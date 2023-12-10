package tests

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

func TestRegisterProject(t *testing.T) {
	ctx := context.Background()

	mockProjectManager := mocks.MockProjectManager{}
	mockProjectManager.SetCreateProject(
		func(ctx context.Context,
			request admin.ProjectRegisterRequest) (*admin.ProjectRegisterResponse, error) {
			return &admin.ProjectRegisterResponse{}, nil
		},
	)
	mockServer := NewMockAdminServer(NewMockAdminServerInput{
		projectManager: &mockProjectManager,
	})

	resp, err := mockServer.RegisterProject(ctx, &admin.ProjectRegisterRequest{
		Project: &admin.Project{
			Id: "project",
		},
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}

func TestListProjects(t *testing.T) {
	mockProjectManager := mocks.MockProjectManager{}
	projects := &admin.Projects{
		Projects: []*admin.Project{
			{
				Id:   "project id",
				Name: "project name",
				Domains: []*admin.Domain{
					{
						Id:   "domain id",
						Name: "domain name",
					},
				},
			},
		},
	}
	mockProjectManager.SetListCallback(func(ctx context.Context, request admin.ProjectListRequest) (*admin.Projects, error) {
		assert.NotNil(t, request)
		return projects, nil
	})
	resp, err := mockProjectManager.ListProjects(context.Background(), admin.ProjectListRequest{})
	assert.Nil(t, err)
	assert.True(t, proto.Equal(projects, resp))
}
