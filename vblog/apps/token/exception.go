package token

import (
	"net/http"

	"github.com/qinchi-ops/govlog/vblog/exception"
)

var (
	ErrAuthFailed          = exception.NewApiException(50001, "用户名或者密码不正确").WithHttpCode(http.StatusUnauthorized)
	ErrAccessTokenExpired  = exception.NewApiException(50002, "AccessToken过期")
	ErrRefreshTokenExpired = exception.NewApiException(50003, "RefreshToken过期")
	ErrRPermissionDeny     = exception.NewApiException(50004, "当前用户无权限访问").WithHttpCode(http.StatusForbidden)
)
