package impl_test

import (
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()

	// req.Role = 1
	// req.Label =
	ins, err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
