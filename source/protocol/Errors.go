package protocol

import "fmt"

const (
	ErrorInvalidRequest  = "invalid_request"
	ErrorUnauthorized    = "unauthorized"
	ErrorForbidden       = "forbidden"
	ErrorNotFound        = "not_found"
	ErrorConflict        = "conflict"
	ErrorUnavailable     = "unavailable"
	ErrorTimeout         = "timeout"
	ErrorInternal        = "internal"
	ErrorNotImplemented  = "not_implemented"
	ErrorProtocolVersion = "protocol_version"
)

type Error struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data,omitempty"`
}

func NewError(code string, message string) *Error {

	return &Error{
		Code:    code,
		Message: message,
		Data:    make(map[string]string),
	}

}

func NewErrorf(code string, format string, values ...any) *Error {

	return NewError(code, fmt.Sprintf(format, values...))

}

func (err *Error) Error() string {

	if err == nil {
		return ""
	}

	if err.Code != "" {
		return err.Code + ": " + err.Message
	}

	return err.Message

}

func (err *Error) With(key string, value string) *Error {

	if err.Data == nil {
		err.Data = make(map[string]string)
	}

	err.Data[key] = value

	return err

}
