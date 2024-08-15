package consts

const (
	OpenAPITitle       = `GFast-v3`
	OpenAPIDescription = `基于 GoFrame2.0的后台管理系统。 Enjoy 💖 `
	OpenAPIContactName = "GFast"
	OpenAPIContactUrl  = "http://www.g-fast.cn"
)

const (
	// WechatCtxKey 本地的登录信息前缀 小程序登录保存在请求参数的key
	WechatCtxKey = "UUGfastWechatContext"
)

// UserVerifyState 实名认证状态
type UserVerifyState int8

const (
	UserVerifyStateIs未认证 UserVerifyState = iota
	UserVerifyStateIs认证中
	UserVerifyStateIs认证通过
	UserVerifyStateIs认证不通过
)
