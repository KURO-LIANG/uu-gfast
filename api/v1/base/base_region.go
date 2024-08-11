// ==========================================================================
// 日期：2023-09-18 16:18:25
// 生成人：agan<960236576@qq.com>
// 功能：行政区域省市区县 接口
// ==========================================================================

package base

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "uu-gfast/api/v1/common"
	"uu-gfast/internal/app/base/model/entity"
)

// RegionSearchReq 查询列表
type RegionSearchReq struct {
	g.Meta           `path:"/region/list" tags:"行政区域省市区县管理" method:"get" summary:"查询列表"`
	ParentId         uint   `json:"parentId" description:"父ID"` //父ID
	RegionName       string `json:"regionName"`                 //名称
	MergerName       string `json:"mergerName"`                 //全称
	ShortName        string `json:"shortName"`                  //简称
	MergerShortName  string `json:"mergerShortName"`            //简称合并
	Level            int    `json:"level"`                      //层级，1是省份，2是城市，3是区县
	CityCode         string `json:"cityCode"`                   //城市代码
	ZipCode          string `json:"zipCode"`                    //邮编号码
	FullPinyin       string `json:"fullPinyin"`                 //全拼
	SimplifiedPinyin string `json:"simplifiedPinyin"`           //简拼
	FirstChar        string `json:"firstChar"`                  //第一个字
	Longitude        string `json:"longitude"`                  //纬度
	Latitude         string `json:"latitude"`                   //经度
	commonApi.PageReq
}

// RegionSearchRes 查询列表返回
type RegionSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Region `json:"list"`
	commonApi.ListRes
}

// RegionAddReq 新增
type RegionAddReq struct {
	g.Meta           `path:"/region/add" tags:"行政区域省市区县管理" method:"post" summary:"新增"`
	ParentId         uint   `json:"parentId" v:"required#父ID不能为空"  dc:"父ID"`                          //父ID
	RegionName       string `json:"regionName" v:"required#名称不能为空"  dc:"名称"`                          //名称
	MergerName       string `json:"mergerName" v:"required#全称不能为空"  dc:"全称"`                          //全称
	ShortName        string `json:"shortName" v:"required#简称不能为空"  dc:"简称"`                           //简称
	MergerShortName  string `json:"mergerShortName" v:"required#简称合并不能为空"  dc:"简称合并"`                 //简称合并
	Level            int    `json:"level" v:"required#层级，1是省份，2是城市，3是区县不能为空"  dc:"层级，1是省份，2是城市，3是区县"` //层级，1是省份，2是城市，3是区县
	CityCode         string `json:"cityCode" v:"required#城市代码不能为空"  dc:"城市代码"`                        //城市代码
	ZipCode          string `json:"zipCode" v:"required#邮编号码不能为空"  dc:"邮编号码"`                         //邮编号码
	FullPinyin       string `json:"fullPinyin" v:"required#全拼不能为空"  dc:"全拼"`                          //全拼
	SimplifiedPinyin string `json:"simplifiedPinyin" v:"required#简拼不能为空"  dc:"简拼"`                    //简拼
	FirstChar        string `json:"firstChar" v:"required#第一个字不能为空"  dc:"第一个字"`                       //第一个字
	Longitude        string `json:"longitude" v:"required#纬度不能为空"  dc:"纬度"`                           //纬度
	Latitude         string `json:"latitude" v:"required#经度不能为空"  dc:"经度"`                            //经度
}

// RegionEditReq 修改
type RegionEditReq struct {
	g.Meta           `path:"/region/edit" tags:"行政区域省市区县管理" method:"put" summary:"修改"`
	Id               uint64 `json:"id" v:"required#不能为空"  dc:""`                                      //
	ParentId         uint   `json:"parentId" v:"required#父ID不能为空"  dc:"父ID"`                          //父ID
	RegionName       string `json:"regionName" v:"required#名称不能为空"  dc:"名称"`                          //名称
	MergerName       string `json:"mergerName" v:"required#全称不能为空"  dc:"全称"`                          //全称
	ShortName        string `json:"shortName" v:"required#简称不能为空"  dc:"简称"`                           //简称
	MergerShortName  string `json:"mergerShortName" v:"required#简称合并不能为空"  dc:"简称合并"`                 //简称合并
	Level            int    `json:"level" v:"required#层级，1是省份，2是城市，3是区县不能为空"  dc:"层级，1是省份，2是城市，3是区县"` //层级，1是省份，2是城市，3是区县
	CityCode         string `json:"cityCode" v:"required#城市代码不能为空"  dc:"城市代码"`                        //城市代码
	ZipCode          string `json:"zipCode" v:"required#邮编号码不能为空"  dc:"邮编号码"`                         //邮编号码
	FullPinyin       string `json:"fullPinyin" v:"required#全拼不能为空"  dc:"全拼"`                          //全拼
	SimplifiedPinyin string `json:"simplifiedPinyin" v:"required#简拼不能为空"  dc:"简拼"`                    //简拼
	FirstChar        string `json:"firstChar" v:"required#第一个字不能为空"  dc:"第一个字"`                       //第一个字
	Longitude        string `json:"longitude" v:"required#纬度不能为空"  dc:"纬度"`                           //纬度
	Latitude         string `json:"latitude" v:"required#经度不能为空"  dc:"经度"`                            //经度
}

// RegionInfoReq 获取信息请求
type RegionInfoReq struct {
	g.Meta `path:"/region/info" tags:"行政区域省市区县管理" method:"get" summary:"获取信息"`
	Id     uint64 `json:"id" v:"required#ID不能为空"`
}

// RegionInfoRes 获取信息返回
type RegionInfoRes struct {
	g.Meta `mime:"application/json"`
	Region *entity.Region `json:"region"`
}

// RegionDeleteReq 删除
type RegionDeleteReq struct {
	g.Meta `path:"/region/delete" tags:"行政区域省市区县管理" method:"delete" summary:"删除"`
	Ids    []uint64 `json:"ids" v:"required#未选择删除的ID" dc:"删除的ID"`
}
