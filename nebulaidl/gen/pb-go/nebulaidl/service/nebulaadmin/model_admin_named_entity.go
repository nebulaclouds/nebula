/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Encapsulates information common to a NamedEntity, a Nebula resource such as a task, workflow or launch plan. A NamedEntity is exclusively identified by its resource type and identifier.
type AdminNamedEntity struct {
	// Resource type of the named entity. One of Task, Workflow or LaunchPlan.
	ResourceType *CoreResourceType `json:"resource_type,omitempty"`
	Id *AdminNamedEntityIdentifier `json:"id,omitempty"`
	// Additional metadata around a named entity.
	Metadata *AdminNamedEntityMetadata `json:"metadata,omitempty"`
}
