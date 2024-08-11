// ==========================================================================
// 日期：2023-09-18 16:18:25
// 生成人：agan<960236576@qq.com>
// 功能：行政区域省市区县 service
// ==========================================================================

package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"uu-gfast/api/v1/base"
	"uu-gfast/internal/app/base/dao"
	"uu-gfast/internal/app/base/model/entity"
	"uu-gfast/library/liberr"
)

type IRegion interface {
	List(ctx context.Context, req *base.RegionSearchReq) (total int, list []*entity.Region, err error)
	Add(ctx context.Context, region *entity.Region) (err error)
	GetInfoById(ctx context.Context, Id uint64) (region *entity.Region, err error)
	Edit(ctx context.Context, region *entity.Region) (err error)
	DeleteByIds(ctx context.Context, ids []uint64) (err error)
}
type regionImpl struct {
}

var (
	regionService = regionImpl{}
)

func Region() IRegion {
	return &regionService
}

// List
func (s *regionImpl) List(ctx context.Context, req *base.RegionSearchReq) (total int, list []*entity.Region, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Region.Ctx(ctx)
		if req.ParentId != 0 {
			m = m.Where(dao.Region.Columns().ParentId+" = ?", req.ParentId)
		}
		if req.RegionName != "" {
			m = m.Where(dao.Region.Columns().RegionName+" like ?", "%"+req.RegionName+"%")
		}
		if req.MergerName != "" {
			m = m.Where(dao.Region.Columns().MergerName+" like ?", "%"+req.MergerName+"%")
		}
		if req.ShortName != "" {
			m = m.Where(dao.Region.Columns().ShortName+" like ?", "%"+req.ShortName+"%")
		}
		if req.MergerShortName != "" {
			m = m.Where(dao.Region.Columns().MergerShortName+" like ?", "%"+req.MergerShortName+"%")
		}
		if req.Level != 0 {
			m = m.Where(dao.Region.Columns().Level+" = ?", req.Level)
		}
		if req.CityCode != "" {
			m = m.Where(dao.Region.Columns().CityCode+" = ?", req.CityCode)
		}
		if req.ZipCode != "" {
			m = m.Where(dao.Region.Columns().ZipCode+" = ?", req.ZipCode)
		}
		if req.FullPinyin != "" {
			m = m.Where(dao.Region.Columns().FullPinyin+" = ?", req.FullPinyin)
		}
		if req.SimplifiedPinyin != "" {
			m = m.Where(dao.Region.Columns().SimplifiedPinyin+" = ?", req.SimplifiedPinyin)
		}
		if req.FirstChar != "" {
			m = m.Where(dao.Region.Columns().FirstChar+" = ?", req.FirstChar)
		}
		if req.Longitude != "" {
			m = m.Where(dao.Region.Columns().Longitude+" = ?", req.Longitude)
		}
		if req.Latitude != "" {
			m = m.Where(dao.Region.Columns().Latitude+" = ?", req.Latitude)
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "查询总数失败")
		err = m.Fields(dao.Region.Columns()).Page(req.PageNum, req.PageSize).Scan(&list)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

// GetInfoById 通过id获取
func (s *regionImpl) GetInfoById(ctx context.Context, id uint64) (region *entity.Region, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Region.Ctx(ctx).Where(dao.Region.Columns().Id+" = ? ", id).Scan(&region)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

// Add 添加
func (s *regionImpl) Add(ctx context.Context, region *entity.Region) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Region.Ctx(ctx).Insert(region)
		liberr.ErrIsNil(ctx, err, "新增数据失败")
	})
	return
}

// Edit 修改
func (s *regionImpl) Edit(ctx context.Context, region *entity.Region) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		d, e1 := s.GetInfoById(ctx, region.Id)
		if e1 != nil {
			liberr.ErrIsNil(ctx, e1, "修改信息失败")
			return
		}
		if d == nil {
			liberr.ErrIsNil(ctx, gerror.New("要修改的数据不存在"))
			return
		}
		_, err = dao.Region.Ctx(ctx).WherePri(region.Id).Update(region)
		liberr.ErrIsNil(ctx, err, "修改信息失败")
	})
	return
}

// DeleteByIds 删除
func (s *regionImpl) DeleteByIds(ctx context.Context, ids []uint64) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Region.Ctx(ctx).Where(dao.Region.Columns().Id+" in (?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
