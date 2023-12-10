package validators

import (
	nebula "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	c "github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/common"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/errors"
)

func validateOutputVar(n c.NodeBuilder, paramName string, errs errors.CompileErrors) (
	param *nebula.Variable, ok bool) {
	if outputs, effectiveOk := validateEffectiveOutputParameters(n, errs.NewScope()); effectiveOk {
		var paramFound bool
		if param, paramFound = findVariableByName(outputs, paramName); !paramFound {
			errs.Collect(errors.NewVariableNameNotFoundErr(n.GetId(), n.GetId(), paramName))
		}
	}

	return param, !errs.HasErrors()
}

func validateInputVar(n c.NodeBuilder, paramName string, requireParamType bool, errs errors.CompileErrors) (param *nebula.Variable, ok bool) {
	if n.GetInterface() == nil {
		return nil, false
	}

	if param, ok = findVariableByName(n.GetInterface().GetInputs(), paramName); ok {
		return
	}

	if !requireParamType {
		if containsBindingByVariableName(n.GetInputs(), paramName) {
			return
		}
	}

	errs.Collect(errors.NewVariableNameNotFoundErr(n.GetId(), n.GetId(), paramName))
	return
}

func validateVarType(nodeID c.NodeID, paramName string, param *nebula.Variable,
	expectedType *nebula.LiteralType, errs errors.CompileErrors) (ok bool) {
	if param.GetType().String() != expectedType.String() {
		errs.Collect(errors.NewMismatchingTypesErr(nodeID, paramName, param.GetType().String(), expectedType.String()))
	}

	return !errs.HasErrors()
}

// Validate parameters have their required attributes set
func validateVariables(nodeID c.NodeID, params *nebula.VariableMap, errs errors.CompileErrors) {
	for paramName, param := range params.Variables {
		if len(paramName) == 0 {
			errs.Collect(errors.NewValueRequiredErr(nodeID, "paramName"))
		}

		if param.Type == nil {
			errs.Collect(errors.NewValueRequiredErr(nodeID, "param.Type"))
		}
	}
}
