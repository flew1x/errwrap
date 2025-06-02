package errwrap

import (
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCCodeFromErrorCode returns grpc code from error code
func GRPCCodeFromErrorCode(code ErrorCode) codes.Code {
	switch code {
	// === 4xx — ===
	case CodeInvalidArgument, CodeValidationFailed, CodeMissingField, CodeOutOfRange:
		return codes.InvalidArgument
	case CodeUnauthenticated, CodeTokenInvalid, CodeAccessExpired:
		return codes.Unauthenticated
	case CodePermissionDenied:
		return codes.PermissionDenied
	case CodeNotFound:
		return codes.NotFound
	case CodeAlreadyExists:
		return codes.AlreadyExists
	case CodeConflict:
		return codes.Aborted
	case CodePreconditionFailed, CodeBusinessRuleViolated:
		return codes.FailedPrecondition
	case CodeRateLimited, CodeQuotaExceeded, CodeBlocked:
		return codes.ResourceExhausted

	// === 5xx — ===
	case CodeInternal, CodeTransactionFailed:
		return codes.Internal
	case CodeTimeout, CodeUpstreamTimeout:
		return codes.DeadlineExceeded
	case CodeDependencyFailed:
		return codes.FailedPrecondition
	case CodeServiceUnavailable:
		return codes.Unavailable
	case CodeSecurityViolation:
		return codes.PermissionDenied
	case CodeCanceled:
		return codes.Canceled
	default:
		return codes.Unknown
	}
}

// ToGRPCStatus converts error to grpc status
func ToGRPCStatus(err error) error {
	var appErr *AppError

	if !errors.As(err, &appErr) {
		return status.Errorf(codes.Unknown, "internal error: %v", err)
	}

	grpcCode := GRPCCodeFromErrorCode(appErr.Code)
	st := status.New(grpcCode, fmt.Sprintf("%s: %v", appErr.Op, appErr.Err))

	detail := &errdetails.ErrorInfo{
		Reason: string(appErr.Code),
		Domain: getDomain(),
		Metadata: map[string]string{
			"operataion": appErr.Op,
		},
	}

	for k, v := range appErr.Meta {
		detail.Metadata[k] = fmt.Sprintf("%v", v)
	}

	stWithDetails, errWithDetails := st.WithDetails(detail)
	if errWithDetails != nil {
		return st.Err()
	}

	return stWithDetails.Err()
}
