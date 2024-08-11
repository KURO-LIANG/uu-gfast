/*
* @desc:公共接口相关
* @company:深圳慢云智能科技有限公司
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/223/30 9:28
 */

package common

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `p:"dateRange"` //日期范围
	PageNum   int      `p:"pageNum"`   //当前页码
	PageSize  int      `p:"pageSize"`  //每页数
	OrderBy   string   //排序方式
}

// Author 授权信息
type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
