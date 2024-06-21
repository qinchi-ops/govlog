package blog

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/common"
)

const (
	AppName = "blogs"
)

type Service interface {
	//文章查看列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	//文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//文章创建
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	//文章更新
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//文章删除
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
}

type QueryBlogRequest struct {
	*common.PageRequest
	KeyWords string `json:"keywords"`
}

type DescribeBlogRequest struct {
	//更新模型 全量/部分更新
	UpdateMode common.UPDATE_MODE `json:"update_mode"`
	//需要更新的数据
	*CreateBlogRequest
}

type UpdateBlogRequest struct {
}

type DeleteBlogRequest struct {
	BlogId int `json:"blog_id"`
}
