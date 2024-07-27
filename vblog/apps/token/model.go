package token

import (
	"encoding/json"
	"time"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/rs/xid"
)

func NewToken(u *user.User) *Token {
	return &Token{
		UserId:                u.Id,
		UserName:              u.Username,
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  3600,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 4,
		CreatedAt:             time.Now().Unix(),
		Role:                  u.Role,
	}
}

func DefaultToken() *Token {
	return &Token{}
}

type Token struct {
	// 该Token是颁发
	UserId int `json:"user_id" gorm:"column:user_id"`
	// 人的名称， user_name
	UserName string `json:"username" gorm:"column:username"`
	// 办法给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	// 过期时间(2h), 单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	// 刷新Token过期时间(7d)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`

	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新实现
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// 额外补充信息, gorm忽略处理
	Role user.Role `json:"role" gorm:"-"`
}

func (t *Token) String() string {
	dj, _ := json.Marshal(t)
	// fmt.Print("findis")
	return string(dj)

}

func (t *Token) IssueTime() time.Time {
	return time.Unix(t.CreatedAt, 0)

}

func (t *Token) AccessTokenDuration() time.Duration {
	return time.Duration(t.AccessTokenExpiredAt * int(time.Second))
}

func (t *Token) RefreshTokenDuration() time.Duration {
	return time.Duration(t.RefreshTokenExpiredAt * int(time.Second))
}

func (t *Token) AccessTokenIsExpired() error {
	// 过期时间: 颁发时间+过期时长
	expieredTime := t.IssueTime().Add(t.AccessTokenDuration())
	// time.Now().Sub(t)
	if time.Since(expieredTime).Seconds() > 0 {
		return ErrAccessTokenExpired
	}
	return nil
}

func (t *Token) RefreshTokenIsExpired() error {
	expieredTime := t.IssueTime().Add(t.RefreshTokenDuration())
	// time.Now().Sub(t)
	if time.Since(expieredTime).Seconds() > 0 {
		return ErrRefreshTokenExpired
	}
	return nil
}
