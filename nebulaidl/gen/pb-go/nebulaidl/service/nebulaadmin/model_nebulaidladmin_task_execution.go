/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Encapsulates all details for a single task execution entity. A task execution represents an instantiated task, including all inputs and additional metadata as well as computed results included state, outputs, and duration-based attributes.
type NebulaidladminTaskExecution struct {
	// Unique identifier for the task execution.
	Id *CoreTaskExecutionIdentifier `json:"id,omitempty"`
	// Path to remote data store where input blob is stored.
	InputUri string `json:"input_uri,omitempty"`
	// Task execution details and results.
	Closure *AdminTaskExecutionClosure `json:"closure,omitempty"`
	// Whether this task spawned nodes.
	IsParent bool `json:"is_parent,omitempty"`
}