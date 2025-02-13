// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: catalog_service/catalog_service.proto

package catalog_service

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

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRequestMultiError, or nil if none found.
func (m *CreateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Image

	// no validation rules for CategoryId

	// no validation rules for BrandId

	for idx, item := range m.GetSkus() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateProductRequestValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateProductRequestValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateProductRequestValidationError{
					field:  fmt.Sprintf("Skus[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateProductRequestMultiError(errors)
	}

	return nil
}

// CreateProductRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRequestMultiError) AllErrors() []error { return m }

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on SKUToCreate with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SKUToCreate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SKUToCreate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SKUToCreateMultiError, or
// nil if none found.
func (m *SKUToCreate) ValidateAll() error {
	return m.validate(true)
}

func (m *SKUToCreate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Slug

	// no validation rules for OriginalPrice

	// no validation rules for InitialStock

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SKUToCreateValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SKUToCreateValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SKUToCreateValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SKUToCreateMultiError(errors)
	}

	return nil
}

// SKUToCreateMultiError is an error wrapping multiple validation errors
// returned by SKUToCreate.ValidateAll() if the designated constraints aren't met.
type SKUToCreateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SKUToCreateMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SKUToCreateMultiError) AllErrors() []error { return m }

// SKUToCreateValidationError is the validation error returned by
// SKUToCreate.Validate if the designated constraints aren't met.
type SKUToCreateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SKUToCreateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SKUToCreateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SKUToCreateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SKUToCreateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SKUToCreateValidationError) ErrorName() string { return "SKUToCreateValidationError" }

// Error satisfies the builtin error interface
func (e SKUToCreateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSKUToCreate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SKUToCreateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SKUToCreateValidationError{}

// Validate checks the field values on AttributeValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AttributeValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeValue with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttributeValueMultiError,
// or nil if none found.
func (m *AttributeValue) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AttributeId

	// no validation rules for Value

	if len(errors) > 0 {
		return AttributeValueMultiError(errors)
	}

	return nil
}

// AttributeValueMultiError is an error wrapping multiple validation errors
// returned by AttributeValue.ValidateAll() if the designated constraints
// aren't met.
type AttributeValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeValueMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeValueMultiError) AllErrors() []error { return m }

// AttributeValueValidationError is the validation error returned by
// AttributeValue.Validate if the designated constraints aren't met.
type AttributeValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValueValidationError) ErrorName() string { return "AttributeValueValidationError" }

// Error satisfies the builtin error interface
func (e AttributeValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValueValidationError{}

// Validate checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductResponseMultiError, or nil if none found.
func (m *CreateProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateProductResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProductResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSkus() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateProductResponseValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateProductResponseValidationError{
						field:  fmt.Sprintf("Skus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateProductResponseValidationError{
					field:  fmt.Sprintf("Skus[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateProductResponseMultiError(errors)
	}

	return nil
}

// CreateProductResponseMultiError is an error wrapping multiple validation
// errors returned by CreateProductResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductResponseMultiError) AllErrors() []error { return m }

// CreateProductResponseValidationError is the validation error returned by
// CreateProductResponse.Validate if the designated constraints aren't met.
type CreateProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductResponseValidationError) ErrorName() string {
	return "CreateProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductResponseValidationError{}

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Image

	// no validation rules for CategoryId

	// no validation rules for BrandId

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on SKU with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *SKU) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SKU with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SKUMultiError, or nil if none found.
func (m *SKU) ValidateAll() error {
	return m.validate(true)
}

func (m *SKU) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProductId

	// no validation rules for Name

	// no validation rules for Slug

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SKUValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SKUValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SKUValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCurrentPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "CurrentPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "CurrentPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCurrentPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SKUValidationError{
				field:  "CurrentPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetInventory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "Inventory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "Inventory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInventory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SKUValidationError{
				field:  "Inventory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SKUValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SKUValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SKUValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SKUMultiError(errors)
	}

	return nil
}

// SKUMultiError is an error wrapping multiple validation errors returned by
// SKU.ValidateAll() if the designated constraints aren't met.
type SKUMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SKUMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SKUMultiError) AllErrors() []error { return m }

// SKUValidationError is the validation error returned by SKU.Validate if the
// designated constraints aren't met.
type SKUValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SKUValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SKUValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SKUValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SKUValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SKUValidationError) ErrorName() string { return "SKUValidationError" }

// Error satisfies the builtin error interface
func (e SKUValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSKU.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SKUValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SKUValidationError{}

// Validate checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PriceMultiError, or nil if none found.
func (m *Price) ValidateAll() error {
	return m.validate(true)
}

func (m *Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for SkuId

	// no validation rules for OriginalPrice

	if all {
		switch v := interface{}(m.GetEffectiveDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriceValidationError{
					field:  "EffectiveDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriceValidationError{
					field:  "EffectiveDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEffectiveDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriceValidationError{
				field:  "EffectiveDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Active

	if len(errors) > 0 {
		return PriceMultiError(errors)
	}

	return nil
}

// PriceMultiError is an error wrapping multiple validation errors returned by
// Price.ValidateAll() if the designated constraints aren't met.
type PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PriceMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PriceMultiError) AllErrors() []error { return m }

// PriceValidationError is the validation error returned by Price.Validate if
// the designated constraints aren't met.
type PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriceValidationError) ErrorName() string { return "PriceValidationError" }

// Error satisfies the builtin error interface
func (e PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriceValidationError{}

// Validate checks the field values on Inventory with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Inventory) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Inventory with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InventoryMultiError, or nil
// if none found.
func (m *Inventory) ValidateAll() error {
	return m.validate(true)
}

func (m *Inventory) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for SkuId

	// no validation rules for Stock

	// no validation rules for Reservations

	if len(errors) > 0 {
		return InventoryMultiError(errors)
	}

	return nil
}

// InventoryMultiError is an error wrapping multiple validation errors returned
// by Inventory.ValidateAll() if the designated constraints aren't met.
type InventoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InventoryMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InventoryMultiError) AllErrors() []error { return m }

// InventoryValidationError is the validation error returned by
// Inventory.Validate if the designated constraints aren't met.
type InventoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InventoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InventoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InventoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InventoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InventoryValidationError) ErrorName() string { return "InventoryValidationError" }

// Error satisfies the builtin error interface
func (e InventoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInventory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InventoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InventoryValidationError{}
