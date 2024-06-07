package ioc_test

import (
	"testing"

	user "github.com/qinchi-ops/govlog/vblog/apps/user/impl"
	"github.com/qinchi-ops/govlog/vblog/ioc"
)

func TestRegistry(t *testing.T) {
	ioc.Controller.Registry("user", &user.UserServiceImpl{})

	t.Logf("%p", ioc.Controller.Get("user"))
}
