package workflowstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	kubeerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned/fake"
	listers "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/listers/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type mockWFNamespaceLister struct {
	listers.NebulaWorkflowNamespaceLister
	GetCb func(name string) (*v1alpha1.NebulaWorkflow, error)
}

func (m *mockWFNamespaceLister) Get(name string) (*v1alpha1.NebulaWorkflow, error) {
	return m.GetCb(name)
}

type mockWFLister struct {
	listers.NebulaWorkflowLister
	V listers.NebulaWorkflowNamespaceLister
}

func (m *mockWFLister) NebulaWorkflows(namespace string) listers.NebulaWorkflowNamespaceLister {
	return m.V
}

func TestPassthroughWorkflowStore_Get(t *testing.T) {
	ctx := context.TODO()

	mockClient := fake.NewSimpleClientset().NebulaworkflowV1alpha1()

	l := &mockWFNamespaceLister{}
	wfStore := NewPassthroughWorkflowStore(ctx, promutils.NewTestScope(), mockClient, &mockWFLister{V: l})

	t.Run("notFound", func(t *testing.T) {
		l.GetCb = func(name string) (*v1alpha1.NebulaWorkflow, error) {
			return nil, kubeerrors.NewNotFound(v1alpha1.Resource(v1alpha1.NebulaWorkflowKind), "name")
		}
		w, err := wfStore.Get(ctx, "ns", "name")
		assert.Error(t, err)
		assert.True(t, IsNotFound(err))
		assert.Nil(t, w)
	})

	t.Run("alreadyExists?", func(t *testing.T) {
		l.GetCb = func(name string) (*v1alpha1.NebulaWorkflow, error) {
			return nil, kubeerrors.NewAlreadyExists(v1alpha1.Resource(v1alpha1.NebulaWorkflowKind), "name")
		}
		w, err := wfStore.Get(ctx, "ns", "name")
		assert.Error(t, err)
		assert.Nil(t, w)
	})

	t.Run("unknownError", func(t *testing.T) {
		l.GetCb = func(name string) (*v1alpha1.NebulaWorkflow, error) {
			return nil, fmt.Errorf("error")
		}
		w, err := wfStore.Get(ctx, "ns", "name")
		assert.Error(t, err)
		assert.Nil(t, w)
	})

	t.Run("success", func(t *testing.T) {
		expW := &v1alpha1.NebulaWorkflow{}
		l.GetCb = func(name string) (*v1alpha1.NebulaWorkflow, error) {
			return expW, nil
		}
		w, err := wfStore.Get(ctx, "ns", "name")
		assert.NoError(t, err)
		assert.Equal(t, expW, w)
	})
}

func dummyWf(namespace, name string) *v1alpha1.NebulaWorkflow {
	return &v1alpha1.NebulaWorkflow{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func TestPassthroughWorkflowStore_UpdateStatus(t *testing.T) {

	ctx := context.TODO()

	mockClient := fake.NewSimpleClientset().NebulaworkflowV1alpha1()
	l := &mockWFNamespaceLister{}
	wfStore := NewPassthroughWorkflowStore(ctx, promutils.NewTestScope(), mockClient, &mockWFLister{V: l})

	const namespace = "test-ns"
	t.Run("notFound", func(t *testing.T) {
		wf := dummyWf(namespace, "x")
		_, err := wfStore.UpdateStatus(ctx, wf, PriorityClassCritical)
		assert.NoError(t, err)
		updated, err := mockClient.NebulaWorkflows(namespace).Get(ctx, "x", v1.GetOptions{})
		assert.Error(t, err)
		assert.Nil(t, updated)
	})

	t.Run("Found-Updated", func(t *testing.T) {
		n := mockClient.NebulaWorkflows(namespace)
		wf := dummyWf(namespace, "x")
		wf.GetExecutionStatus().UpdatePhase(v1alpha1.WorkflowPhaseSucceeding, "", nil)
		wf.ResourceVersion = "r1"
		_, err := n.Create(ctx, wf, v1.CreateOptions{})
		assert.NoError(t, err)
		updated, err := n.Get(ctx, "x", v1.GetOptions{})
		if assert.NoError(t, err) {
			assert.Equal(t, v1alpha1.WorkflowPhaseSucceeding, updated.GetExecutionStatus().GetPhase())
			wf.GetExecutionStatus().UpdatePhase(v1alpha1.WorkflowPhaseFailed, "", &core.ExecutionError{})
			_, err := wfStore.UpdateStatus(ctx, wf, PriorityClassCritical)
			assert.NoError(t, err)
			newVal, err := n.Get(ctx, "x", v1.GetOptions{})
			assert.NoError(t, err)
			assert.Equal(t, v1alpha1.WorkflowPhaseFailed, newVal.GetExecutionStatus().GetPhase())
		}
	})

}
