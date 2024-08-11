/*
* @desc:token功能
 */

package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast-token/gftoken"
	"uu-gfast/internal/app/common/consts"
	"uu-gfast/internal/app/common/model"
)

type IGfToken interface {
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	Middleware(group *ghttp.RouterGroup) error
	ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	IsLogin(r *ghttp.Request) (b bool, failed *gftoken.AuthFailed)
	GetRequestToken(r *ghttp.Request) (token string)
	RemoveToken(ctx context.Context, token string) (err error)
}

type gfTokenImpl struct {
	*gftoken.GfToken
}

var gT = gfTokenImpl{
	GfToken: gftoken.NewGfToken(),
}

var wT = gfTokenImpl{
	GfToken: gftoken.NewGfToken(),
}

// GfToken 后台token
func GfToken(options *model.TokenOptions) IGfToken {
	gT.GfToken = buildTokenInstance(options)
	return &gT
}

// WechatToken 微信token
func WechatToken(options *model.TokenOptions) IGfToken {
	wT.GfToken = buildTokenInstance(options)
	return &wT
}

// buildTokenInstance buildTokenInstance
func buildTokenInstance(options *model.TokenOptions) *gftoken.GfToken {
	var fun gftoken.OptionFunc
	if options.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return gftoken.NewGfToken(
		gftoken.WithCacheKey(options.CacheKey),
		gftoken.WithTimeout(options.Timeout),
		gftoken.WithMaxRefresh(options.MaxRefresh),
		gftoken.WithMultiLogin(options.MultiLogin),
		gftoken.WithExcludePaths(options.ExcludePaths),
		fun,
	)
}
