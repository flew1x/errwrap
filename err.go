package errwrap

import (
	"errors"
	"fmt"
	"sync"
)

var (
	defaultDomain     = "service"
	defaultDomainOnce sync.Once
)

func SetDomain(domain string) {
	defaultDomainOnce.Do(func() {
		defaultDomain = domain
	})
}

// ErrorCode is an error code
type ErrorCode string

const (
	// Common
	CodeUnknown  ErrorCode = "UNKNOWN"
	CodeInternal ErrorCode = "INTERNAL"
	CodeCanceled ErrorCode = "CANCELED"
	CodeTimeout  ErrorCode = "TIMEOUT"

	// Validation
	CodeInvalidArgument  ErrorCode = "INVALID_ARGUMENT"
	CodeValidationFailed ErrorCode = "VALIDATION_FAILED"
	CodeOutOfRange       ErrorCode = "OUT_OF_RANGE"
	CodeMissingField     ErrorCode = "MISSING_FIELD"

	// Authentication
	CodeUnauthenticated  ErrorCode = "UNAUTHENTICATED"
	CodePermissionDenied ErrorCode = "PERMISSION_DENIED"
	CodeAccessExpired    ErrorCode = "ACCESS_EXPIRED"
	CodeTokenInvalid     ErrorCode = "TOKEN_INVALID"

	// Resources
	CodeNotFound          ErrorCode = "NOT_FOUND"
	CodeAlreadyExists     ErrorCode = "ALREADY_EXISTS"
	CodeConflict          ErrorCode = "CONFLICT"
	CodeResourceExhausted ErrorCode = "RESOURCE_EXHAUSTED"

	// Network
	CodeDependencyFailed   ErrorCode = "DEPENDENCY_FAILED"
	CodeUpstreamTimeout    ErrorCode = "UPSTREAM_TIMEOUT"
	CodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"

	// Rate limiting
	CodeRateLimited   ErrorCode = "RATE_LIMITED"
	CodeQuotaExceeded ErrorCode = "QUOTA_EXCEEDED"
	CodeBlocked       ErrorCode = "BLOCKED"

	// Business
	CodeBusinessRuleViolated ErrorCode = "BUSINESS_RULE_VIOLATED"
	CodePreconditionFailed   ErrorCode = "PRECONDITION_FAILED"

	//Security
	CodeSecurityViolation ErrorCode = "SECURITY_VIOLATION"

	// Transaction
	CodeTransactionFailed ErrorCode = "TRANSACTION_FAILED"
)

// AppError is an error wrapper
type AppError struct {
	Op     string
	Code   ErrorCode
	Domain string
	Err    error
	Meta   map[string]any
}

// Error returns error message
func (e *AppError) Error() string {
	if e.Meta != nil {
		return fmt.Sprintf("%s: [%s] %v | meta: %v", e.Op, e.Code, e.Err, e.Meta)
	}

	return fmt.Sprintf("%s: [%s] %v", e.Op, e.Code, e.Err)
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// CodeOf returns error code
func CodeOf(err error) ErrorCode {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr.Code
	}

	return CodeUnknown
}

// Wrap wraps an error into AppError
func Wrap(op string, code ErrorCode, err error, meta map[string]any) error {
	if err == nil {
		return nil
	}

	return &AppError{
		Op:   op,
		Code: code,
		Err:  err,
		Meta: meta,
	}
}
