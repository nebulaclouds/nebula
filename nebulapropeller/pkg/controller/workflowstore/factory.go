package workflowstore

import (
	"context"
	"fmt"

	nebulaworkflowv1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned/typed/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/listers/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

func NewWorkflowStore(ctx context.Context, cfg *Config, lister v1alpha1.NebulaWorkflowLister,
	workflows nebulaworkflowv1alpha1.NebulaworkflowV1alpha1Interface, scope promutils.Scope) (NebulaWorkflow, error) {

	var workflowStore NebulaWorkflow
	var err error

	switch cfg.Policy {
	case PolicyInMemory:
		workflowStore = NewInMemoryWorkflowStore()
	case PolicyPassThrough:
		workflowStore = NewPassthroughWorkflowStore(ctx, scope, workflows, lister)
	case PolicyTrackTerminated:
		workflowStore = NewPassthroughWorkflowStore(ctx, scope, workflows, lister)
		workflowStore, err = NewTerminatedTrackingStore(ctx, scope, workflowStore)
	case PolicyResourceVersionCache:
		workflowStore = NewPassthroughWorkflowStore(ctx, scope, workflows, lister)
		workflowStore, err = NewTerminatedTrackingStore(ctx, scope, workflowStore)
		workflowStore = NewResourceVersionCachingStore(ctx, scope, workflowStore)
	}

	if err != nil {
		return nil, err
	} else if workflowStore == nil {
		return nil, fmt.Errorf("empty workflow store config")
	}

	return workflowStore, err
}
