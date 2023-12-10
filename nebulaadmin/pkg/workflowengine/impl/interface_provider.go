package impl

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/common"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

// Satisfies
// https://github.com/nebulaclouds/nebulapropeller/blob/9428f5ca8e8fd84189a9105cad55176c8bd9d31c/pkg/compiler/common/reader.go#L49

type InterfaceProvider = common.InterfaceProvider

type LaunchPlanInterfaceProvider struct {
	expectedInputs  core.ParameterMap
	expectedOutputs core.VariableMap
	identifier      *core.Identifier
}

func (p *LaunchPlanInterfaceProvider) GetID() *core.Identifier {
	return p.identifier
}
func (p *LaunchPlanInterfaceProvider) GetExpectedInputs() *core.ParameterMap {
	return &p.expectedInputs

}
func (p *LaunchPlanInterfaceProvider) GetExpectedOutputs() *core.VariableMap {
	return &p.expectedOutputs
}

func NewLaunchPlanInterfaceProvider(launchPlan models.LaunchPlan, identifier core.Identifier) (common.InterfaceProvider, error) {
	var closure admin.LaunchPlanClosure
	err := proto.Unmarshal(launchPlan.Closure, &closure)
	if err != nil {
		logger.Errorf(context.TODO(), "Failed to transform launch plan: %v", err)
		return &LaunchPlanInterfaceProvider{}, err
	}
	return &LaunchPlanInterfaceProvider{
		expectedInputs:  *closure.ExpectedInputs,
		expectedOutputs: *closure.ExpectedOutputs,
		identifier:      &identifier,
	}, nil
}
