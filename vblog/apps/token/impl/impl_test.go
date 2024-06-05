package impl_test

import (
	"context"

	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/apps/token/impl"
	user "github.com/qinchi-ops/govlog/vblog/apps/user/impl"
)

var (
	// 声明被测试的对象
	serviceImpl token.Service
	ctx         = context.Background()
)

func init() {
	//使用构造函数
	// serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())
	serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())
}

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123")
	tk, err := serviceImpl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)

}

func TestRovolkToken(t *testing.T) {
	req := token.NewRovolkTokenRequest("cpfh8qnnl5306htovqv0", "cpfh8qnnl5306htovqvg")
	tk, err := serviceImpl.RovolkToken(ctx, req)
	if err != nil {
		t.Fatal(err)

	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cpfr22fnl533ri0eu9m0")
	tk, err := serviceImpl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)

	}
	t.Log(tk)
}
