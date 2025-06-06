package errs

const (
	ErrBadRequest       = 400_000
	ErrMsgBadRequest    = "Bad Request"
	ErrNotFound         = 404_000
	ErrMsgNotFound      = "Not Found"
	ErrUnauthorized     = 401_000
	ErrMsgUnauthorized  = "Unauthorized"
	ErrConflict         = 402_000
	ErrMsgConflict      = "Conflict"
	ErrAlreadyExists    = 400_100
	ErrMsgAlreadyExists = "Already Exists"
	ErrPermissionDenied = 403_000
	ErrMsgPermissionDenied = "Permission Denied"

	ErrInternalError = 500_000
	ErrMsgInternalError = "Internal Server Error"
	ErrDatabase      = 500_100
	ErrMsgDatabase   = "Database Error"
	ErrKafka         = 500_200
	ErrMsgKafka      = "Kafka Error"
	ErrThirdParty    = 500_400
	ErrMsgThirdParty = "Third Party Error"
)

var ErrMapping = map[int]string{
	ErrBadRequest:       ErrMsgBadRequest,
	ErrNotFound:         ErrMsgNotFound,
	ErrUnauthorized:     ErrMsgUnauthorized,
	ErrConflict:         ErrMsgConflict,
	ErrAlreadyExists:    ErrMsgAlreadyExists,
	ErrPermissionDenied: ErrMsgPermissionDenied,
	ErrInternalError:    ErrMsgInternalError,
	ErrDatabase:         ErrMsgDatabase,
	ErrKafka:            ErrMsgKafka,
	ErrThirdParty:       ErrMsgThirdParty,
}
