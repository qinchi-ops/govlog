package token

import "context"

type Service interface {
	//令牌颁发
	// 1. 调用user模块来校验用户名密码const
	// 2.  。。。。
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	//令牌撤销
	//删除令牌
	RovolkToken(context.Context, *RovolkTokenRequest) (*Token, error)
	//令牌校验，校验令牌合法
	//检查令牌是否我们颁发（我们村数据）
	//判断令牌是否过期
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
		IsMember: false,
	}
}

type IssueTokenRequest struct {
	Username string
	Password string
	//记住密码
	IsMember bool
}

func NewRovolkTokenRequest(at, rt string) *RovolkTokenRequest {
	return &RovolkTokenRequest{
		AccessToken: at,
		RefreshToen: rt,
	}
}

type RovolkTokenRequest struct {
	// 你需要撤销的令牌
	// AccessToken,RefreshToken 构成了一对username/password

	AccessToken string
	//你需要知道正确的刷新Token
	RefreshToen string
}

func NewValidateTokenRequest(at string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: at,
	}
}

type ValidateTokenRequest struct {
	// 你需要撤销的令牌
	// AccessToken,RefreshToken 构成了一对username/password
	AccessToken string
}
