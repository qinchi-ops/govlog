package token_test

import (
	"testing"
	"time"

	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/apps/user"
)

func Test_TokenString(t *testing.T) {

	tk := token.Token{
		UserId: 1,
		Role:   user.ROLE_ADMIN,
	}
	//失败,退出测试
	// t.Fatal()
	//打印数据
	t.Log(tk.String())

}

func Test_TokenExpired(t *testing.T) {
	now := time.Now().Unix()

	tk := token.Token{
		UserId:               1,
		Role:                 user.ROLE_ADMIN,
		AccessTokenExpiredAt: 1,
		CreatedAt:            now,
	}
	//失败,退出测试
	// t.Fatal()
	//打印数据
	t.Log(tk.AccessTokenIsExpeired())

}
