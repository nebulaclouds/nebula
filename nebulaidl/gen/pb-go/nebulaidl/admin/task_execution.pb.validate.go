// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nebulaidl/admin/task_execution.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = core.TaskExecution_Phase(0)
)

// define the regex for a UUID once up-front
var _task_execution_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TaskExecutionGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskExecutionGetRequestValidationError is the validation error returned by
// TaskExecutionGetRequest.Validate if the designated constraints aren't met.
type TaskExecutionGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionGetRequestValidationError) ErrorName() string {
	return "TaskExecutionGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionGetRequestValidationError{}

// Validate checks the field values on TaskExecutionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNodeExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionListRequestValidationError{
				field:  "NodeExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Limit

	// no validation rules for Token

	// no validation rules for Filters

	if v, ok := interface{}(m.GetSortBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionListRequestValidationError{
				field:  "SortBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskExecutionListRequestValidationError is the validation error returned by
// TaskExecutionListRequest.Validate if the designated constraints aren't met.
type TaskExecutionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionListRequestValidationError) ErrorName() string {
	return "TaskExecutionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionListRequestValidationError{}

// Validate checks the field values on TaskExecution with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TaskExecution) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InputUri

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsParent

	return nil
}

// TaskExecutionValidationError is the validation error returned by
// TaskExecution.Validate if the designated constraints aren't met.
type TaskExecutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionValidationError) ErrorName() string { return "TaskExecutionValidationError" }

// Error satisfies the builtin error interface
func (e TaskExecutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionValidationError{}

// Validate checks the field values on TaskExecutionList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TaskExecutionList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTaskExecutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionListValidationError{
					field:  fmt.Sprintf("TaskExecutions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// TaskExecutionListValidationError is the validation error returned by
// TaskExecutionList.Validate if the designated constraints aren't met.
type TaskExecutionListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionListValidationError) ErrorName() string {
	return "TaskExecutionListValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionListValidationError{}

// Validate checks the field values on TaskExecutionClosure with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionClosure) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Phase

	for idx, item := range m.GetLogs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionClosureValidationError{
					field:  fmt.Sprintf("Logs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCustomInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "CustomInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Reason

	// no validation rules for TaskType

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionClosureValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EventVersion

	for idx, item := range m.GetReasons() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionClosureValidationError{
					field:  fmt.Sprintf("Reasons[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.OutputResult.(type) {

	case *TaskExecutionClosure_OutputUri:
		// no validation rules for OutputUri

	case *TaskExecutionClosure_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionClosureValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TaskExecutionClosure_OutputData:

		if v, ok := interface{}(m.GetOutputData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionClosureValidationError{
					field:  "OutputData",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TaskExecutionClosureValidationError is the validation error returned by
// TaskExecutionClosure.Validate if the designated constraints aren't met.
type TaskExecutionClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionClosureValidationError) ErrorName() string {
	return "TaskExecutionClosureValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionClosureValidationError{}

// Validate checks the field values on Reason with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Reason) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReasonValidationError{
				field:  "OccurredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Message

	return nil
}

// ReasonValidationError is the validation error returned by Reason.Validate if
// the designated constraints aren't met.
type ReasonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReasonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReasonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReasonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReasonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReasonValidationError) ErrorName() string { return "ReasonValidationError" }

// Error satisfies the builtin error interface
func (e ReasonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReason.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReasonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReasonValidationError{}

// Validate checks the field values on TaskExecutionGetDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionGetDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskExecutionGetDataRequestValidationError is the validation error returned
// by TaskExecutionGetDataRequest.Validate if the designated constraints
// aren't met.
type TaskExecutionGetDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionGetDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionGetDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionGetDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionGetDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionGetDataRequestValidationError) ErrorName() string {
	return "TaskExecutionGetDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionGetDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionGetDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionGetDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionGetDataRequestValidationError{}

// Validate checks the field values on TaskExecutionGetDataResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionGetDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataResponseValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataResponseValidationError{
				field:  "Outputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFullInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataResponseValidationError{
				field:  "FullInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFullOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataResponseValidationError{
				field:  "FullOutputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetNebulaUrls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionGetDataResponseValidationError{
				field:  "NebulaUrls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskExecutionGetDataResponseValidationError is the validation error returned
// by TaskExecutionGetDataResponse.Validate if the designated constraints
// aren't met.
type TaskExecutionGetDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionGetDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionGetDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionGetDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionGetDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionGetDataResponseValidationError) ErrorName() string {
	return "TaskExecutionGetDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionGetDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionGetDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionGetDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionGetDataResponseValidationError{}
