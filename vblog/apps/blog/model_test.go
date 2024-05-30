package blog_test

import (
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
)

func Test_BlogString(t *testing.T) {
	ins := blog.NewBlog()
	//失败,退出测试
	// t.Fatal()
	//打印数据
	t.Log(ins.ChangedBlogStatusRequest)

}

// func Test_CreateBlogRequestString(t *testing.T) {
// 	ins := blog.NewBlog()
// 	//失败,退出测试
// 	// t.Fatal()
// 	//打印数据
// 	t.Log(ins)

// }
