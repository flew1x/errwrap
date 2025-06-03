package errwrap

// Meta keys
const (
	MetaStage  = "stage"
	MetaReason = "reason"
	MetaField  = "field"
	MetaUserID = "user_id"
	MetaSQL    = "sql"
)

// Reason values
const (
	ReasonUnauthorized        = "unauthorized"
	ReasonTokenInvalid        = "token_invalid"
	ReasonDatabaseUnavailable = "database_unavailable"
	ReasonExecSQLError        = "exec_sql_error"
	ReasonScanRowError        = "scan_row_error"
	ReasonBuildSQLError       = "build_sql_error"
	ReasonValidationFailed    = "validation_failed"
)

// Stage values
const (
	StagePrepareRequest = "prepare_request"
	StageBuildSQL       = "build_sql"
	StageExecSQL        = "exec_sql"
	StageScanRow        = "scan_row"
	StageValidation     = "validation"
	StageTokenDecode    = "token_decode"
	StageCheckAuth      = "check_auth"
	StageTransformData  = "transform_data"
	StageSendRequest    = "send_request"
)
