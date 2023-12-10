// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nebulaidl/admin/signal.proto

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
)

// define the regex for a UUID once up-front
var _signal_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SignalGetOrCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignalGetOrCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalGetOrCreateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalGetOrCreateRequestValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SignalGetOrCreateRequestValidationError is the validation error returned by
// SignalGetOrCreateRequest.Validate if the designated constraints aren't met.
type SignalGetOrCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalGetOrCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalGetOrCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalGetOrCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalGetOrCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalGetOrCreateRequestValidationError) ErrorName() string {
	return "SignalGetOrCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignalGetOrCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignalGetOrCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalGetOrCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalGetOrCreateRequestValidationError{}

// Validate checks the field values on SignalListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignalListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetWorkflowExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalListRequestValidationError{
				field:  "WorkflowExecutionId",
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
			return SignalListRequestValidationError{
				field:  "SortBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SignalListRequestValidationError is the validation error returned by
// SignalListRequest.Validate if the designated constraints aren't met.
type SignalListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalListRequestValidationError) ErrorName() string {
	return "SignalListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignalListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignalListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalListRequestValidationError{}

// Validate checks the field values on SignalList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SignalList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSignals() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignalListValidationError{
					field:  fmt.Sprintf("Signals[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// SignalListValidationError is the validation error returned by
// SignalList.Validate if the designated constraints aren't met.
type SignalListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalListValidationError) ErrorName() string { return "SignalListValidationError" }

// Error satisfies the builtin error interface
func (e SignalListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignalList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalListValidationError{}

// Validate checks the field values on SignalSetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignalSetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalSetRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalSetRequestValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SignalSetRequestValidationError is the validation error returned by
// SignalSetRequest.Validate if the designated constraints aren't met.
type SignalSetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalSetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalSetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalSetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalSetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalSetRequestValidationError) ErrorName() string { return "SignalSetRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignalSetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignalSetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalSetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalSetRequestValidationError{}

// Validate checks the field values on SignalSetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignalSetResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SignalSetResponseValidationError is the validation error returned by
// SignalSetResponse.Validate if the designated constraints aren't met.
type SignalSetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalSetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalSetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalSetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalSetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalSetResponseValidationError) ErrorName() string {
	return "SignalSetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SignalSetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignalSetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalSetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalSetResponseValidationError{}

// Validate checks the field values on Signal with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Signal) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignalValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SignalValidationError is the validation error returned by Signal.Validate if
// the designated constraints aren't met.
type SignalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignalValidationError) ErrorName() string { return "SignalValidationError" }

// Error satisfies the builtin error interface
func (e SignalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignalValidationError{}