/*
* @desc:后台路由
* @company:xxxx
* @Author: KURO
* @Date:   2023/8/222/18 17:34
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"uu-gfast/internal/app/common/controller"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/pub", func(group *ghttp.RouterGroup) {
		group.Group("/captcha", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Captcha,
			)
		})
	})
}
func BindUploadController(group *ghttp.RouterGroup) {
	group.Group("/upload", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Upload,
		)
	})
}
