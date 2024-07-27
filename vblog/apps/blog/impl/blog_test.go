package impl_test

import (
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/common"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "Go全站开发test2"
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

// UPDATE `blogs` SET `created_at`=1718952015,`updated_at`=1719399590,`title`='xxx',`author`='Test',`content`='MD内容填充',`summary`='文章摘要信息',`tags`='{}' WHERE `id` = 1
func TestPatchBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("1")
	req.UpdateMode = common.UPDATE_MODE_PATCH
	req.Title = "xxx"
	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestPutBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("1")
	req.UpdateMode = common.UPDATE_MODE_PUT
	req.Title = "更新后文章标题PUT"
	req.Author = "PUT"
	req.Content = "PUT"

	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlogStatus(t *testing.T) {
	req := blog.NewUpdateBlogStatusRequest("1")
	req.SetStatus(blog.STATUS_PUBLISH)

	ins, err := serviceImpl.UpdateBlogStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// DELETE FROM `blogs` WHERE id = '1'
func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("1")

	ins, err := serviceImpl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
