/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Request to terminate an in-progress execution.  This action is irreversible. If an execution is already terminated, this request will simply be a no-op. This request will fail if it references a non-existent execution. If the request succeeds the phase \"ABORTED\" will be recorded for the termination with the optional cause added to the output_result.
type AdminExecutionTerminateRequest struct {
	// Uniquely identifies the individual workflow execution to be terminated.
	Id *CoreWorkflowExecutionIdentifier `json:"id,omitempty"`
	// Optional reason for aborting.
	Cause string `json:"cause,omitempty"`
}
