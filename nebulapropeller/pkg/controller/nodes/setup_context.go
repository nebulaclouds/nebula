package nodes

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type setupContext struct {
	enq   func(string)
	scope promutils.Scope
}

func (s *setupContext) EnqueueOwner() func(string) {
	return s.enq
}

func (s *setupContext) OwnerKind() string {
	return v1alpha1.NebulaWorkflowKind
}

func (s *setupContext) MetricsScope() promutils.Scope {
	return s.scope
}

func (c *recursiveNodeExecutor) newSetupContext(_ context.Context) interfaces.SetupContext {
	return &setupContext{
		enq:   c.enqueueWorkflow,
		scope: c.metrics.Scope,
	}
}
