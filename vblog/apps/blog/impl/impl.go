package impl

import (
	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller.Registry(blog.AppName, &BlogServiceImpl{})
}

func (i *BlogServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetDB()
	return nil
}

// 需要资源
// 需要数据库操作
type BlogServiceImpl struct {
	// db conn 共享对象
	// mysql host port  ....
	db *gorm.DB
}
