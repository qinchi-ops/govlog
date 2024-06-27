package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/qinchi-ops/govlog/vblog/responese"
)

// Gin Web中间件,  我们需要在中间件注入到请求的链路当中，然后由Gin框架来调用
// HandlerFunc defines the handler used by gin middleware as return value.
// type HandlerFunc func(*Context)
// 加一下中间件处理函数 GIn HandlerFunc

func Auth(ctx *gin.Context) {
	//补充鉴权逻辑
	accessToken, _ := ctx.Cookie(token.COOKIE_TOKEN_KEY)
	tk, err := ioc.Controller.Get(token.AppName).(token.Service).ValidateToken(ctx.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		//相应报错
		responese.Failed(token.ErrAuthFailed.WithMessage(err.Error()), ctx)
		ctx.Abort()
	} else {
		// 鉴权成功, 请求继续往后面进行

		// 后面的handler 怎么知道 鉴权成功了, 当前是谁在访问这个接口
		// 请求的上下文:
		// 怎么把中间件请求结果，添加到请求的上下文中
		// 	// Keys is a key/value pair exclusively for the context of each request.
		// Keys map[string]any
		// Gin 采用一个map对象来维护中间传递的数据
		// context.WithValue()
		ctx.Set(token.GIN_TOKEN_KEY_NAME, tk)
		ctx.Next()
	}
}
