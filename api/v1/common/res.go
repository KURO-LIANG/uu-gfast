/*
* @desc:返回响应公共参数
* @company:深圳慢云智能科技有限公司
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/2210/27 16:30
 */

package common

import "github.com/gogf/gf/v2/frame/g"

// NoneRes 不响应任何数据
type NoneRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
