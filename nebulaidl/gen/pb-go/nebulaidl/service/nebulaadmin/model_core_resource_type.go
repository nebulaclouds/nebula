/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin
// CoreResourceType : Indicates a resource type within Nebula.   - DATASET: A dataset represents an entity modeled in Nebula DataCatalog. A Dataset is also a versioned entity and can be a compilation of multiple individual objects. Eventually all Catalog objects should be modeled similar to Nebula Objects. The Dataset entities makes it possible for the UI  and CLI to act on the objects  in a similar manner to other Nebula objects
type CoreResourceType string

// List of coreResourceType
const (
	CoreResourceTypeUNSPECIFIED CoreResourceType = "UNSPECIFIED"
	CoreResourceTypeTASK CoreResourceType = "TASK"
	CoreResourceTypeWORKFLOW CoreResourceType = "WORKFLOW"
	CoreResourceTypeLAUNCH_PLAN CoreResourceType = "LAUNCH_PLAN"
	CoreResourceTypeDATASET CoreResourceType = "DATASET"
)