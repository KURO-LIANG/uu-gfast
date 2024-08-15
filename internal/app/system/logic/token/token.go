/*
* @desc:token功能
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/229/27 17:01
 */

package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
	"uu-gfast/internal/app/common/consts"
	commonModel "uu-gfast/internal/app/common/model"
	"uu-gfast/internal/app/system/service"
	"uu-gfast/library/liberr"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *commonModel.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	liberr.ErrIsNil(ctx, err)
	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	service.RegisterGToken(New())
}
