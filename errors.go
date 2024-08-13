package xerr

import (
	"fmt"
)

type CodeError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("Code: %d, msg: %s", e.Code, e.Message)
}

func NewErrCodeMsg(errCode int64, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Message: errMsg}
}

func NewErrCode(errCode int64) *CodeError {
	return &CodeError{
		Code:    errCode,
		Message: MapErrMsg(errCode),
	}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{Code: ServerCommonError, Message: errMsg}
}
