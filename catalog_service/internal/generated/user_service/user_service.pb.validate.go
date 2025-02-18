// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user_service/user_service.proto

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

// Validate checks the field values on SigninRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SigninRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SigninRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SigninRequestMultiError, or
// nil if none found.
func (m *SigninRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SigninRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return SigninRequestMultiError(errors)
	}

	return nil
}

// SigninRequestMultiError is an error wrapping multiple validation errors
// returned by SigninRequest.ValidateAll() if the designated constraints
// aren't met.
type SigninRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SigninRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SigninRequestMultiError) AllErrors() []error { return m }

// SigninRequestValidationError is the validation error returned by
// SigninRequest.Validate if the designated constraints aren't met.
type SigninRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SigninRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SigninRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SigninRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SigninRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SigninRequestValidationError) ErrorName() string { return "SigninRequestValidationError" }

// Error satisfies the builtin error interface
func (e SigninRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSigninRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SigninRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SigninRequestValidationError{}

// Validate checks the field values on SigninResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SigninResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SigninResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SigninResponseMultiError,
// or nil if none found.
func (m *SigninResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SigninResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SigninResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SessionId

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if all {
		switch v := interface{}(m.GetAccessTokenExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "AccessTokenExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "AccessTokenExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccessTokenExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SigninResponseValidationError{
				field:  "AccessTokenExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRefreshTokenExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "RefreshTokenExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SigninResponseValidationError{
					field:  "RefreshTokenExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshTokenExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SigninResponseValidationError{
				field:  "RefreshTokenExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SigninResponseMultiError(errors)
	}

	return nil
}

// SigninResponseMultiError is an error wrapping multiple validation errors
// returned by SigninResponse.ValidateAll() if the designated constraints
// aren't met.
type SigninResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SigninResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SigninResponseMultiError) AllErrors() []error { return m }

// SigninResponseValidationError is the validation error returned by
// SigninResponse.Validate if the designated constraints aren't met.
type SigninResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SigninResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SigninResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SigninResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SigninResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SigninResponseValidationError) ErrorName() string { return "SigninResponseValidationError" }

// Error satisfies the builtin error interface
func (e SigninResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSigninResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SigninResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SigninResponseValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserResponseMultiError, or nil if none found.
func (m *GetUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserResponseMultiError(errors)
	}

	return nil
}

// GetUserResponseMultiError is an error wrapping multiple validation errors
// returned by GetUserResponse.ValidateAll() if the designated constraints
// aren't met.
type GetUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserResponseMultiError) AllErrors() []error { return m }

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetPasswordChangedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "PasswordChangedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "PasswordChangedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPasswordChangedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "PasswordChangedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}
