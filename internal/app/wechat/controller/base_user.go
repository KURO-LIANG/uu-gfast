// ==========================================================================
// 日期：2023-04-18 07:17:01
// 生成人：liangqing
// 功能：微信用户 controller
// ==========================================================================

package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"uu-gfast/api/v1/wechat"
	"uu-gfast/internal/app/base/service"
	wechatService "uu-gfast/internal/app/wechat/service"
)

type baseUserController struct {
}

var BaseUser = new(baseUserController)

// UserInfo 获取用户信息
func (c *baseUserController) UserInfo(ctx context.Context, req *wechat.UserInfoReq) (res *wechat.UserInfoRes, err error) {
	res = new(wechat.UserInfoRes)
	userId := wechatService.Context().GetUserId(ctx)
	if req.Phone != "" {
		res.UserInfo, err = service.BaseUser().GetInfoByMobile(ctx, req.Phone)
	} else if req.Id != "" {
		res.UserInfo, err = service.BaseUser().GetInfo(ctx, gconv.Uint64(req.Id))
	} else {
		res.UserInfo, err = service.BaseUser().GetInfo(ctx, userId)
	}
	return
}

// UserInfoList 获取用户列表请求
func (c *baseUserController) UserInfoList(ctx context.Context, req *wechat.UserInfoListReq) (res *wechat.UserInfoListRes, err error) {
	res = new(wechat.UserInfoListRes)
	res.List, err = service.BaseUser().GetInfoList(ctx, req.Ids)
	return
}

// UpdateUser 更新用户信息
func (c *baseUserController) UpdateUser(ctx context.Context, req *wechat.UpdateUserReq) (res *wechat.UserInfoRes, err error) {
	res = new(wechat.UserInfoRes)
	userId := wechatService.Context().GetUserId(ctx)
	res.UserInfo, err = service.BaseUser().UpdateUser(ctx, req, userId)
	return
}
