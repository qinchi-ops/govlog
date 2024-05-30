package token_test

import (
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/token"
)

func Test_TokenString(t *testing.T) {

	tk := token.Token{
		UserId: "admin",
		Role:   1,
	}
	//失败,退出测试
	// t.Fatal()
	//打印数据
	t.Log(tk.String())

}
