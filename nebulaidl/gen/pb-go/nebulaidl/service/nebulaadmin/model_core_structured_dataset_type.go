/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

type CoreStructuredDatasetType struct {
	// A list of ordered columns this schema comprises of.
	Columns []StructuredDatasetTypeDatasetColumn `json:"columns,omitempty"`
	// This is the storage format, the format of the bits at rest parquet, feather, csv, etc. For two types to be compatible, the format will need to be an exact match.
	Format string `json:"format,omitempty"`
	// This is a string representing the type that the bytes in external_schema_bytes are formatted in. This is an optional field that will not be used for type checking.
	ExternalSchemaType string `json:"external_schema_type,omitempty"`
	// The serialized bytes of a third-party schema library like Arrow. This is an optional field that will not be used for type checking.
	ExternalSchemaBytes string `json:"external_schema_bytes,omitempty"`
}