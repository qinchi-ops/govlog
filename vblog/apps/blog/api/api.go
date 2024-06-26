package api

import (
	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/ioc"
)

// 核心就是出来api请求
// func NewTokenApiHandler(TokenServiceImpl token.Service) *TokenApiHandler {
// func NewTokenApiHandler() *BlogApiHandler {
// 	return &BlogApiHandler{
// 		// token: TokenServiceImpl,
// 		token: ioc.Controller.Get(blog.AppName).(blog.Service),
// 	}
// }

func init() {
	ioc.Api.Registry(blog.AppName, &BlogApiHandler{})
}

type BlogApiHandler struct {
	//依赖TokenServiceImpl
	//impl.TokenServiceImpl
	//依赖接口
	svc blog.Service
}

func (h *BlogApiHandler) Init() error {
	h.svc = ioc.Controller.Get(blog.AppName).(blog.Service)
	// 把自己注册到Root Router
	// /vblog/v1/ 前置
	subRouter := conf.C().Application.GinRootRouter().Group("blogs")
	h.Registry(subRouter)

	return nil

}
