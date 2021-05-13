package lib

import (
	"goframework/util"
)

const (
	InternaError    int = 1000
	SystemError     int = 1001
	ParamError      int = 1002
	BadRequestError int = 1003
	NotFoundError   int = 1004
)

//自定义api错误结构体
type Error struct {
	Code    int
	Message string
}

func (err *Error) Error() string {
	return err.Message
}

func NewParamError(msg string) *Error {
	return &Error{
		Code:    ParamError,
		Message: msg,
	}
}

func NewAutoParamError(err error) *Error {
	msg := util.ValidateParams(err)
	return NewParamError(msg)
}

func NewBadRequestError(msg string) *Error {
	return &Error{
		Code:    BadRequestError,
		Message: msg,
	}
}

func NewInternalError(msg string) *Error {
	return &Error{
		Code:    InternaError,
		Message: msg,
	}
}

func NewNotFoundError(msg string) *Error {
	return &Error{
		Code:    NotFoundError,
		Message: msg,
	}
}
