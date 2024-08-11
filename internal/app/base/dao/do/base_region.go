// 功能：行政区域省市区县 do
package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Region is the golang structure of table base_region for DAO operations like Where/Data.
type Region struct {
	g.Meta           `orm:"table:base_region, do:true"`
	Id               interface{} //
	ParentId         interface{} // 父ID
	RegionName       interface{} // 名称
	MergerName       interface{} // 全称
	ShortName        interface{} // 简称
	MergerShortName  interface{} // 简称合并
	Level            interface{} // 层级，1是省份，2是城市，3是区县
	CityCode         interface{} // 城市代码
	ZipCode          interface{} // 邮编号码
	FullPinyin       interface{} // 全拼
	SimplifiedPinyin interface{} // 简拼
	FirstChar        interface{} // 第一个字
	Longitude        interface{} // 纬度
	Latitude         interface{} // 经度
}
