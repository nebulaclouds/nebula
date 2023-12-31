package controller

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	listers "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/listers/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
)

var wfs = []*v1alpha1.NebulaWorkflow{
	{
		ExecutionID: v1alpha1.ExecutionID{
			WorkflowExecutionIdentifier: &core.WorkflowExecutionIdentifier{
				Project: "proj",
				Domain:  "dev",
				Name:    "name",
			},
		},
	},
	{
		ExecutionID: v1alpha1.ExecutionID{
			WorkflowExecutionIdentifier: &core.WorkflowExecutionIdentifier{
				Project: "proj",
				Domain:  "dev",
				Name:    "name",
			},
		},
	},
	{
		ExecutionID: v1alpha1.ExecutionID{
			WorkflowExecutionIdentifier: &core.WorkflowExecutionIdentifier{
				Project: "proj2",
				Domain:  "dev",
				Name:    "name",
			},
		},
	},
}

func TestNewResourceLevelMonitor(t *testing.T) {
	lm := ResourceLevelMonitor{}
	res := lm.countList(context.Background(), wfs)
	assert.Equal(t, 2, res["proj"]["dev"])
	assert.Equal(t, 1, res["proj2"]["dev"])
}

type mockWFLister struct {
	listers.NebulaWorkflowLister
}

func (m mockWFLister) List(_ labels.Selector) (ret []*v1alpha1.NebulaWorkflow, err error) {
	return wfs, nil
}

func TestResourceLevelMonitor_collect(t *testing.T) {
	scope := promutils.NewScope("testscope")
	g := labeled.NewGauge("unittest", "testing", scope)
	lm := &ResourceLevelMonitor{
		Scope:          scope,
		CollectorTimer: scope.MustNewStopWatch("collection_cycle", "Measures how long it takes to run a collection", time.Millisecond),
		levels:         g,
		lister:         mockWFLister{},
	}
	lm.collect(context.Background())

	var expected = `
		# HELP testscope:unittest testing
		# TYPE testscope:unittest gauge
		testscope:unittest{domain="dev",project="proj", task="",wf=""} 2
		testscope:unittest{domain="dev",project="proj2", task="",wf=""} 1
	`

	err := testutil.CollectAndCompare(g.GaugeVec, strings.NewReader(expected))
	assert.NoError(t, err)
}

func init() {
	labeled.SetMetricKeys(contextutils.ProjectKey, contextutils.DomainKey, contextutils.WorkflowIDKey, contextutils.TaskIDKey)
}
