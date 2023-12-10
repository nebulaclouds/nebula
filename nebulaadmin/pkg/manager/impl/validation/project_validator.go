package validation

import (
	"context"

	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/impl/shared"
	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

const projectID = "project_id"
const projectName = "project_name"
const projectDescription = "project_description"
const maxNameLength = 64
const maxDescriptionLength = 300
const maxLabelArrayLength = 16

func ValidateProjectRegisterRequest(request admin.ProjectRegisterRequest) error {
	if request.Project == nil {
		return shared.GetMissingArgumentError(shared.Project)
	}
	project := *request.Project
	if err := ValidateEmptyStringField(project.Name, projectName); err != nil {
		return err
	}
	return ValidateProject(project)
}

func ValidateProject(project admin.Project) error {
	if err := ValidateEmptyStringField(project.Id, projectID); err != nil {
		return err
	}
	if err := validateLabels(project.Labels); err != nil {
		return err
	}
	if errs := validation.IsDNS1123Label(project.Id); len(errs) > 0 {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument, "invalid project id [%s]: %v", project.Id, errs)
	}
	if err := ValidateMaxLengthStringField(project.Name, projectName, maxNameLength); err != nil {
		return err
	}
	if err := ValidateMaxLengthStringField(project.Description, projectDescription, maxDescriptionLength); err != nil {
		return err
	}
	if project.Domains != nil {
		return errors.NewNebulaAdminError(codes.InvalidArgument,
			"Domains are currently only set system wide. Please retry without domains included in your request.")
	}
	return nil
}

// Validates that a specified project and domain combination has been registered and exists in the db.
func ValidateProjectAndDomain(
	ctx context.Context, db repositoryInterfaces.Repository, config runtimeInterfaces.ApplicationConfiguration, projectID, domainID string) error {
	project, err := db.ProjectRepo().Get(ctx, projectID)
	if err != nil {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] and domain [%s] are registered, err: [%+v]",
			projectID, domainID, err)
	}
	if *project.State != int32(admin.Project_ACTIVE) {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument,
			"project [%s] is not active", projectID)
	}
	var validDomain bool
	domains := config.GetDomainsConfig()
	for _, domain := range *domains {
		if domain.ID == domainID {
			validDomain = true
			break
		}
	}
	if !validDomain {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument, "domain [%s] is unrecognized by system", domainID)
	}
	return nil
}

func ValidateProjectForUpdate(
	ctx context.Context, db repositoryInterfaces.Repository, projectID string) error {

	project, err := db.ProjectRepo().Get(ctx, projectID)
	if err != nil {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] is registered, err: [%+v]",
			projectID, err)
	}
	if *project.State != int32(admin.Project_ACTIVE) {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument,
			"project [%s] is not active", projectID)
	}
	return nil
}

// ValidateProjectExists doesn't check that the project is active. This is used to get Project level attributes, which you should
// be able to do even for inactive projects.
func ValidateProjectExists(
	ctx context.Context, db repositoryInterfaces.Repository, projectID string) error {

	_, err := db.ProjectRepo().Get(ctx, projectID)
	if err != nil {
		return errors.NewNebulaAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] exists, err: [%+v]",
			projectID, err)
	}
	return nil
}
