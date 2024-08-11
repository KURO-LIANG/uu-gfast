package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonService "uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/gen/controller"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/tools/gen", func(group *ghttp.RouterGroup) {
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		//生成代码
		group.Bind(controller.ToolsGenTableController)
	})
}
