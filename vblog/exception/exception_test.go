package exception_test

import (
	"encoding/json"
	"testing"

	"github.com/qinchi-ops/govlog/vblog/exception"
)

func CheckIsError() error {
	return exception.NewApiException(50001, "用户名或者密码不正确")
}

func TestException(t *testing.T) {
	err := CheckIsError()
	t.Log(err)
	// 怎么获取Errorcode ,断言这个接口的对象具体类型
	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.Message)
		t.Log(v.String())
	}
	// 前端想要获取一个完整的ApiException
	dj, _ := json.MarshalIndent(err, "", " ")
	t.Log(string(dj))
}
