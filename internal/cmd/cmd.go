package cmd

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"uu-gfast/internal/consts"
	"uu-gfast/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
			g.Log().Info(ctx, gbase64.MustDecodeString(consts.Logo), "Version:", consts.Version)
			sAdmin := g.Server("admin")
			sAdmin.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
			})
			err = sAdmin.Start()
			if err != nil {
				return err
			}
			enhanceOpenAPIDoc(sAdmin)

			// 微信小程序
			wechatApp := g.Server("wechatApp")
			wechatApp.Group("/", func(group *ghttp.RouterGroup) {
				router.BindWechatController(group)
			})

			err = wechatApp.Start()
			if err != nil {
				return err
			}
			enhanceOpenAPIDoc(wechatApp)

			g.Wait()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
}
