/*
* @desc:字典数据管理
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/223/18 11:57
 */

package controller

import (
	"context"
	"uu-gfast/api/v1/system"
	commonService "uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/system/service"
)

var DictData = dictDataController{}

type dictDataController struct {
}

// GetDictData 获取字典数据
func (c *dictDataController) GetDictData(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error) {
	res, err = commonService.SysDictData().GetDictWithDataByType(ctx, req)
	return
}

// List 获取字典数据列表
func (c *dictDataController) List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error) {
	res, err = commonService.SysDictData().List(ctx, req)
	return
}

// Add 添加字典数据
func (c *dictDataController) Add(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error) {
	err = commonService.SysDictData().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取对应的字典数据
func (c *dictDataController) Get(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error) {
	res, err = commonService.SysDictData().Get(ctx, req.DictCode)
	return
}

// Edit 修改字典数据
func (c *dictDataController) Edit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error) {
	err = commonService.SysDictData().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

func (c *dictDataController) Delete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error) {
	err = commonService.SysDictData().Delete(ctx, req.Ids)
	return
}
