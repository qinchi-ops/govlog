package common

import "github.com/bytedance/sonic"

type Print struct {
}

func (p *Print) String() string {
	var err error
	ret, err := sonic.MarshalString(p)
	if err != nil {
		return ""
	}
	return ret
}
