/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// The unique identifier for a successfully created execution. If the name was *not* specified in the create request, this identifier will include a generated name.
type AdminExecutionCreateResponse struct {
	Id *CoreWorkflowExecutionIdentifier `json:"id,omitempty"`
}
