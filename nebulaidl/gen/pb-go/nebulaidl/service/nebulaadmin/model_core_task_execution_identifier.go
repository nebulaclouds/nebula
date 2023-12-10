/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Encapsulation of fields that identify a Nebula task execution entity.
type CoreTaskExecutionIdentifier struct {
	TaskId *CoreIdentifier `json:"task_id,omitempty"`
	NodeExecutionId *CoreNodeExecutionIdentifier `json:"node_execution_id,omitempty"`
	RetryAttempt int64 `json:"retry_attempt,omitempty"`
}
