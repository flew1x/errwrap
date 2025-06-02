package errwrap

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
