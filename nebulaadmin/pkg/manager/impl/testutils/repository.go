package testutils

import (
	"context"

	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	repositoryMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

func GetRepoWithDefaultProjectAndErr(err error) repositoryInterfaces.Repository {
	repo := repositoryMocks.NewMockRepository()
	repo.ProjectRepo().(*repositoryMocks.MockProjectRepo).GetFunction = func(
		ctx context.Context, projectID string) (models.Project, error) {
		activeState := int32(admin.Project_ACTIVE)
		return models.Project{State: &activeState}, err
	}
	return repo
}

func GetRepoWithDefaultProject() repositoryInterfaces.Repository {
	repo := repositoryMocks.NewMockRepository()
	repo.ProjectRepo().(*repositoryMocks.MockProjectRepo).GetFunction = func(
		ctx context.Context, projectID string) (models.Project, error) {
		activeState := int32(admin.Project_ACTIVE)
		return models.Project{State: &activeState}, nil
	}
	return repo
}
