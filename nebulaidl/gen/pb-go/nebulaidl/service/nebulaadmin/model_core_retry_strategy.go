/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Retry strategy associated with an executable unit.
type CoreRetryStrategy struct {
	// Number of retries. Retries will be consumed when the job fails with a recoverable error. The number of retries must be less than or equals to 10.
	Retries int64 `json:"retries,omitempty"`
}
