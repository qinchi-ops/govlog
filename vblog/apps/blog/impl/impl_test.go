package impl_test

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/qinchi-ops/govlog/vblog/test"

	// 导入被测试的全部对象
	_ "github.com/qinchi-ops/govlog/vblog/apps"
)

var (
	// 声明被测试的对象
	serviceImpl blog.Service
	ctx         = context.Background()
)

// 招到对象
func init() {
	test.DevelopmentSetup()
	// serviceImpl = impl.NewUserServiceImpl()

	serviceImpl = ioc.Controller.Get(blog.AppName).(blog.Service)
}
