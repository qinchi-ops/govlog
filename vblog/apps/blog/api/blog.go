package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/common"
	"github.com/qinchi-ops/govlog/vblog/responese"
)

func (h *BlogApiHandler) Registry(appRouter gin.IRouter) {
	appRouter.POST("/", h.CreateBlog)
	appRouter.GET("/", h.QueryBlog)
	appRouter.GET("/:id", h.DescribeBlog)
	appRouter.PUT("/:id", h.PutUpdateBlog)
	appRouter.PATCH("/id", h.PatchUpdateBlog)
	appRouter.PUT("/:id/status", h.UpdateBlogStatus)
	appRouter.DELETE("/:id", h.DeleteBlog)

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

}

func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {

}
func (h *BlogApiHandler) PutUpdateBlog(ctx *gin.Context) {

}
func (h *BlogApiHandler) PatchUpdateBlog(ctx *gin.Context) {

}
func (h *BlogApiHandler) UpdateBlogStatus(ctx *gin.Context) {

}
func (h *BlogApiHandler) DeleteBlog(ctx *gin.Context) {

}
