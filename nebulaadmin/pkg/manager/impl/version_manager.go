package impl

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	adminversion "github.com/nebulaclouds/nebula/nebulastdlib/version"
)

type VersionManager struct {
	Version   string
	Build     string
	BuildTime string
}

func (v *VersionManager) GetVersion(ctx context.Context, r *admin.GetVersionRequest) (*admin.GetVersionResponse, error) {
	return &admin.GetVersionResponse{
		ControlPlaneVersion: &admin.Version{
			Version:   v.Version,
			Build:     v.Build,
			BuildTime: v.BuildTime,
		},
	}, nil
}

func NewVersionManager() interfaces.VersionInterface {
	return &VersionManager{
		Build:     adminversion.Build,
		Version:   adminversion.Version,
		BuildTime: adminversion.BuildTime,
	}
}
