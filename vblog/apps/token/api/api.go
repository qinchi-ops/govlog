package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/responese"
)

// 核心就是出来api请求
func NewTokenApiHandler(TokenServiceImpl token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		token: TokenServiceImpl,
	}
}

type TokenApiHandler struct {
	//依赖TokenServiceImpl
	//impl.TokenServiceImpl
	//依赖接口
	token token.Service
}

func (h *TokenApiHandler) Registry(appRouter gin.IRouter) {
	// r := gin.Default()
	// /api/v1 ---> root group
	// RouterGroup ---> gin.IRouter
	// r.Group("api").Group("v1").Group("tokens")

	// POST /api/v1/tokens/ ---> Login
	appRouter.POST("/", h.Login)
	appRouter.DELETE("/", h.Logout)

}

// 颁发令牌
// Body {} ---> Token {}
// Gin Handler http Request http Response
func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1.获取HTTP请求<---->IssueTokenRequest

	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		responese.Failed(err, c)
		// c.JSON(http.StatusBadRequest, err)
		return
	}
	//  标准的 http request
	// c.Request
	// Bind BindJSON
	//
	// io.ReadAll(c.Request.Body)
	// json.Unmarshal()
	// URL 参数
	// c.Query()
	// c.Request.URL.Query().Get("page_size")
	// 路径参数
	// /vblog/api/v1/users/1 URL
	// /vblog/api/v1/users/:id   id --> 1
	// c.Param("id")
	// Header参数
	// c.Request.Header.Get("Authtication")
	// c.GetHeader("Authtication")

	// 2. 业务处理
	tk, err := h.token.IssueToken(c.Request.Context(), req)
	if err != nil {
		// 不要用我们自定义的异常 "xxxx"
		// 破坏了接口返回的数据
		// c.JSON(http.StatusBadRequest, err)
		responese.Failed(err, c)
		return
	}
	// 3. 返回结果
	//
	// 线返回数据  c.Abort() 中断请求, 不然请求继续执行后面逻辑
	// c.JSON()
	// json.Marshal(data)
	// c.Writer.Write(dj)

	// // 这种返回的响应头
	// c.Writer.Header().Set()

	// // 返回HTTP协议的状态码
	// c.Writer.WriteHeader()

	// // 返回Body
	// c.Writer.Write()
	// 做个数据格式封装 YAML, XML, ....
	// c.JSON()
	// c.YAML()

	// 通过SetCookie发Cookie 设置到浏览器中, 让浏览器发送请求的时候都代码
	// 你使用的http客户端 支不支持  Set-Cookie头

	c.SetCookie(
		token.COOKIE_TOKEN_KEY,
		tk.AccessToken,
		tk.RefreshTokenExpiredAt,
		"/",
		conf.C().Application.Domain,
		false,
		false,
	)
	responese.Success(tk, c)
	// c.JSON(http.StatusOK, tk)
}

func (h *TokenApiHandler) Logout(c *gin.Context) {
	// 1. 获取HTTP 请求 <---> IssueTokenRequest
	// DELETE 方法 一般情况不带Body
	// 敏感信息放Header或者Body
	ak, err := c.Cookie(token.COOKIE_TOKEN_KEY)
	if err != nil {
		responese.Failed(err, c)
		return
	}
	rt := c.GetHeader(token.REFRESH_TOKEN_KEY)
	req := token.NewRovolkTokenRequest(ak, rt)
	tk, err := h.token.RovolkToken(c.Request.Context(), req)
	if err != nil {
		responese.Failed(err, c)
		return
	}
	responese.Success(tk, c)

}
