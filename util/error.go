package util

import "fmt"

type CommonError struct {
	apiName string
	ErrCode int64  `json:"error_code"`
	Message string `json:"message"`
}

func (c *CommonError) Error() string {
	return fmt.Sprintf("%s Error , errcode=%d , errmsg=%s", c.apiName, c.ErrCode, c.Message)
}

// NewCommonError 新建 CommonError 错误，对于无 errcode 和 errmsg 的返回也可以返回该通用错误
func NewCommonError(apiName string, code int64, msg string) *CommonError {
	return &CommonError{
		apiName: apiName,
		ErrCode: code,
		Message: msg,
	}
}
