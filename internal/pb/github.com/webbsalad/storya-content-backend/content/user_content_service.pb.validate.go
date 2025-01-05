// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/content/user_content_service.proto

package content

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

// define the regex for a UUID once up-front
var _user_content_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetUserItemsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserItemsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserItemsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserItemsRequestMultiError, or nil if none found.
func (m *GetUserItemsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserItemsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = GetUserItemsRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ContentType

	if len(errors) > 0 {
		return GetUserItemsRequestMultiError(errors)
	}

	return nil
}

func (m *GetUserItemsRequest) _validateUuid(uuid string) error {
	if matched := _user_content_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetUserItemsRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserItemsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserItemsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserItemsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserItemsRequestMultiError) AllErrors() []error { return m }

// GetUserItemsRequestValidationError is the validation error returned by
// GetUserItemsRequest.Validate if the designated constraints aren't met.
type GetUserItemsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserItemsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserItemsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserItemsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserItemsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserItemsRequestValidationError) ErrorName() string {
	return "GetUserItemsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserItemsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserItemsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserItemsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserItemsRequestValidationError{}

// Validate checks the field values on GetUserItemsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserItemsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserItemsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserItemsResponseMultiError, or nil if none found.
func (m *GetUserItemsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserItemsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUserItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUserItemsResponseValidationError{
						field:  fmt.Sprintf("UserItems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUserItemsResponseValidationError{
						field:  fmt.Sprintf("UserItems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserItemsResponseValidationError{
					field:  fmt.Sprintf("UserItems[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetUserItemsResponseMultiError(errors)
	}

	return nil
}

// GetUserItemsResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserItemsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserItemsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserItemsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserItemsResponseMultiError) AllErrors() []error { return m }

// GetUserItemsResponseValidationError is the validation error returned by
// GetUserItemsResponse.Validate if the designated constraints aren't met.
type GetUserItemsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserItemsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserItemsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserItemsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserItemsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserItemsResponseValidationError) ErrorName() string {
	return "GetUserItemsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserItemsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserItemsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserItemsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserItemsResponseValidationError{}

// Validate checks the field values on GetValuedRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetValuedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetValuedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetValuedRequestMultiError, or nil if none found.
func (m *GetValuedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetValuedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = GetValuedRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ContentType

	// no validation rules for Value

	if len(errors) > 0 {
		return GetValuedRequestMultiError(errors)
	}

	return nil
}

func (m *GetValuedRequest) _validateUuid(uuid string) error {
	if matched := _user_content_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetValuedRequestMultiError is an error wrapping multiple validation errors
// returned by GetValuedRequest.ValidateAll() if the designated constraints
// aren't met.
type GetValuedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetValuedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetValuedRequestMultiError) AllErrors() []error { return m }

// GetValuedRequestValidationError is the validation error returned by
// GetValuedRequest.Validate if the designated constraints aren't met.
type GetValuedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetValuedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetValuedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetValuedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetValuedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetValuedRequestValidationError) ErrorName() string { return "GetValuedRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetValuedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetValuedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetValuedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetValuedRequestValidationError{}

// Validate checks the field values on GetValuedResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetValuedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetValuedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetValuedResponseMultiError, or nil if none found.
func (m *GetValuedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetValuedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetValuedResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetValuedResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetValuedResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetValuedResponseMultiError(errors)
	}

	return nil
}

// GetValuedResponseMultiError is an error wrapping multiple validation errors
// returned by GetValuedResponse.ValidateAll() if the designated constraints
// aren't met.
type GetValuedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetValuedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetValuedResponseMultiError) AllErrors() []error { return m }

// GetValuedResponseValidationError is the validation error returned by
// GetValuedResponse.Validate if the designated constraints aren't met.
type GetValuedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetValuedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetValuedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetValuedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetValuedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetValuedResponseValidationError) ErrorName() string {
	return "GetValuedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetValuedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetValuedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetValuedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetValuedResponseValidationError{}

// Validate checks the field values on AddRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddRequestMultiError, or
// nil if none found.
func (m *AddRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetItemId()); err != nil {
		err = AddRequestValidationError{
			field:  "ItemId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ContentType

	// no validation rules for Value

	if len(errors) > 0 {
		return AddRequestMultiError(errors)
	}

	return nil
}

func (m *AddRequest) _validateUuid(uuid string) error {
	if matched := _user_content_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// AddRequestMultiError is an error wrapping multiple validation errors
// returned by AddRequest.ValidateAll() if the designated constraints aren't met.
type AddRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddRequestMultiError) AllErrors() []error { return m }

// AddRequestValidationError is the validation error returned by
// AddRequest.Validate if the designated constraints aren't met.
type AddRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddRequestValidationError) ErrorName() string { return "AddRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddRequestValidationError{}

// Validate checks the field values on RemoveItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RemoveItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveItemRequestMultiError, or nil if none found.
func (m *RemoveItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetItemId()); err != nil {
		err = RemoveItemRequestValidationError{
			field:  "ItemId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ContentType

	if len(errors) > 0 {
		return RemoveItemRequestMultiError(errors)
	}

	return nil
}

func (m *RemoveItemRequest) _validateUuid(uuid string) error {
	if matched := _user_content_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// RemoveItemRequestMultiError is an error wrapping multiple validation errors
// returned by RemoveItemRequest.ValidateAll() if the designated constraints
// aren't met.
type RemoveItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveItemRequestMultiError) AllErrors() []error { return m }

// RemoveItemRequestValidationError is the validation error returned by
// RemoveItemRequest.Validate if the designated constraints aren't met.
type RemoveItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveItemRequestValidationError) ErrorName() string {
	return "RemoveItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveItemRequestValidationError{}
