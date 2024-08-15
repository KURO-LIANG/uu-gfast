package utils

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

func GetRequestInfo(ctx context.Context) (ip string, userAgent string) {
	request := ghttp.RequestFromCtx(ctx)
	userAgent = request.Get("User-Agent").String()
	ip = getRealIP(request)
	return
}

func getRealIP(r *ghttp.Request) string {
	// 获取 X-Forwarded-For 头部字段
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For 头部可能包含多个 IP 地址，使用逗号分隔
		ips := strings.Split(xff, ",")
		// 返回第一个 IP 地址，通常是真实客户端的 IP 地址
		return strings.TrimSpace(ips[0])
	}
	// 如果没有 X-Forwarded-For 头部，直接返回 RemoteAddr
	return strings.Split(r.RemoteAddr, ":")[0]
}
