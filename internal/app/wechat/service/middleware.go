package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	commonService "uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/wechat/model"
)

type IMiddleware interface {
	CtxInit(r *ghttp.Request)
}

type middlewareImpl struct{}

var middleService = middlewareImpl{}

func Middleware() IMiddleware {
	return &middleService
}

// CtxInit 自定义上下文对象
func (s *middlewareImpl) CtxInit(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := commonService.WechatTokenInstance().ParseToken(r)
	if err != nil {
		r.Response.WriteJson(map[string]interface{}{
			"code":   401,
			"notice": err.Error(),
		})
		return
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
