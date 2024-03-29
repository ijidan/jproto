// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gateway.proto

package proto_build

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

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetGatewayId()) < 1 {
		err := RegisterRequestValidationError{
			field:  "GatewayId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetClientId()) < 1 {
		err := RegisterRequestValidationError{
			field:  "ClientId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterResponseMultiError, or nil if none found.
func (m *RegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterResponseMultiError(errors)
	}

	return nil
}

// RegisterResponseMultiError is an error wrapping multiple validation errors
// returned by RegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type RegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterResponseMultiError) AllErrors() []error { return m }

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on UnRegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnRegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnRegisterRequestMultiError, or nil if none found.
func (m *UnRegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnRegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetGatewayId()) < 1 {
		err := UnRegisterRequestValidationError{
			field:  "GatewayId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetClientId()) < 1 {
		err := UnRegisterRequestValidationError{
			field:  "ClientId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UnRegisterRequestMultiError(errors)
	}

	return nil
}

// UnRegisterRequestMultiError is an error wrapping multiple validation errors
// returned by UnRegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type UnRegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnRegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnRegisterRequestMultiError) AllErrors() []error { return m }

// UnRegisterRequestValidationError is the validation error returned by
// UnRegisterRequest.Validate if the designated constraints aren't met.
type UnRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnRegisterRequestValidationError) ErrorName() string {
	return "UnRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UnRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnRegisterRequestValidationError{}

// Validate checks the field values on UnRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnRegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnRegisterResponseMultiError, or nil if none found.
func (m *UnRegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UnRegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnRegisterResponseMultiError(errors)
	}

	return nil
}

// UnRegisterResponseMultiError is an error wrapping multiple validation errors
// returned by UnRegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type UnRegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnRegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnRegisterResponseMultiError) AllErrors() []error { return m }

// UnRegisterResponseValidationError is the validation error returned by
// UnRegisterResponse.Validate if the designated constraints aren't met.
type UnRegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnRegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnRegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnRegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnRegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnRegisterResponseValidationError) ErrorName() string {
	return "UnRegisterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UnRegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnRegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnRegisterResponseValidationError{}

// Validate checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageRequestMultiError, or nil if none found.
func (m *SendMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetGatewayId()) < 1 {
		err := SendMessageRequestValidationError{
			field:  "GatewayId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetData()) < 1 {
		err := SendMessageRequestValidationError{
			field:  "Data",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendMessageRequestMultiError(errors)
	}

	return nil
}

// SendMessageRequestMultiError is an error wrapping multiple validation errors
// returned by SendMessageRequest.ValidateAll() if the designated constraints
// aren't met.
type SendMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageRequestMultiError) AllErrors() []error { return m }

// SendMessageRequestValidationError is the validation error returned by
// SendMessageRequest.Validate if the designated constraints aren't met.
type SendMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageRequestValidationError) ErrorName() string {
	return "SendMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageRequestValidationError{}

// Validate checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageResponseMultiError, or nil if none found.
func (m *SendMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetData()) < 1 {
		err := SendMessageResponseValidationError{
			field:  "Data",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendMessageResponseMultiError(errors)
	}

	return nil
}

// SendMessageResponseMultiError is an error wrapping multiple validation
// errors returned by SendMessageResponse.ValidateAll() if the designated
// constraints aren't met.
type SendMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageResponseMultiError) AllErrors() []error { return m }

// SendMessageResponseValidationError is the validation error returned by
// SendMessageResponse.Validate if the designated constraints aren't met.
type SendMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageResponseValidationError) ErrorName() string {
	return "SendMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageResponseValidationError{}

// Validate checks the field values on PushToAllRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PushToAllRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushToAllRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PushToAllRequestMultiError, or nil if none found.
func (m *PushToAllRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PushToAllRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetGatewayId()) < 1 {
		err := PushToAllRequestValidationError{
			field:  "GatewayId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetData()) < 1 {
		err := PushToAllRequestValidationError{
			field:  "Data",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PushToAllRequestMultiError(errors)
	}

	return nil
}

// PushToAllRequestMultiError is an error wrapping multiple validation errors
// returned by PushToAllRequest.ValidateAll() if the designated constraints
// aren't met.
type PushToAllRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushToAllRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushToAllRequestMultiError) AllErrors() []error { return m }

// PushToAllRequestValidationError is the validation error returned by
// PushToAllRequest.Validate if the designated constraints aren't met.
type PushToAllRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushToAllRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushToAllRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushToAllRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushToAllRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushToAllRequestValidationError) ErrorName() string { return "PushToAllRequestValidationError" }

// Error satisfies the builtin error interface
func (e PushToAllRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushToAllRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushToAllRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushToAllRequestValidationError{}

// Validate checks the field values on PushToAllResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PushToAllResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushToAllResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PushToAllResponseMultiError, or nil if none found.
func (m *PushToAllResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PushToAllResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetData() <= 0 {
		err := PushToAllResponseValidationError{
			field:  "Data",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PushToAllResponseMultiError(errors)
	}

	return nil
}

// PushToAllResponseMultiError is an error wrapping multiple validation errors
// returned by PushToAllResponse.ValidateAll() if the designated constraints
// aren't met.
type PushToAllResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushToAllResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushToAllResponseMultiError) AllErrors() []error { return m }

// PushToAllResponseValidationError is the validation error returned by
// PushToAllResponse.Validate if the designated constraints aren't met.
type PushToAllResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushToAllResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushToAllResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushToAllResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushToAllResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushToAllResponseValidationError) ErrorName() string {
	return "PushToAllResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PushToAllResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushToAllResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushToAllResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushToAllResponseValidationError{}
