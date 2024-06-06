package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/token"
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
	responese.Success(tk, c)
	// c.JSON(http.StatusOK, tk)
}
