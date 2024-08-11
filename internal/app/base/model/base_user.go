// 功能：用户

package model

type BaseUserInfo struct {
	Id       uint64 `json:"id" description:"id"`              // id
	NickName string `json:"nickName" description:"用户昵称"`      // 用户昵称
	Avatar   string `json:"avatar" description:"用户头像"`        // 用户头像
	Phone    string `json:"phone" description:"手机号"`          // 手机号
	MaOpenId string `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId  string `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
}
