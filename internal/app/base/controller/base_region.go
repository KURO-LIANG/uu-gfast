// ==========================================================================
// 日期：2023-09-18 16:18:25
// 生成人：agan<960236576@qq.com>
// 功能：行政区域省市区县 controller
// ==========================================================================

package controller

import (
	"context"
	"uu-gfast/api/v1/base"
	"uu-gfast/api/v1/common"
	"uu-gfast/internal/app/base/model/entity"
	"uu-gfast/internal/app/base/service"
)

type regionController struct {
}

var Region = new(regionController)

// List 列表
func (c *regionController) List(ctx context.Context, req *base.RegionSearchReq) (res *base.RegionSearchRes, err error) {
	res = new(base.RegionSearchRes)
	res.Total, res.List, err = service.Region().List(ctx, req)
	return
}

// Add 添加
func (c *regionController) Add(ctx context.Context, req *base.RegionAddReq) (res *common.NoneRes, err error) {
	err = service.Region().Add(ctx, &entity.Region{
		ParentId:         req.ParentId,
		RegionName:       req.RegionName,
		MergerName:       req.MergerName,
		ShortName:        req.ShortName,
		MergerShortName:  req.MergerShortName,
		Level:            req.Level,
		CityCode:         req.CityCode,
		ZipCode:          req.ZipCode,
		FullPinyin:       req.FullPinyin,
		SimplifiedPinyin: req.SimplifiedPinyin,
		FirstChar:        req.FirstChar,
		Longitude:        req.Longitude,
		Latitude:         req.Latitude,
	})
	return
}

// Get 获取
func (c *regionController) Get(ctx context.Context, req *base.RegionInfoReq) (res *base.RegionInfoRes, err error) {
	res = new(base.RegionInfoRes)
	res.Region, err = service.Region().GetInfoById(ctx, req.Id)
	return
}

// Edit 修改
func (c *regionController) Edit(ctx context.Context, req *base.RegionEditReq) (res *common.NoneRes, err error) {
	err = service.Region().Edit(ctx, &entity.Region{
		Id:               req.Id,
		ParentId:         req.ParentId,
		RegionName:       req.RegionName,
		MergerName:       req.MergerName,
		ShortName:        req.ShortName,
		MergerShortName:  req.MergerShortName,
		Level:            req.Level,
		CityCode:         req.CityCode,
		ZipCode:          req.ZipCode,
		FullPinyin:       req.FullPinyin,
		SimplifiedPinyin: req.SimplifiedPinyin,
		FirstChar:        req.FirstChar,
		Longitude:        req.Longitude,
		Latitude:         req.Latitude,
	})
	return
}

// Delete 删除
func (c *regionController) Delete(ctx context.Context, req *base.RegionDeleteReq) (res *common.NoneRes, err error) {
	err = service.Region().DeleteByIds(ctx, req.Ids)
	return
}
