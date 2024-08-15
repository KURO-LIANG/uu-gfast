// ==========================================================================
// 日期：2023-04-18 07:17:01
// 生成人：liangqing
// 功能：微信用户 接口
// ==========================================================================

package wechat

import (
	"github.com/gogf/gf/v2/frame/g"
	"uu-gfast/internal/app/base/model/entity"
)

// UserInfoReq 获取信息请求
type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"微信用户" method:"get" summary:"获取用户信息"`
	Phone  string `json:"phone" description:"手机号"`
	Id     string `json:"id" description:"用户ID"`
}

// UserInfoRes 获取信息返回
type UserInfoRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *entity.BaseUserInfo `json:"userInfo"`
}

// UserInfoListReq 获取用户列表请求
type UserInfoListReq struct {
	g.Meta `path:"/user/infoList" tags:"微信用户" method:"post" summary:"获取用户列表请求"`
	Ids    []uint64 `json:"ids" v:"required#用户id列表不能为空" description:"用户ID列表"`
}

// UserInfoListRes 获取用户列表返回
type UserInfoListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.BaseUserInfo `json:"list"`
}

// UpdateUserReq 更新用户头像或昵称
type UpdateUserReq struct {
	g.Meta   `path:"/user/updateUser" tags:"微信用户" method:"post" summary:"更新用户头像或昵称"`
	NickName string `json:"nickName" description:"用户昵称"`          // 用户昵称
	Avatar   string `json:"avatar" description:"用户头像"`            // 用户头像
	Gender   string `json:"gender" description:"性别 0-未知，1-男，2-女"` // 性别
}

// SendVerCodeToOriginPhoneReq 向原手机号发送验证码
type SendVerCodeToOriginPhoneReq struct {
	g.Meta `path:"/sendOriginVerCode" tags:"微信用户" method:"post" summary:"向原手机号发送验证码"`
}

type UpdateUserPhoneReq struct {
	g.Meta     `path:"/updatePhone" tags:"微信用户" method:"post" summary:"修改手机号"`
	Phone      string `json:"phone" v:"required#手机号不能为空" description:"手机号"`
	OldVerCode string `json:"oldVerCode" description:"旧手机验证码"`
	NewVerCode string `json:"newVerCode" v:"required#新手机验证码不能为空" description:"新手机验证码"`
}
