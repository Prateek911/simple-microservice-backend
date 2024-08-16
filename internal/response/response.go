package response

type ErrorType string

const (
	ErrorNone           ErrorType = ""
	ErrorTimeout        ErrorType = "timeout"
	ErrorCanceled       ErrorType = "canceled"
	ErrorExec           ErrorType = "execution"
	ErrorBadData        ErrorType = "bad_data"
	ErrorInternal       ErrorType = "internal"
	ErrorUnavailable    ErrorType = "unavailable"
	ErrorNotFound       ErrorType = "not_found"
	ErrorNotImplemented ErrorType = "not_implemented"
	ErrorUnauthorized   ErrorType = "unauthorized"
	ErrorForbidden      ErrorType = "forbidden"
	ErrorConflict       ErrorType = "conflict"
)

type BaseApiError interface {
	Type() ErrorType
	isNil() bool
	Error() string
	ToError() error
}

type ApiError struct {
	Typ ErrorType
	Err error
}

func (a *ApiError) Type() ErrorType {
	return a.Typ
}

func (a *ApiError) isNil() bool {
	return a == nil || a.Err == nil
}

func (a *ApiError) Error() string {
	if a == nil || a.Err == nil {
		return ""
	}
	return a.Err.Error()

}

func BadRequest(err error) *ApiError {
	return &ApiError{
		Typ: ErrorBadData,
		Err: err,
	}
}

func RequestTimeOut(err error) *ApiError {
	return &ApiError{
		Typ: ErrorTimeout,
		Err: err,
	}
}

func InternalError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorInternal,
		Err: err,
	}
}

func Unauthorized(err error) *ApiError {
	return &ApiError{
		Typ: ErrorUnauthorized,
		Err: err,
	}
}

func NotFoundError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorNotFound,
		Err: err,
	}
}

func ForbiddenError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorForbidden,
		Err: err,
	}
}

func UnauthorizedError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorUnauthorized,
		Err: err,
	}
}
