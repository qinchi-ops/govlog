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
	//文章发布
	UpdateBlogStatus(context.Context, *ChangedBlogStatusRequest) (*Blog, error)
	//文章删除
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: common.NewPageRequest(),
	}

}

type QueryBlogRequest struct {
	*common.PageRequest
	KeyWords string `json:"keywords"`
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}

type DescribeBlogRequest struct {
	BlogId string `json:"blog_id"`
}

type UpdateBlogRequest struct {
	//更新模型 全量/部分更新
	UpdateMode common.UPDATE_MODE `json:"update_mode"`
	//需要更新的数据
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	BlogId string `json:"blog_id"`
}
