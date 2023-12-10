package impl

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/transformers/k8s"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type builderMetrics struct {
	Scope                  promutils.Scope
	WorkflowBuildSuccesses prometheus.Counter
	WorkflowBuildFailures  prometheus.Counter
	InvalidExecutionID     prometheus.Counter
}

type nebulaWorkflowBuilder struct {
	metrics builderMetrics
}

func (b *nebulaWorkflowBuilder) Build(
	wfClosure *core.CompiledWorkflowClosure, inputs *core.LiteralMap, executionID *core.WorkflowExecutionIdentifier,
	namespace string) (*v1alpha1.NebulaWorkflow, error) {
	nebulaWorkflow, err := k8s.BuildNebulaWorkflow(wfClosure, inputs, executionID, namespace)
	if err != nil {
		b.metrics.WorkflowBuildFailures.Inc()
		return nil, err
	}
	b.metrics.WorkflowBuildSuccesses.Inc()
	return nebulaWorkflow, nil
}

func newBuilderMetrics(scope promutils.Scope) builderMetrics {
	return builderMetrics{
		Scope: scope,
		WorkflowBuildSuccesses: scope.MustNewCounter("build_successes",
			"count of workflows built by propeller without error"),
		WorkflowBuildFailures: scope.MustNewCounter("build_failures",
			"count of workflows built by propeller with errors"),
	}
}

func NewNebulaWorkflowBuilder(scope promutils.Scope) interfaces.NebulaWorkflowBuilder {
	return &nebulaWorkflowBuilder{
		metrics: newBuilderMetrics(scope),
	}
}
