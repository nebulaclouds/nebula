package workflowstore

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

//go:generate mockery -all

type PriorityClass int

const (
	PriorityClassCritical PriorityClass = iota
	PriorityClassRegular
)

// NebulaWorkflow store interface provides an abstraction of accessing the actual NebulaWorkflow object.
type NebulaWorkflow interface {
	Get(ctx context.Context, namespace, name string) (*v1alpha1.NebulaWorkflow, error)
	UpdateStatus(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass PriorityClass) (
		newWF *v1alpha1.NebulaWorkflow, err error)
	Update(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass PriorityClass) (
		newWF *v1alpha1.NebulaWorkflow, err error)
}
