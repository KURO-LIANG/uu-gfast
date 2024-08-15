// 功能：微信用户 internal
package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseUserDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns BaseUserColumns // columns contains all the column names of Table for convenient usage.
}
type BaseUserColumns struct {
	UserId              string //id
	Uid                 string //Uid
	NickName            string //用户昵称
	Avatar              string //用户头像
	Gender              string //用户性别
	MaOpenId            string //小程序openid
	UnionId             string //微信开放平台id
	Phone               string //手机号
	Email               string //手机号
	RelName             string //手机号
	CredentialType      string //证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证
	CredentialCode      string //证件号码
	UserImgUrl          string //用户照片
	IdCardFrontImgUrl   string //证件正面照片
	IdCardBackImgUrl    string //证件背面照片
	VerifyState         string //实名认证状态，0-未认证，1-认证中，2-认证中，3-认证不通过
	VerifyTime          string // 认证时间
	AuditTime           string // 审核时间
	AuditUser           string // 审核人
	AuditRemark         string // 审核意见
	LastLoginTime       string //最近登录时间
	LastLoginIp         string //最近登录IP
	LastLoginInfo       string //最近登录设备信息
	CreatedAt           string //创建时间
	UpdatedAt           string //修改时间
	DeletedAt           string //删除时间
	SubscribeNum        string
	SubscribeScene      string
	SubscribeTime       string
	CancelSubscribeTime string
	QRSceneStr          string
	Language            string
	Subscribe           string
	LongAndLati         string
}

// baseUserColumns holds the columns for table base_user.
var baseUserColumns = BaseUserColumns{
	UserId:              "user_id",
	Uid:                 "uid",
	NickName:            "nick_name",
	Avatar:              "avatar",
	Gender:              "gender",
	MaOpenId:            "ma_open_id",
	UnionId:             "union_id",
	Phone:               "phone",
	Email:               "email",
	RelName:             "rel_name",
	CredentialType:      "credential_type",
	CredentialCode:      "credential_code",
	UserImgUrl:          "user_img_url",
	IdCardFrontImgUrl:   "id_card_front_img_url",
	IdCardBackImgUrl:    "id_card_back_img_url",
	VerifyState:         "verify_state",
	VerifyTime:          "verify_time",
	AuditTime:           "audit_time",
	AuditUser:           "audit_user",
	AuditRemark:         "audit_remark",
	LastLoginTime:       "last_login_time",
	LastLoginIp:         "last_login_ip",
	LastLoginInfo:       "last_login_info",
	SubscribeNum:        "subscribe_num",
	SubscribeScene:      "subscribe_scene",
	SubscribeTime:       "subscribe_time",
	CancelSubscribeTime: "cancel_subscribe_time",
	QRSceneStr:          "qr_scene_str",
	Language:            "language",
	Subscribe:           "subscribe",
	LongAndLati:         "long_and_lati",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// NewBaseUserDao creates and returns a new DAO object for table data access.
func NewBaseUserDao() *BaseUserDao {
	return &BaseUserDao{
		group:   "default",
		table:   "base_user",
		columns: baseUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current do.
func (dao *BaseUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current do.
func (dao *BaseUserDao) Columns() BaseUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current do.
func (dao *BaseUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
