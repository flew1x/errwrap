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
