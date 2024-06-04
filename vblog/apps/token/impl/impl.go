package impl

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/apps/token"
	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"gorm.io/gorm"
)

func NewTokenServiceImpl(userServiceImpl user.Service) *TokenServiceImpl {
	// 每个业务对象,都可能依赖到数据库
	// db =create db
	return &TokenServiceImpl{
		db:   *conf.C().Mysql.GetDB(),
		user: userServiceImpl,
	}
}

// 需要资源
// 需要数据库操作
type TokenServiceImpl struct {
	// db conn 共享对象
	// mysql host port  ....
	db   gorm.DB
	user user.Service
}

// 令牌颁发
func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	// 1. 查询用户对象
	queryUser := user.NewQueryUserRequest()
	queryUser.Username = in.Username

	us, err := i.user.QueryUser(ctx, queryUser)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, token.ErrAuthFailed
	}
	// //2.比对用户密码
	if err := us.Items[0].CheckPassword(in.Password); err != nil {
		return nil, token.ErrAuthFailed
	}

	return nil, nil
	// 3.颁发一个令牌(Token)

	// 4.把令牌存储在数据库里面

	// 5. 返回令牌
}

// 令牌撤销
func (i *TokenServiceImpl) RovolkToken(ctx context.Context, in *token.RovolkTokenRequest) (*token.Token, error) {
	//直接删除数据库里面存储的token
	return nil, nil
}

// 令牌校验，校验令牌合法
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	// 1. 查询出Token, Where AccessToken 来查询conf
	// 2. 判断Token是否过期， [1,先判断RefreshToken有没有过期 2，AccessTOken有没有过期 ]
	// 3. Token合法[1，是我颁发的，2，没有过期]
	//返回查询到的Token
	return nil, nil
}
