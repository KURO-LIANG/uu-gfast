// ==========================================================================
// 日期：2023-04-18 07:17:01
// 生成人：liangqing
// 功能：微信用户 service
// ==========================================================================

package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"uu-gfast/api/v1/base"
	"uu-gfast/api/v1/wechat"
	"uu-gfast/internal/app/base/dao"
	"uu-gfast/internal/app/base/model/entity"
	"uu-gfast/internal/app/system/service"
	"uu-gfast/internal/consts"
	"uu-gfast/library/liberr"
)

type IBaseUser interface {
	List(ctx context.Context, req *base.UserSearchReq) (total int, list []*entity.BaseSearchUser, err error)
	Add(ctx context.Context, baseUser *entity.BaseUser) (id int64, err error)
	GetInfoById(ctx context.Context, Id uint64) (baseUser *entity.BaseUser, err error)
	GetInfoByUid(ctx context.Context, uid string) (baseUser *entity.BaseUser, err error)
	Edit(ctx context.Context, baseUser *entity.BaseUser) (err error)
	GetByMaOpenId(ctx context.Context, openid string) (baseUser *entity.BaseUser, err error)
	GetInfo(ctx context.Context, id uint64) (userInfo *entity.BaseUserInfo, err error)
	GetByMobile(ctx context.Context, mobile string) (user *entity.BaseUser, err error)
	UpdateUser(ctx context.Context, req *wechat.UpdateUserReq, userId uint64) (userInfo *entity.BaseUserInfo, err error)
	GetInfoByMobile(ctx context.Context, mobile string) (userInfo *entity.BaseUserInfo, err error)
	GetInfoList(ctx context.Context, ids []uint64) (list []*entity.BaseUserInfo, err error)
	// UserVerifyAudit 用户实名认证审核
	UserVerifyAudit(ctx context.Context, UserId uint64, verifyState int, auditRemark string) (err error)
	// GetByCredentialCode 通过证件号获取用户信息
	GetByCredentialCode(ctx context.Context, credentialCode string, uid string) (baseUser *entity.BaseUser, err error)
	// GetByUnionId 通过unionid获取用户信息
	GetByUnionId(ctx context.Context, unionid string) (user *entity.BaseUser, err error)
}
type baseUserImpl struct {
}

var (
	baseUserService = baseUserImpl{}
)

func BaseUser() IBaseUser {
	return &baseUserService
}

// List 列表
func (s *baseUserImpl) List(ctx context.Context, req *base.UserSearchReq) (total int, list []*entity.BaseSearchUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.BaseUser.Ctx(ctx).As("u").LeftJoin("base_channel c", "u.channel_id=c.id")
		if req.NickName != "" {
			m = m.Where("u.nick_name like ?", "%"+req.NickName+"%")
		}
		if req.Phone != "" {
			m = m.Where("u.phone", req.Phone)
		}
		if req.T == 2 {
			m = m.WhereGT("u.verify_state", 0)
		}
		if req.VerifyState != "" {
			m = m.Where("u.verify_state", req.VerifyState)
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "查询总数失败")
		err = m.Fields("u.user_id,u.nick_name,u.avatar,u.phone,u.email,u.rel_name,u.credential_type,u.credential_code,u.user_img_url,u.id_card_front_img_url,u.id_card_back_img_url,u.verify_state,u.verify_time,u.audit_time,u.audit_user,u.audit_remark,c.channel_name,u.last_login_time,u.last_login_ip,u.last_login_info,u.created_at").
			OrderDesc(dao.BaseUser.Columns().UserId).
			Page(req.PageNum, req.PageSize).Scan(&list)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

func (s *baseUserImpl) GetInfoList(ctx context.Context, ids []uint64) (list []*entity.BaseUserInfo, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			if len(ids) == 0 {
				liberr.ErrIsNil(ctx, gerror.New("参数有误！"))
			}
			err = dao.BaseUser.Ctx(ctx).WhereIn(dao.BaseUser.Columns().UserId, ids).Scan(&list)
			liberr.ErrIsNil(ctx, err, "查询数据失败")
		})
		return err
	})
	return
}

// GetInfoById 通过id获取
func (s *baseUserImpl) GetInfoById(ctx context.Context, id uint64) (baseUser *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().UserId, id).Scan(&baseUser)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

func (s *baseUserImpl) GetByCredentialCode(ctx context.Context, credentialCode string, uid string) (baseUser *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BaseUser.Ctx(ctx).
			Where(dao.BaseUser.Columns().CredentialCode, credentialCode).
			WhereNot(dao.BaseUser.Columns().Uid, uid).
			Scan(&baseUser)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

// GetInfoByUid 通过uid获取
func (s *baseUserImpl) GetInfoByUid(ctx context.Context, uid string) (baseUser *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().Uid, uid).Scan(&baseUser)
		liberr.ErrIsNil(ctx, err, "查询数据失败")
	})
	return
}

func (s *baseUserImpl) GetInfoByMobile(ctx context.Context, phone string) (userInfo *entity.BaseUserInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err := s.GetByMobile(ctx, phone)
		liberr.ErrIsNil(ctx, err, "用户查询失败")
		if user == nil || user.UserId == 0 {
			liberr.ErrIsNil(ctx, gerror.New("无此用户信息"))
		}
		_ = gconv.Struct(user, &userInfo)
	})
	return
}

func (s *baseUserImpl) GetInfo(ctx context.Context, id uint64) (userInfo *entity.BaseUserInfo, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			var user *entity.BaseUser
			err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().UserId, id).Scan(&user)
			liberr.ErrIsNil(ctx, err, "查询失败")
			if user == nil {
				liberr.ErrIsNil(ctx, gerror.New("无此用户"))
			}
			_ = gconv.Struct(user, &userInfo)
		})
		return err
	})
	return
}

// Add 添加
func (s *baseUserImpl) Add(ctx context.Context, baseUser *entity.BaseUser) (id int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.BaseUser.Ctx(ctx).InsertAndGetId(baseUser)
		liberr.ErrIsNil(ctx, err, "新增数据失败")
	})
	return
}

// Edit 修改
func (s *baseUserImpl) Edit(ctx context.Context, baseUser *entity.BaseUser) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.BaseUser.Ctx(ctx).WherePri(baseUser.UserId).Update(baseUser)
		liberr.ErrIsNil(ctx, err, "修改信息失败")
	})
	return
}

func (s *baseUserImpl) GetByMaOpenId(ctx context.Context, miniOpenId string) (baseUser *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().MaOpenId, miniOpenId).Scan(&baseUser)
		liberr.ErrIsNil(ctx, err, "信息查询失败")
	})
	return
}

func (s *baseUserImpl) GetByUnionId(ctx context.Context, unionid string) (baseUser *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 查询平台账户
		err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().UnionId, unionid).Scan(&baseUser)
		liberr.ErrIsNil(ctx, err, "信息查询失败")
	})
	return
}

func (s *baseUserImpl) GetByMobile(ctx context.Context, phone string) (user *entity.BaseUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BaseUser.Ctx(ctx).Where(dao.BaseUser.Columns().Phone, phone).Fields(dao.BaseUser.Columns()).Scan(&user)
		liberr.ErrIsNil(ctx, err, "查询用户失败")
	})
	return
}

func (s *baseUserImpl) UpdateUser(ctx context.Context, req *wechat.UpdateUserReq, userId uint64) (userInfo *entity.BaseUserInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if req.NickName == "" && req.Avatar == "" && req.Gender == "" {
			liberr.ErrIsNil(ctx, gerror.New("参数有误"))
		}
		user, e := s.GetInfoById(ctx, userId)
		liberr.ErrIsNil(ctx, e, "查询用户信息失败")
		if user == nil {
			liberr.ErrIsNil(ctx, gerror.New("您的信息有误"))
		}
		if req.Avatar != "" {
			user.Avatar = req.Avatar
		}
		if req.NickName != "" {
			user.NickName = req.NickName
		}
		if req.Gender != "" {
			user.Gender = gconv.Int(req.Gender)
		}
		_, err = dao.BaseUser.Ctx(ctx).OmitEmpty().WherePri(userId).Update(user)
		liberr.ErrIsNil(ctx, err, "更新信息失败")
		_ = gconv.Struct(user, &userInfo)
	})
	return
}

// UserVerifyAudit 用户实名认证审核
func (s *baseUserImpl) UserVerifyAudit(ctx context.Context, UserId uint64, verifyState int, auditRemark string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err1 := s.GetInfoById(ctx, UserId)
		liberr.ErrIsNil(ctx, err1)
		if user == nil {
			liberr.ErrIsNil(ctx, gerror.New("用户信息不存在"))
		}
		if user.VerifyState != consts.UserVerifyStateIs认证中 {
			liberr.ErrIsNil(ctx, gerror.New("用户认证状态异常"))
		}
		_, err = dao.BaseUser.Ctx(ctx).OmitEmpty().WherePri(UserId).Update(g.Map{
			dao.BaseUser.Columns().VerifyState: verifyState,
			dao.BaseUser.Columns().AuditRemark: auditRemark,
			dao.BaseUser.Columns().AuditUser:   service.Context().GetLoginUser(ctx).UserName,
			dao.BaseUser.Columns().AuditTime:   gtime.Now(),
		})
		liberr.ErrIsNil(ctx, err, "修改信息失败")
	})
	return
}
