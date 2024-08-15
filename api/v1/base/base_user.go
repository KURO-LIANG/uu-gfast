// ==========================================================================
// 日期：2023-04-18 07:17:01
// 生成人：liangqing
// 功能：微信用户 接口
// ==========================================================================

package base

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "uu-gfast/api/v1/common"
	"uu-gfast/internal/app/base/model/entity"
)

// UserSearchReq 查询列表
type UserSearchReq struct {
	g.Meta      `path:"/baseUser/list/{t}" tags:"微信用户管理" method:"get" summary:"查询列表"`
	NickName    string `json:"nickName"`    //用户昵称
	Phone       string `json:"phone"`       //手机号
	T           int8   `json:"t"`           // 1-用户列表，2-审核列表
	VerifyState string `json:"verifyState"` //实名认证审核状态
	commonApi.PageReq
}

// UserSearchRes 查询列表返回
type UserSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.BaseSearchUser `json:"list"`
	commonApi.ListRes
}

// UserVerifyAuditReq 实名认证请求
type UserVerifyAuditReq struct {
	g.Meta      `path:"/baseUser/verify/audit" tags:"微信用户管理" method:"post" summary:"实名认证"`
	UserId      uint64 `json:"userId"`
	VerifyState int    `json:"verifyState"` //审核状态
	AuditRemark string `json:"auditRemark"` //审核意见
}
