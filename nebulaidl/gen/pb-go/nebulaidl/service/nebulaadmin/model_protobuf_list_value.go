/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// `ListValue` is a wrapper around a repeated field of values.  The JSON representation for `ListValue` is JSON array.
type ProtobufListValue struct {
	// Repeated field of dynamically typed values.
	Values []ProtobufValue `json:"values,omitempty"`
}