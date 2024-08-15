package controller

import (
	"context"
	"uu-gfast/api/v1/common"
	"uu-gfast/api/v1/wechat"
	"uu-gfast/internal/app/wechat/service"
)

type wechatLoginController struct {
}

var WechatLogin = new(wechatLoginController)

func (c *wechatLoginController) WechatOpenidLogin(ctx context.Context, req *wechat.OpenIdLoginReq) (res *wechat.LoginRes, err error) {
	res = new(wechat.LoginRes)
	res.UserInfo, res.Token, res.MaOpenId, res.UnionId, res.ExpireTime, err = service.WechatLogin().WechatGetOpenid(ctx, req.Code)
	return
}
func (c *wechatLoginController) WechatPhoneLogin(ctx context.Context, req *wechat.PhoneLoginReq) (res *wechat.LoginRes, err error) {
	res = new(wechat.LoginRes)
	res.UserInfo, res.Token, res.ExpireTime, err = service.WechatLogin().WechatPhoneLogin(ctx, req.Code, req.MiniOpenId, req.Unionid, req.MiniAppId, req.NickName, req.Avatar)
	return
}

// UserVerifyInfo 用户实名认证信息
func (c *wechatLoginController) UserVerifyInfo(ctx context.Context, req *wechat.UserVerifyInfoReq) (res *wechat.UserVerifyInfoRes, err error) {
	res = new(wechat.UserVerifyInfoRes)
	res.VerifyInfo, err = service.WechatLogin().GetUserVerify(ctx, req.Uid)
	return
}

// UserVerifyCommit 提交实名认证
func (c *wechatLoginController) UserVerifyCommit(ctx context.Context, req *wechat.UserVerifyCommitReq) (res *common.NoneRes, err error) {
	err = service.WechatLogin().UserVerifyCommit(ctx, req.Uid, req.RelName, req.CredentialType, req.CredentialCode, req.UserImgUrl, req.IdCardFrontImgUrl, req.IdCardBackImgUrl)
	return
}
