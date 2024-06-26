package impl_test

import (
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "Go全站开发"
	req.Author = "Test"
	req.Summary = "文章摘要信息"
	req.Content = "MD内容填充"
	// req.CreateBy = "owner"

	// req.Role = 1
	// req.Label =
	ins, err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()

	ins, err := serviceImpl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.String())
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest("1")

	ins, err := serviceImpl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.String())
}
