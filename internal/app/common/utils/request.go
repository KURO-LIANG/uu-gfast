package utils

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"net/http"
	"time"
	"uu-gfast/library/liberr"
)

// Request 封装请求 Post请求时，data可以是多参数&连接，如："name=kuro&age=18"；可以是map类型，如：g.map{"name":"kuro","age":18}；可以是json类型，如：`{"name":"kuro","age":18}`
func Request(ctx context.Context, url string, method string, data interface{}, header map[string]string) (resData map[string]interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		c := g.Client().Timeout(3 * time.Second)
		if header != nil {
			c.SetHeaderMap(header)
		}
		var r *gclient.Response
		if method == http.MethodPost {
			r, err = c.SetHeader("content-type", "application/json").Post(ctx, url, data)
		} else if method == http.MethodGet {
			r, err = c.Get(ctx, url, data)
		} else if method == http.MethodPut {
			r, err = c.SetHeader("content-type", "application/json").Put(ctx, url, data)
		} else if method == http.MethodDelete {
			r, err = c.Delete(ctx, url, data)
		} else {
			g.Log().Error(ctx, "请求方法不匹配！")
			err = gerror.New("请求方法不匹配！")
			return
		}

		if err != nil {
			g.Log().Error(ctx, "未能成功发起远程请求！", err.Error())
			return
		}
		r.RawDump()
		defer func() {
			if r != nil {
				r.Close()
			}
		}()

		str := r.ReadAllString()
		if str == "" {
			g.Log().Error(ctx, "请求失败，未能成功获得响应数据！")
			return
		}
		g.Log().Info(ctx, "<== "+r.Status+" "+str)
		err = json.Unmarshal([]byte(str), &resData)
		if err != nil {
			liberr.ErrIsNil(ctx, err, "请求失败，解析响应数据失败")
			return
		}
	})
	return
}
