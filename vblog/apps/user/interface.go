package user

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/common"
)

// user.Service
// 用户接口管理
// 接口定义的原则: 站在调用方(使用者)的角度来设计接口
// userServiceImpl.CreateUser(ctx, *CreateUserRequest)
type Service interface {
	//用户创建
	// 1. 用户取消了请求这么办?
	// 2. 后面要做trace,Trace ID 怎么传递
	// 3. 多个接口,需要做事物(Session),这个事物对象怎么传递进来
	// 4. CreateUser(username,password string)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	//用户查询
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryUserRequest struct {
	*common.PageRequest
}