/*
* @desc:角色管理
* @company:深圳慢云智能科技有限公司
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/223/30 9:08
 */

package controller

import (
	"context"
	"uu-gfast/api/v1/system"
	"uu-gfast/internal/app/system/service"
)

var Role = roleController{}

type roleController struct {
	BaseController
}

// List 角色列表
func (c *roleController) List(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res, err = service.SysRole().GetRoleListSearch(ctx, req)
	return
}

// GetParams 获取角色表单参数
func (c *roleController) GetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error) {
	res = new(system.RoleGetParamsRes)
	res.Menu, err = service.SysAuthRule().GetMenuList(ctx)
	return
}

// Add 添加角色信息
func (c *roleController) Add(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	err = service.SysRole().AddRole(ctx, req)
	return
}

// Get 获取角色信息
func (c *roleController) Get(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	res = new(system.RoleGetRes)
	res.Role, err = service.SysRole().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.SysRole().GetFilteredNamedPolicy(ctx, req.Id)
	return
}

// Edit 修改角色信息
func (c *roleController) Edit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.SysRole().EditRole(ctx, req)
	return
}

// Delete 删除角色
func (c *roleController) Delete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = service.SysRole().DeleteByIds(ctx, req.Ids)
	return
}
