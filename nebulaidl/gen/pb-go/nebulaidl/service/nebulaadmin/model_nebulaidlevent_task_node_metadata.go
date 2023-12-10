/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

type NebulaidleventTaskNodeMetadata struct {
	// Captures the status of caching for this execution.
	CacheStatus *CoreCatalogCacheStatus `json:"cache_status,omitempty"`
	CatalogKey *CoreCatalogMetadata `json:"catalog_key,omitempty"`
	// Captures the status of cache reservations for this execution.
	ReservationStatus *CatalogReservationStatus `json:"reservation_status,omitempty"`
	CheckpointUri string `json:"checkpoint_uri,omitempty"`
	// In the case this task launched a dynamic workflow we capture its structure here.
	DynamicWorkflow *NebulaidleventDynamicWorkflowNodeMetadata `json:"dynamic_workflow,omitempty"`
}