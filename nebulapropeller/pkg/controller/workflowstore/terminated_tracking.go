package workflowstore

import (
	"context"
	"fmt"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/fastcheck"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

func workflowKey(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}

// A specialized store that stores a LRU cache of all the workflows that are in a terminal phase.
// Terminated workflows are ignored (Get returns a nil).
// Processing terminated NebulaWorkflows can occur when workflow updates are reported after a workflow has already completed.
type terminatedTracking struct {
	w                NebulaWorkflow
	terminatedFilter fastcheck.Filter
}

func (t *terminatedTracking) Get(ctx context.Context, namespace, name string) (*v1alpha1.NebulaWorkflow, error) {
	if t.terminatedFilter.Contains(ctx, []byte(workflowKey(namespace, name))) {
		return nil, ErrWorkflowTerminated
	}

	return t.w.Get(ctx, namespace, name)
}

func (t *terminatedTracking) UpdateStatus(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass PriorityClass) (
	newWF *v1alpha1.NebulaWorkflow, err error) {
	newWF, err = t.w.UpdateStatus(ctx, workflow, priorityClass)
	if err != nil {
		return nil, err
	}

	if newWF != nil {
		if newWF.GetExecutionStatus().IsTerminated() {
			t.terminatedFilter.Add(ctx, []byte(workflowKey(workflow.Namespace, workflow.Name)))
		}
	}

	return newWF, nil
}

func (t *terminatedTracking) Update(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass PriorityClass) (
	newWF *v1alpha1.NebulaWorkflow, err error) {
	newWF, err = t.w.Update(ctx, workflow, priorityClass)
	if err != nil {
		return nil, err
	}

	if newWF != nil {
		if newWF.GetExecutionStatus().IsTerminated() {
			t.terminatedFilter.Add(ctx, []byte(workflowKey(workflow.Namespace, workflow.Name)))
		}
	}

	return newWF, nil
}

func NewTerminatedTrackingStore(_ context.Context, scope promutils.Scope, workflowStore NebulaWorkflow) (NebulaWorkflow, error) {
	filter, err := fastcheck.NewLRUCacheFilter(1000, scope.NewSubScope("terminated_filter"))
	if err != nil {
		return nil, err
	}

	return &terminatedTracking{
		w:                workflowStore,
		terminatedFilter: filter,
	}, nil
}
