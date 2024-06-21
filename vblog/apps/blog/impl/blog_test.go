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
