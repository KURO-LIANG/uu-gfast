// ==========================================================================
// 日期：2023-04-18 07:17:01
// 生成人：liangqing
// 功能：微信用户 controller
// ==========================================================================

package controller

import (
	"context"
	"uu-gfast/api/v1/base"
	"uu-gfast/api/v1/common"
	"uu-gfast/internal/app/base/service"
)

type baseUserController struct {
}

var BaseUser = new(baseUserController)

// List 列表
func (c *baseUserController) List(ctx context.Context, req *base.UserSearchReq) (res *base.UserSearchRes, err error) {
	res = new(base.UserSearchRes)
	res.Total, res.List, err = service.BaseUser().List(ctx, req)
	return
}

// UserVerifyAudit 用户校验审核
func (c *baseUserController) UserVerifyAudit(ctx context.Context, req *base.UserVerifyAuditReq) (res *common.NoneRes, err error) {
	res = new(common.NoneRes)
	err = service.BaseUser().UserVerifyAudit(ctx, req.UserId, req.VerifyState, req.AuditRemark)
	return
}
