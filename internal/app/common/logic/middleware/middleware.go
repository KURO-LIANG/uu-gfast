/*
* @desc:中间件处理
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/229/28 9:08
 */

package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"uu-gfast/internal/app/common/service"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
