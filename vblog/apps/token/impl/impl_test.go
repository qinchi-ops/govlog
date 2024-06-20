package impl_test

import (
	"context"

	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"github.com/qinchi-ops/govlog/vblog/test"
)

var (
	// 声明被测试的对象
	serviceImpl token.Service
	ctx         = context.Background()
)

func init() {
	test.DevelopmentSetup()
	//使用构造函数
	// serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())

	//去ioc中获取被测试的业务对象
	serviceImpl = ioc.Controller.Get(token.AppName).(token.Service)
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
