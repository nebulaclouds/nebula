/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

import (
	"time"
)

// A container holding the compiled workflow produced from the WorkflowSpec and additional metadata.
type AdminWorkflowClosure struct {
	// Represents the compiled representation of the workflow from the specification provided.
	CompiledWorkflow *CoreCompiledWorkflowClosure `json:"compiled_workflow,omitempty"`
	// Time at which the workflow was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
}