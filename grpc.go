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

func GRPCCodeToErrorCode(code codes.Code) ErrorCode {
	switch code {
	case codes.InvalidArgument:
		return CodeInvalidArgument
	case codes.Unauthenticated:
		return CodeUnauthenticated
	case codes.PermissionDenied:
		return CodePermissionDenied
	case codes.NotFound:
		return CodeNotFound
	case codes.AlreadyExists:
		return CodeAlreadyExists
	case codes.Aborted:
		return CodeConflict
	case codes.FailedPrecondition:
		return CodePreconditionFailed
	case codes.ResourceExhausted:
		return CodeRateLimited
	case codes.Internal:
		return CodeInternal
	case codes.DeadlineExceeded:
		return CodeTimeout
	case codes.Unavailable:
		return CodeServiceUnavailable
	case codes.Canceled:
		return CodeCanceled
	default:
		return CodeUnknown
	}
}

// ToGRPCStatus converts error to grpc status
func ToGRPCStatus(err error) error {
	var appErr *ErrorInfo

	if !errors.As(err, &appErr) {
		return status.Errorf(codes.Unknown, "internal error: %v", err)
	}

	grpcCode := GRPCCodeFromErrorCode(appErr.Code)
	st := status.New(grpcCode, fmt.Sprintf("%s: %v", appErr.Op, appErr.Err))

	reason := "unspecified"
	if r, ok := appErr.Meta[MetaReason]; ok {
		reason = fmt.Sprintf("%v", r)
	}

	detail := &errdetails.ErrorInfo{
		Reason: string(appErr.Code),
		Domain: getDomain(),
		Metadata: map[string]string{
			"operation": appErr.Op,
			"reason":    reason,
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
