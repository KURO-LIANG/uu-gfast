package wechat

import (
	"github.com/gogf/gf/v2/frame/g"
	"uu-gfast/internal/app/base/model/entity"
)

type OpenIdLoginReq struct {
	g.Meta `path:"/login" tags:"微信小程序" method:"GET" summary:"自动登录"`
	Code   string `json:"code" v:"required#授权码不能为空"  description:"授权码"`
}

type PhoneLoginReq struct {
	g.Meta     `path:"/login/phone" tags:"微信小程序" method:"POST" summary:"手机号授权登录"`
	Code       string `json:"code" v:"required#授权码不能为空"  description:"授权码"`
	MiniOpenId string `json:"miniOpenId" v:"required#miniOpenId不能为空" description:"miniOpenId"`
	Unionid    string `json:"unionid" description:"unionid"`
	MiniAppId  string `json:"miniAppId" v:"required#小程序APPID不能为空"  description:"小程序APPID"`
	NickName   string `json:"nickName" description:"用户昵称"` // 用户昵称
	Avatar     string `json:"avatar" description:"用户头像"`   // 用户头像
}

type LoginRes struct {
	g.Meta     `mime:"application/json"`
	UserInfo   *entity.BaseUserInfo `json:"userInfo" description:"用户信息"`
	Token      string               `json:"token" description:"token"`
	ExpireTime int64                `json:"expireTime" description:"过期时间"`
	MaOpenId   string               `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId    string               `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
}

// UserVerifyInfoReq 指定用户实名认证信息
type UserVerifyInfoReq struct {
	g.Meta `path:"/login/verify/info" tags:"微信用户" method:"get" summary:"获取用户实名认证"`
	Uid    string `p:"uid" v:"required#用户ID不能为空" dc:"uid"` //用户ID
}

// UserVerifyInfoRes 指定用户实名认证信息返回
type UserVerifyInfoRes struct {
	g.Meta     `mime:"application/json"`
	VerifyInfo *entity.BaseUserVerifyInfo `json:"verifyInfo" description:"实名认证信息"`
}

// UserVerifyCommitReq 提交实名认证
type UserVerifyCommitReq struct {
	g.Meta            `path:"/login/verify/commit" tags:"微信用户" method:"post" summary:"提交实名认证"`
	Uid               string `p:"uid" v:"required#用户ID不能为空" dc:"uid"`       //用户ID
	RelName           string `p:"relName" v:"required#姓名不能为空" dc:"relName"` //姓名
	CredentialType    int    `p:"credentialType" v:"required#证件类型不能为空" dc:"credentialType"`
	CredentialCode    string `p:"credentialCode" v:"required#证件号码未输入" dc:"证件号码"`
	UserImgUrl        string `p:"userImgUrl"  dc:"用户照片"`
	IdCardFrontImgUrl string `p:"idCardFrontImgUrl"  dc:"证件正面照片"`
	IdCardBackImgUrl  string `p:"idCardBackImgUrl"  dc:"证件背面照片"`
}
