package token

import "github.com/qinchi-ops/govlog/vblog/exception"

var (
	ErrAuthFailed = exception.NewApiException(50001, "用户名或者密码不正确")
)
