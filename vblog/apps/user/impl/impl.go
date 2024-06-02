package impl

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/common"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"gorm.io/gorm"
)

func NewUserServiceImpl() *UserServiceImpl {
	// 每个业务对象,都可能依赖到数据库
	// db =create db
	return &UserServiceImpl{
		db: *conf.C().Mysql.GetDB(),
	}
}

// 需要资源
// 需要数据库操作
type UserServiceImpl struct {
	// db conn 共享对象
	// mysql host port  ....
	db gorm.DB
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	// 1.校验请求的合法性
	if err := common.Validate(in); err != nil {
		return nil, err
	}
	//不要存储用户铭文密码
	if err := in.HashPassword(); err != nil {
		return nil, err

	}
	// 2. 创建user对象(资源)
	ins := user.NewUser(in)
	// 3. user对象保持入库
	/*
		读取数据库配置
		读取数据库连接
		操作连接 保证数据
	*/
	if err := i.db.Save(ins).Error; err != nil {
		return nil, err
	}
	// 4. 返回保持后的user对象
	return ins, nil
}

func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	return nil, nil
}
