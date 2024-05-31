package impl

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"gorm.io/gorm"
)

func NewUserServiceImpl() *UserServiceImpl {
	// 每个业务对象,都可能依赖到数据库
	// db =create db
	return &UserServiceImpl{}
}

// 需要资源
// 需要数据库操作
type UserServiceImpl struct {
	// db conn 共享对象
	// mysql host port  ....
	db gorm.DB
}

func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (
	*user.User, error) {
	// 1.校验请求的合法性
	// 2. 创建user对象(资源)
	// 3. user对象保持入库
	/*
	 */
	// 4. 返回保持后的user对象
	return nil, nil
}

func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	return nil, nil
}