package impl

import (
	"google.golang.org/grpc/codes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

func addMapValues(overrides map[string]string, defaultValues map[string]string) map[string]string {
	if defaultValues == nil {
		defaultValues = map[string]string{}
	}
	if overrides == nil {
		return defaultValues
	}
	for label, value := range overrides {
		defaultValues[label] = value
	}
	return defaultValues
}

func addPermissions(securityCtx *core.SecurityContext, roleNameKey string, nebulaWf *v1alpha1.NebulaWorkflow) {
	if securityCtx == nil || securityCtx.RunAs == nil {
		return
	}
	nebulaWf.SecurityContext = *securityCtx
	if len(securityCtx.RunAs.IamRole) > 0 {
		if nebulaWf.Annotations == nil {
			nebulaWf.Annotations = map[string]string{}
		}
		nebulaWf.Annotations[roleNameKey] = securityCtx.RunAs.IamRole
	}
	if len(securityCtx.RunAs.K8SServiceAccount) > 0 {
		nebulaWf.ServiceAccountName = securityCtx.RunAs.K8SServiceAccount
	}
}

func addExecutionOverrides(taskPluginOverrides []*admin.PluginOverride,
	workflowExecutionConfig *admin.WorkflowExecutionConfig, recoveryExecution *core.WorkflowExecutionIdentifier,
	taskResources *interfaces.TaskResources, nebulaWf *v1alpha1.NebulaWorkflow) {
	executionConfig := v1alpha1.ExecutionConfig{
		TaskPluginImpls: make(map[string]v1alpha1.TaskPluginOverride),
		RecoveryExecution: v1alpha1.WorkflowExecutionIdentifier{
			WorkflowExecutionIdentifier: recoveryExecution,
		},
	}
	for _, override := range taskPluginOverrides {
		executionConfig.TaskPluginImpls[override.TaskType] = v1alpha1.TaskPluginOverride{
			PluginIDs:             override.PluginId,
			MissingPluginBehavior: override.MissingPluginBehavior,
		}

	}
	if workflowExecutionConfig != nil {
		executionConfig.MaxParallelism = uint32(workflowExecutionConfig.MaxParallelism)

		if workflowExecutionConfig.GetInterruptible() != nil {
			interruptible := workflowExecutionConfig.GetInterruptible().GetValue()
			executionConfig.Interruptible = &interruptible
		}

		executionConfig.OverwriteCache = workflowExecutionConfig.GetOverwriteCache()

		envs := make(map[string]string)
		if workflowExecutionConfig.GetEnvs() != nil {
			for _, v := range workflowExecutionConfig.GetEnvs().Values {
				envs[v.Key] = v.Value
			}
			executionConfig.EnvironmentVariables = envs
		}
	}
	if taskResources != nil {
		var requests = v1alpha1.TaskResourceSpec{}
		if !taskResources.Defaults.CPU.IsZero() {
			requests.CPU = taskResources.Defaults.CPU
		}
		if !taskResources.Defaults.Memory.IsZero() {
			requests.Memory = taskResources.Defaults.Memory
		}
		if !taskResources.Defaults.EphemeralStorage.IsZero() {
			requests.EphemeralStorage = taskResources.Defaults.EphemeralStorage
		}
		if !taskResources.Defaults.Storage.IsZero() {
			requests.Storage = taskResources.Defaults.Storage
		}
		if !taskResources.Defaults.GPU.IsZero() {
			requests.GPU = taskResources.Defaults.GPU
		}

		var limits = v1alpha1.TaskResourceSpec{}
		if !taskResources.Limits.CPU.IsZero() {
			limits.CPU = taskResources.Limits.CPU
		}
		if !taskResources.Limits.Memory.IsZero() {
			limits.Memory = taskResources.Limits.Memory
		}
		if !taskResources.Limits.EphemeralStorage.IsZero() {
			limits.EphemeralStorage = taskResources.Limits.EphemeralStorage
		}
		if !taskResources.Limits.Storage.IsZero() {
			limits.Storage = taskResources.Limits.Storage
		}
		if !taskResources.Limits.GPU.IsZero() {
			limits.GPU = taskResources.Limits.GPU
		}
		executionConfig.TaskResources = v1alpha1.TaskResources{
			Requests: requests,
			Limits:   limits,
		}
	}
	nebulaWf.ExecutionConfig = executionConfig
}

func PrepareNebulaWorkflow(data interfaces.ExecutionData, nebulaWorkflow *v1alpha1.NebulaWorkflow) error {
	if data.ExecutionID == nil {
		return errors.NewNebulaAdminErrorf(codes.Internal, "invalid execution id")
	}
	if nebulaWorkflow == nil {
		return errors.NewNebulaAdminErrorf(codes.Internal, "missing Nebula Workflow")
	}

	// add the executionId so Propeller can send events back that are associated with the ID
	nebulaWorkflow.ExecutionID = v1alpha1.WorkflowExecutionIdentifier{
		WorkflowExecutionIdentifier: data.ExecutionID,
	}
	// add the acceptedAt timestamp so propeller can emit latency metrics.
	acceptAtWrapper := v1.NewTime(data.ExecutionParameters.AcceptedAt)
	nebulaWorkflow.AcceptedAt = &acceptAtWrapper

	// add permissions from auth and security context. Adding permissions from auth would be removed once all clients
	// have migrated over to security context
	addPermissions(data.ExecutionParameters.ExecutionConfig.SecurityContext,
		data.ExecutionParameters.RoleNameKey, nebulaWorkflow)

	labels := addMapValues(data.ExecutionParameters.Labels, nebulaWorkflow.Labels)
	nebulaWorkflow.Labels = labels
	annotations := addMapValues(data.ExecutionParameters.Annotations, nebulaWorkflow.Annotations)
	nebulaWorkflow.Annotations = annotations
	if nebulaWorkflow.WorkflowMeta == nil {
		nebulaWorkflow.WorkflowMeta = &v1alpha1.WorkflowMeta{}
	}
	nebulaWorkflow.WorkflowMeta.EventVersion = v1alpha1.EventVersion(data.ExecutionParameters.EventVersion)
	addExecutionOverrides(data.ExecutionParameters.TaskPluginOverrides, data.ExecutionParameters.ExecutionConfig,
		data.ExecutionParameters.RecoveryExecution, data.ExecutionParameters.TaskResources, nebulaWorkflow)

	if data.ExecutionParameters.RawOutputDataConfig != nil {
		nebulaWorkflow.RawOutputDataConfig = v1alpha1.RawOutputDataConfig{
			RawOutputDataConfig: data.ExecutionParameters.RawOutputDataConfig,
		}
	}

	return nil
}
