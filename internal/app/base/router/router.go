package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	baseController "uu-gfast/internal/app/base/controller"
	commonService "uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/system/service"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/base", func(group *ghttp.RouterGroup) {
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		// 不想要登录的路由
		//行政区域
		group.Bind(baseController.Region)
		// 登录鉴权中间件，下方的路由都会鉴权
		err := service.GfToken().Middleware(group)
		if err != nil {
			return
		}

		// 需要登录鉴权的业务路由

	})
}
