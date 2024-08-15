// 功能：微信用户 do
package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseUser is the golang structure of table base_user for DAO operations like Where/Data.
type BaseUser struct {
	g.Meta            `orm:"table:base_user, do:true"`
	UserId            interface{} // id
	Uid               interface{} // id
	NickName          interface{} // 用户昵称
	Avatar            interface{} // 用户头像
	Phone             interface{} // 手机号
	Email             interface{} // 邮箱
	RelName           interface{} // 真实姓名
	CredentialType    interface{} // 证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证
	CredentialCode    interface{} // 证件号码
	UserImgUrl        interface{} // 用户照片
	IdCardFrontImgUrl interface{} // 证件正面照片
	IdCardBackImgUrl  interface{} // 证件背面照片
	VerifyState       interface{} // 实名认证状态，0-未认证，1-认证中，2-认证中，3-认证不通过
	VerifyTime        interface{} // 认证时间
	AuditTime         interface{} // 审核时间
	AuditUser         interface{} // 审核人
	AuditRemark       interface{} // 审核意见
	Score             interface{} // 剩余积分
	TotalScore        interface{} // 累计积分
	LastLoginTime     *gtime.Time // 最近登录时间
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 修改时间
	DeletedAt         *gtime.Time // 删除时间
}
