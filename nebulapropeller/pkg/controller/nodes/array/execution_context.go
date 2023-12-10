package array

import (
	"strconv"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/executors"
)

const (
	NebulaK8sArrayIndexVarName string = "NEBULA_K8S_ARRAY_INDEX"
	JobIndexVarName           string = "BATCH_JOB_ARRAY_INDEX_VAR_NAME"
)

type arrayExecutionContext struct {
	executors.ExecutionContext
	executionConfig    v1alpha1.ExecutionConfig
	currentParallelism *uint32
}

func (a *arrayExecutionContext) GetExecutionConfig() v1alpha1.ExecutionConfig {
	return a.executionConfig
}

func (a *arrayExecutionContext) CurrentParallelism() uint32 {
	return *a.currentParallelism
}

func (a *arrayExecutionContext) IncrementParallelism() uint32 {
	*a.currentParallelism = *a.currentParallelism + 1
	return *a.currentParallelism
}

func newArrayExecutionContext(executionContext executors.ExecutionContext, subNodeIndex int, currentParallelism *uint32, maxParallelism uint32) *arrayExecutionContext {
	executionConfig := executionContext.GetExecutionConfig()
	if executionConfig.EnvironmentVariables == nil {
		executionConfig.EnvironmentVariables = make(map[string]string)
	}
	executionConfig.EnvironmentVariables[JobIndexVarName] = NebulaK8sArrayIndexVarName
	executionConfig.EnvironmentVariables[NebulaK8sArrayIndexVarName] = strconv.Itoa(subNodeIndex)
	executionConfig.MaxParallelism = maxParallelism

	return &arrayExecutionContext{
		ExecutionContext:   executionContext,
		executionConfig:    executionConfig,
		currentParallelism: currentParallelism,
	}
}
