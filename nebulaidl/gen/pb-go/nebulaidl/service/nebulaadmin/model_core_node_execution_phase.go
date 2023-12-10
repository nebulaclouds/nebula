/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

type CoreNodeExecutionPhase string

// List of coreNodeExecutionPhase
const (
	CoreNodeExecutionPhaseUNDEFINED CoreNodeExecutionPhase = "UNDEFINED"
	CoreNodeExecutionPhaseQUEUED CoreNodeExecutionPhase = "QUEUED"
	CoreNodeExecutionPhaseRUNNING CoreNodeExecutionPhase = "RUNNING"
	CoreNodeExecutionPhaseSUCCEEDED CoreNodeExecutionPhase = "SUCCEEDED"
	CoreNodeExecutionPhaseFAILING CoreNodeExecutionPhase = "FAILING"
	CoreNodeExecutionPhaseFAILED CoreNodeExecutionPhase = "FAILED"
	CoreNodeExecutionPhaseABORTED CoreNodeExecutionPhase = "ABORTED"
	CoreNodeExecutionPhaseSKIPPED CoreNodeExecutionPhase = "SKIPPED"
	CoreNodeExecutionPhaseTIMED_OUT CoreNodeExecutionPhase = "TIMED_OUT"
	CoreNodeExecutionPhaseDYNAMIC_RUNNING CoreNodeExecutionPhase = "DYNAMIC_RUNNING"
	CoreNodeExecutionPhaseRECOVERED CoreNodeExecutionPhase = "RECOVERED"
)