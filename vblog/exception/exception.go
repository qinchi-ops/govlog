package exception

import "encoding/json"

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// 用于描述业务异常
// 实现
type ApiException struct {
	// 业务异常的编码，50001表示token过期
	Code int `json:"code"`
	// 异常描述信息
	Message string `json:"message"`
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// type error interface {
// 	Error() string
// }

func (e *ApiException) Error() string {
	return e.Message
}

func (e *ApiException) String() string {
	dj, _ := json.MarshalIndent(e, "", " ")
	return string(dj)
}
