package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/common"
	"github.com/qinchi-ops/govlog/vblog/exception"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/qinchi-ops/govlog/vblog/middleware"
	"github.com/qinchi-ops/govlog/vblog/responese"
)

func (h *BlogApiHandler) Registry(appRouter gin.IRouter) {
	//不需要鉴权，公开访问
	appRouter.GET("/", h.QueryBlog)
	appRouter.GET("/:id", h.DescribeBlog)

	//需要鉴权，变更认证
	appRouter.Use(middleware.Auth)
	appRouter.POST("/", middleware.RequireRole(user.ROLE_AUTHOR), h.CreateBlog)
	appRouter.PUT("/:id", middleware.RequireRole(user.ROLE_AUTHOR), h.PutUpdateBlog)
	appRouter.PATCH("/:id", middleware.RequireRole(user.ROLE_AUTHOR), h.PatchUpdateBlog)
	appRouter.POST("/:id/status", middleware.RequireRole(user.ROLE_AUTHOR), h.UpdateBlogStatus)
	appRouter.DELETE("/:id", middleware.RequireRole(user.ROLE_AUTHOR), h.DeleteBlog)

}

/*
	//文章查看列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	//文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//文章创建
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	//文章更新
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//文章发布
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
	//文章删除
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
*/

func (h *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	// 1.获取用户登陆
	req := blog.NewQueryBlogRequest()
	req.PageRequest = common.NewPageRequestFromGin(ctx)
	req.KeyWords = ctx.Query("keywords")

	// 2. 业务处理
	set, err := h.svc.QueryBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(set, ctx)
}

func (h *BlogApiHandler) DescribeBlog(ctx *gin.Context) {
	tk, _ := ctx.Cookie(token.COOKIE_TOKEN_KEY)
	ioc.Controller.Get(token.AppName).(token.Service).ValidateToken(ctx.Request.Context(), token.NewValidateTokenRequest(tk))
	// 1.获取用户登陆
	req := blog.NewDescribeBlogRequest(ctx.Param("id"))

	// 2. 业务处理
	ins, err := h.svc.DescribeBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}

func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	// 1.获取用户登陆
	req := blog.NewCreateBlogRequest()
	if err := ctx.Bind(req); err != nil {
		responese.Failed(exception.ErrValidateFailed(err.Error()), ctx)
	}

	// 补充上下文中注入的 中间数据
	if v, ok := ctx.Get(token.GIN_TOKEN_KEY_NAME); ok {
		req.CreateBy = v.(*token.Token).UserName
	}
	// 2. 业务处理
	ins, err := h.svc.CreateBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}
func (h *BlogApiHandler) PutUpdateBlog(ctx *gin.Context) {

	// 1.获取用户登陆
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT
	//body
	if err := ctx.Bind(req.CreateBlogRequest); err != nil {
		responese.Failed(exception.ErrValidateFailed(err.Error()), ctx)
	}
	// 2. 业务处理
	ins, err := h.svc.UpdateBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}
func (h *BlogApiHandler) PatchUpdateBlog(ctx *gin.Context) {
	// 1.获取用户登陆
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH
	//body
	if err := ctx.Bind(req.CreateBlogRequest); err != nil {
		responese.Failed(exception.ErrValidateFailed(err.Error()), ctx)
	}

	// 2. 业务处理
	ins, err := h.svc.UpdateBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}
func (h *BlogApiHandler) UpdateBlogStatus(ctx *gin.Context) {
	// 1.获取用户登陆
	req := blog.NewUpdateBlogStatusRequest(ctx.Param("id"))

	//body
	if err := ctx.Bind(req.ChangedBlogStatusRequest); err != nil {
		responese.Failed(exception.ErrValidateFailed(err.Error()), ctx)
	}
	// req.Status =
	// 2. 业务处理
	ins, err := h.svc.UpdateBlogStatus(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}
func (h *BlogApiHandler) DeleteBlog(ctx *gin.Context) {
	// 1.获取用户登陆
	req := blog.NewDeleteBlogRequest(ctx.Param("id"))

	// 2. 业务处理
	ins, err := h.svc.DeleteBlog(ctx, req)
	if err != nil {
		responese.Failed(err, ctx)
		return
	}

	// 3. 返回结果
	responese.Success(ins, ctx)

}
