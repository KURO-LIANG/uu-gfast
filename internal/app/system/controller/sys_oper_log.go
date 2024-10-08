/*
* @desc:系统后台操作日志
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/229/21 16:10
 */

package controller

import (
	"context"
	"uu-gfast/api/v1/system"
	"uu-gfast/internal/app/system/service"
)

var OperLog = new(operateLogController)

type operateLogController struct {
	BaseController
}

// List 列表
func (c *operateLogController) List(ctx context.Context, req *system.SysOperLogSearchReq) (res *system.SysOperLogSearchRes, err error) {
	res, err = service.OperateLog().List(ctx, req)
	return
}

// Get 获取操作日志
func (c *operateLogController) Get(ctx context.Context, req *system.SysOperLogGetReq) (res *system.SysOperLogGetRes, err error) {
	res = new(system.SysOperLogGetRes)
	res.SysOperLogInfoRes, err = service.OperateLog().GetByOperId(ctx, req.OperId)
	return
}

func (c *operateLogController) Delete(ctx context.Context, req *system.SysOperLogDeleteReq) (res *system.SysOperLogDeleteRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperIds)
	return
}

func (c *operateLogController) Clear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
