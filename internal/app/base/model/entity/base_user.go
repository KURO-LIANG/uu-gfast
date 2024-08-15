// 功能：用户统一账号，手机号是全平台登录的手机号，存的是用户的基本信息
package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"uu-gfast/internal/consts"
)

// BaseUser is the golang structure for table base_user.
type BaseUser struct {
	UserId              uint64                 `json:"userId" description:"用户ID"`        // 用户ID
	Uid                 string                 `json:"uid" description:"用户UUID"`         // 用户UUID
	MaOpenId            string                 `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId             string                 `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
	NickName            string                 `json:"nickName" description:"用户昵称"`      // 用户昵称
	Avatar              string                 `json:"avatar" description:"用户头像"`        // 用户头像
	Gender              int                    `json:"gender" description:"性别 0-未知，1-男，2-女"`
	Phone               string                 `json:"phone" description:"手机号"` // 手机号
	Email               string                 `json:"email" description:"邮箱"`
	RelName             string                 `json:"relName" description:"真实姓名"`
	CredentialType      int                    `json:"credentialType" description:"证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证"`
	CredentialCode      string                 `json:"credentialCode" description:"证件号码"`
	UserImgUrl          string                 `json:"userImgUrl" description:"用户照片"`
	IdCardFrontImgUrl   string                 `json:"idCardFrontImgUrl" description:"证件正面照片"`
	IdCardBackImgUrl    string                 `json:"idCardBackImgUrl" description:"证件背面照片"`
	VerifyTime          *gtime.Time            `json:"verifyTime" description:"实名认证时间"`
	VerifyState         consts.UserVerifyState `json:"verifyState" description:"实名认证状态，0-未认证，1-认证中，2-认证通过，3-认证不通过"`
	AuditTime           *gtime.Time            `json:"auditTime" description:"审核时间"`
	AuditUser           string                 `json:"auditUser" description:"审核人"`
	AuditRemark         string                 `json:"auditRemark" description:"审核意见"`
	LastLoginTime       *gtime.Time            `json:"lastLoginTime" description:"上次登录时间"`
	LastLoginIp         string                 `json:"lastLoginIp" description:"上次登录IP"`
	LastLoginInfo       string                 `json:"lastLoginInfo" description:"上次登录设备信息"`
	SubscribeNum        int                    `json:"subscribeNum" description:"关注次数"`
	SubscribeScene      string                 `json:"subscribeScene" description:"返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他"`
	SubscribeTime       *gtime.Time            `json:"subscribeTime" description:"关注时间"`
	CancelSubscribeTime *gtime.Time            `json:"cancelSubscribeTime" description:"取消关注时间"`
	QRSceneStr          string                 `json:"qrSceneStr" description:"二维码扫码场景"`
	Language            string                 `json:"language" description:"语言"`
	Subscribe           int                    `json:"subscribe" description:"是否关注公众号 0-否；1-是；"`
	LongAndLati         string                 `json:"longAndLati" description:"经纬度  经度,纬度"`
	CreatedAt           *gtime.Time            `json:"createdAt" description:"创建时间"` // 创建时间
	UpdatedAt           *gtime.Time            `json:"updatedAt" description:"修改时间"` // 修改时间
	UpdatedBy           string                 `json:"updated_by" description:"修改人"` // 修改人
	DeletedAt           *gtime.Time            `json:"deletedAt" description:"删除时间"` // 删除时间
}

type BaseUserInfo struct {
	UserId              uint64      `json:"userId" description:"用户ID"`   // 用户ID
	Uid                 string      `json:"uid" description:"用户UUID"`    // 用户UUID
	NickName            string      `json:"nickName" description:"用户昵称"` // 用户昵称
	Avatar              string      `json:"avatar" description:"用户头像"`   // 用户头像
	Gender              int         `json:"gender" description:"性别 0-未知，1-男，2-女"`
	Phone               string      `json:"phone" description:"手机号"`          // 手机号
	MaOpenId            string      `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId             string      `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
	SubscribeNum        int         `json:"subscribeNum" description:"关注次数"`
	SubscribeScene      string      `json:"subscribeScene" description:"返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他"`
	SubscribeTime       *gtime.Time `json:"subscribeTime" description:"关注时间"`
	CancelSubscribeTime *gtime.Time `json:"cancelSubscribeTime" description:"取消关注时间"`
	QRSceneStr          string      `json:"qrSceneStr" description:"二维码扫码场景"`
	Language            string      `json:"language" description:"语言"`
	Subscribe           int         `json:"subscribe" description:"是否关注公众号 0-否；1-是；"`
	LongAndLati         string      `json:"longAndLati" description:"经纬度  经度,纬度"`
	SmsNum              int         `json:"smsNum" description:"短信数量"`
	VmsNum              int         `json:"vmsNum" description:"语音数量"`
}

type BaseSearchUser struct {
	UserId            uint64                 `json:"userId" description:"用户ID"`   // 用户ID
	Uid               string                 `json:"uid" description:"用户UUID"`    // 用户UUID
	NickName          string                 `json:"nickName" description:"用户昵称"` // 用户昵称
	Avatar            string                 `json:"avatar" description:"用户头像"`   // 用户头像
	Gender            int                    `json:"gender" description:"性别 0-未知，1-男，2-女"`
	Phone             string                 `json:"phone" description:"手机号"` // 手机号
	Email             string                 `json:"email" description:"邮箱"`
	RelName           string                 `json:"relName" description:"真实姓名"`
	CredentialType    int                    `json:"credentialType" description:"证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证"`
	CredentialCode    string                 `json:"credentialCode" description:"证件号码"`
	UserImgUrl        string                 `json:"userImgUrl" description:"用户照片"`
	IdCardFrontImgUrl string                 `json:"idCardFrontImgUrl" description:"证件正面照片"`
	IdCardBackImgUrl  string                 `json:"idCardBackImgUrl" description:"证件背面照片"`
	VerifyState       consts.UserVerifyState `json:"verifyState" description:"实名认证状态，0-未认证，1-认证中，2-认证通过，3-认证不通过"`
	VerifyTime        *gtime.Time            `json:"verifyTime" description:"实名认证时间"`
	AuditTime         *gtime.Time            `json:"auditTime" description:"审核时间"`
	AuditUser         string                 `json:"auditUser" description:"审核人"`
	AuditRemark       string                 `json:"auditRemark" description:"审核意见"`
	ChannelName       string                 `json:"channelName" description:"渠道名称"` // 渠道ID
	LastLoginTime     *gtime.Time            `json:"lastLoginTime" description:"上次登录时间"`
	LastLoginIp       string                 `json:"lastLoginIp" description:"上次登录IP"`
	LastLoginInfo     string                 `json:"lastLoginInfo" description:"上次登录设备信息"`
	CreatedAt         *gtime.Time            `json:"createdAt" description:"创建时间"` // 创建时间
}

// BaseUserVerifyInfo 用户实名认证信息
type BaseUserVerifyInfo struct {
	Mobile            string                 `json:"mobile" description:"手机号"`
	RelName           string                 `json:"relName" description:"真实姓名"`
	CredentialType    int                    `json:"credentialType" description:"证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证"`
	CredentialCode    string                 `json:"credentialCode" description:"证件号码"`
	UserImgUrl        string                 `json:"userImgUrl" description:"用户照片"`
	IdCardFrontImgUrl string                 `json:"idCardFrontImgUrl" description:"证件正面照片"`
	IdCardBackImgUrl  string                 `json:"idCardBackImgUrl" description:"证件背面照片"`
	VerifyState       consts.UserVerifyState `json:"verifyState" description:"实名认证状态，0-未认证，1-认证中，2-认证中，3-认证不通过"`
}
