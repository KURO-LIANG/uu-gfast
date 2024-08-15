package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonService "uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/wechat/controller"
	"uu-gfast/internal/app/wechat/service"
)

func BindController(group *ghttp.RouterGroup) {
	// 微信小程序
	group.Group("/mini", func(group *ghttp.RouterGroup) {
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		// 登录
		group.Bind(controller.WechatLogin)

		// 登录鉴权中间件，下方的路由都会鉴权
		commonService.WechatTokenInstance().Middleware(group)
		//绑定用户信息中间件
		group.Middleware(service.Middleware().CtxInit)
		// 需要登录鉴权的业务路由

		// 用户信息
		group.Bind(controller.BaseUser)

	})
}
