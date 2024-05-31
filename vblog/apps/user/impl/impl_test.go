package impl_test

import (
	"context"
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/apps/user/impl"
)

var (
	// 声明被测试的对象
	serviceImpl user.Service
	ctx         = context.Background()
)

// 招到对象
func init() {
	serviceImpl = &impl.UserServiceImpl{}
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal()
	}
	t.Log(ins)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	ins, err := serviceImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal()
	}
	t.Log(ins)
}
