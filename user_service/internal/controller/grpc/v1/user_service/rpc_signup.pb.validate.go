// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user_service/rpc_signup.proto

package user_service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SignupRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupRequestMultiError, or
// nil if none found.
func (m *SignupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return SignupRequestMultiError(errors)
	}

	return nil
}

// SignupRequestMultiError is an error wrapping multiple validation errors
// returned by SignupRequest.ValidateAll() if the designated constraints
// aren't met.
type SignupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupRequestMultiError) AllErrors() []error { return m }

// SignupRequestValidationError is the validation error returned by
// SignupRequest.Validate if the designated constraints aren't met.
type SignupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupRequestValidationError) ErrorName() string { return "SignupRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupRequestValidationError{}

// Validate checks the field values on SignupResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupResponseMultiError,
// or nil if none found.
func (m *SignupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignupResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignupResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignupResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SignupResponseMultiError(errors)
	}

	return nil
}

// SignupResponseMultiError is an error wrapping multiple validation errors
// returned by SignupResponse.ValidateAll() if the designated constraints
// aren't met.
type SignupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupResponseMultiError) AllErrors() []error { return m }

// SignupResponseValidationError is the validation error returned by
// SignupResponse.Validate if the designated constraints aren't met.
type SignupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupResponseValidationError) ErrorName() string { return "SignupResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupResponseValidationError{}
