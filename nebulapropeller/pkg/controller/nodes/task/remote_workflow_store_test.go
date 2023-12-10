package task

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func createInmemoryStore(t testing.TB) *storage.DataStore {
	cfg := storage.Config{
		Type: storage.TypeMemory,
	}

	d, err := storage.NewDataStore(&cfg, promutils.NewTestScope())
	assert.NoError(t, err)

	return d
}

func Test_cacheNebulaWorkflow(t *testing.T) {
	store := createInmemoryStore(t)
	t.Run("cache CRD", func(t *testing.T) {
		expected := &v1alpha1.NebulaWorkflow{
			TypeMeta:   v1.TypeMeta{},
			ObjectMeta: v1.ObjectMeta{},
			WorkflowSpec: &v1alpha1.WorkflowSpec{
				ID: "abc",
				Connections: v1alpha1.Connections{
					Downstream: map[v1alpha1.NodeID][]v1alpha1.NodeID{},
					Upstream:   map[v1alpha1.NodeID][]v1alpha1.NodeID{},
				},
				DeprecatedConnections: v1alpha1.DeprecatedConnections{
					DownstreamEdges: map[v1alpha1.NodeID][]v1alpha1.NodeID{},
					UpstreamEdges:   map[v1alpha1.NodeID][]v1alpha1.NodeID{},
				},
			},
		}

		ctx := context.TODO()
		location := storage.DataReference("somekey/file.json")
		r := RemoteFileWorkflowStore{store: store}
		assert.NoError(t, r.PutNebulaWorkflowCRD(ctx, expected, location))
		actual, err := r.GetWorkflowCRD(ctx, location)
		assert.NoError(t, err)
		if diff := deep.Equal(expected, actual); len(diff) > 0 {
			t.Errorf("Cached() Diff = %v\r\n got = %v\r\n want = %v", diff, actual, expected)
		}
	})

	t.Run("cache CRD with deprecatedConnections", func(t *testing.T) {
		expected := &v1alpha1.NebulaWorkflow{
			TypeMeta:   v1.TypeMeta{},
			ObjectMeta: v1.ObjectMeta{},
			WorkflowSpec: &v1alpha1.WorkflowSpec{
				ID: "abc",
				DeprecatedConnections: v1alpha1.DeprecatedConnections{
					DownstreamEdges: map[v1alpha1.NodeID][]v1alpha1.NodeID{},
					UpstreamEdges:   map[v1alpha1.NodeID][]v1alpha1.NodeID{},
				},
			},
			ExecutionConfig: v1alpha1.ExecutionConfig{
				TaskResources: v1alpha1.TaskResources{
					Requests: v1alpha1.TaskResourceSpec{
						CPU:              resource.MustParse("1"),
						Memory:           resource.MustParse("1"),
						Storage:          resource.MustParse("1"),
						EphemeralStorage: resource.MustParse("1"),
						GPU:              resource.MustParse("1"),
					},
					Limits: v1alpha1.TaskResourceSpec{
						CPU:              resource.MustParse("1"),
						Memory:           resource.MustParse("1"),
						Storage:          resource.MustParse("1"),
						EphemeralStorage: resource.MustParse("1"),
						GPU:              resource.MustParse("1"),
					},
				},
			},
		}

		ctx := context.TODO()
		location := storage.DataReference("somekey/file.json")
		r := RemoteFileWorkflowStore{store: store}
		assert.NoError(t, r.PutNebulaWorkflowCRD(ctx, expected, location))
		actual, err := r.GetWorkflowCRD(ctx, location)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("cache compiled workflow", func(t *testing.T) {
		expected := &core.CompiledWorkflowClosure{
			Primary: &core.CompiledWorkflow{
				Template: &core.WorkflowTemplate{
					Id: &core.Identifier{
						ResourceType: core.ResourceType_WORKFLOW,
						Project:      "proj",
						Domain:       "domain",
						Name:         "name",
						Version:      "version",
					},
				},
			},
		}

		ctx := context.TODO()
		location := storage.DataReference("somekey/dynamic_compiled.pb")
		r := RemoteFileWorkflowStore{store: store}
		assert.NoError(t, r.PutCompiledNebulaWorkflow(ctx, expected, location))
		actual, err := r.GetCompiledWorkflow(ctx, location)
		assert.NoError(t, err)
		assert.True(t, proto.Equal(expected, actual))
	})
}
