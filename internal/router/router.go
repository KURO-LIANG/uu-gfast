/*
* @desc:路由绑定
* @company:深圳慢云智能科技有限公司
* @Author: KURO
* @Date:   2023/8/222/18 16:23
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	baseRouter "uu-gfast/internal/app/base/router"
	commonRouter "uu-gfast/internal/app/common/router"
	commonService "uu-gfast/internal/app/common/service"
	genRouter "uu-gfast/internal/app/gen/router"
	systemRouter "uu-gfast/internal/app/system/router"
	"uu-gfast/library/libRouter"
)

var R = new(Router)

type Router struct{}

// BindController 绑定访问路由
func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//跨域处理，安全起见正式环境请注释该行
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定生成代码路由
		genRouter.BindController(group)
		// 绑定后台路由
		systemRouter.R.BindController(ctx, group)
		// 绑定公共路由
		commonRouter.R.BindController(ctx, group)
		// 绑定基础公共路由
		baseRouter.BindController(group)

		//自动绑定定义的模块
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
