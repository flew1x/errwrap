package errwrap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// HTTPCodeFromErrorCode returns http code from error code
func HTTPCodeFromErrorCode(code ErrorCode) int {
	switch code {
	case CodeInvalidArgument, CodeValidationFailed, CodeMissingField, CodeOutOfRange:
		return http.StatusBadRequest
	case CodeUnauthenticated, CodeTokenInvalid, CodeAccessExpired:
		return http.StatusUnauthorized
	case CodePermissionDenied, CodeSecurityViolation:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeAlreadyExists, CodeConflict:
		return http.StatusConflict
	case CodeRateLimited:
		return http.StatusTooManyRequests
	case CodeServiceUnavailable, CodeUpstreamTimeout:
		return http.StatusServiceUnavailable
	case CodeTimeout:
		return http.StatusGatewayTimeout
	case CodeInternal, CodeTransactionFailed, CodeDependencyFailed:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// ErrorResponse defines error response
type ErrorResponse struct {
	Code      ErrorCode      `json:"code"`
	Message   string         `json:"message"`
	Operation string         `json:"operation,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
}

// WriteHTTPError writes AppError to http.ResponseWriter
func WriteHTTPError(w http.ResponseWriter, err error) {
	var appErr *ErrorInfo
	if !errors.As(err, &appErr) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	status := HTTPCodeFromErrorCode(appErr.Code)

	resp := ErrorResponse{
		Code:      appErr.Code,
		Message:   appErr.Err.Error(),
		Operation: appErr.Op,
		Meta:      appErr.Meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("failed to write error response: %v\n", err)
	}
}
