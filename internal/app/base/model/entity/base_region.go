// 功能：行政区域省市区县 model
package entity

// Region is the golang structure for table base_region.
type Region struct {
	Id               uint64 `json:"id" description:""`                     //
	ParentId         uint   `json:"parentId" description:"父ID"`            // 父ID
	RegionName       string `json:"regionName" description:"名称"`           // 名称
	MergerName       string `json:"mergerName" description:"全称"`           // 全称
	ShortName        string `json:"shortName" description:"简称"`            // 简称
	MergerShortName  string `json:"mergerShortName" description:"简称合并"`    // 简称合并
	Level            int    `json:"level" description:"层级，1是省份，2是城市，3是区县"` // 层级，1是省份，2是城市，3是区县
	CityCode         string `json:"cityCode" description:"城市代码"`           // 城市代码
	ZipCode          string `json:"zipCode" description:"邮编号码"`            // 邮编号码
	FullPinyin       string `json:"fullPinyin" description:"全拼"`           // 全拼
	SimplifiedPinyin string `json:"simplifiedPinyin" description:"简拼"`     // 简拼
	FirstChar        string `json:"firstChar" description:"第一个字"`          // 第一个字
	Longitude        string `json:"longitude" description:"纬度"`            // 纬度
	Latitude         string `json:"latitude" description:"经度"`             // 经度
}
