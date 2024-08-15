package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"uu-gfast/internal/app/base/dao"
	"uu-gfast/internal/app/base/model/entity"
	baseService "uu-gfast/internal/app/base/service"
	"uu-gfast/internal/app/common/lock"
	"uu-gfast/internal/app/common/service"
	"uu-gfast/internal/app/common/utils"
	"uu-gfast/internal/consts"
	"uu-gfast/library/liberr"
)

type IWechatLogin interface {
	WechatGetOpenid(ctx context.Context, code string) (userInfo *entity.BaseUserInfo, token string, miniOpenid string, unionid string, expireTime int64, err error)
	WechatPhoneLogin(ctx context.Context, code string, miniOpenid string, unionid string, miniAppId string, nickName string, avatar string) (userInfo *entity.BaseUserInfo, token string, expireTime int64, err error)
	// GetUserVerify 获取用户指定实名认证信息
	GetUserVerify(ctx context.Context, uid string) (userVerifyInfo *entity.BaseUserVerifyInfo, err error)
	// UserVerifyCommit 用户提交实名认证
	UserVerifyCommit(ctx context.Context, uid string, relName string, credentialType int, credentialCode string, userImgUrl string, idCardFrontImgUrl string, idCardBackImgUrl string) (err error)
}

type wechatLoginImpl struct {
}

var (
	wechatLoginService = wechatLoginImpl{}
)

func WechatLogin() IWechatLogin {
	return &wechatLoginService
}

// WechatGetOpenid 通过微信小程序登录code自动登录
func (s *wechatLoginImpl) WechatGetOpenid(ctx context.Context, code string) (userInfo *entity.BaseUserInfo, token string, miniOpenid string, unionid string, expireTime int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// code换取用户openid
		miniOpenid, unionid, err = WechatApp().GetWxUserOpenid(ctx, code)
		liberr.ErrIsNil(ctx, err)
		// 自动登录 获取平台账户用户信息
		var baseUser *entity.BaseUser
		// openid查询账户信息
		baseUser, err = baseService.BaseUser().GetByMaOpenId(ctx, miniOpenid)
		liberr.ErrIsNil(ctx, err, "登录失败")
		if baseUser == nil {
			if unionid == "" {
				return
			}
			// 没有平台账户，通过unionid查询是否有其他账户
			baseUser, err = baseService.BaseUser().GetByUnionId(ctx, unionid)
			liberr.ErrIsNil(ctx, err, "登录失败")
			if baseUser == nil {
				// 没有其他平台账户
				return
			}
			// 有其他账户，更新账户
			baseUser.MaOpenId = miniOpenid
			err = baseService.BaseUser().Edit(ctx, baseUser)
			liberr.ErrIsNil(ctx, err, "用户注册失败")
		}
		if baseUser == nil {
			baseUser = new(entity.BaseUser)
			baseUser.MaOpenId = miniOpenid
			baseUser.UnionId = unionid
			baseUser.CreatedAt = gtime.Now()
		}

		// 获取登录信息
		ip, deviceInfo := utils.GetRequestInfo(ctx)
		// 更新用户信息
		baseUser.LastLoginIp = ip
		baseUser.LastLoginInfo = deviceInfo
		baseUser.LastLoginTime = gtime.Now()
		if baseUser.Uid == "" {
			baseUser.Uid = guid.S()
		}
		err = baseService.BaseUser().Edit(ctx, baseUser)
		liberr.ErrIsNil(ctx, err, "登录失败")
		_ = gconv.Struct(baseUser, &userInfo)

		// 查询该用户是否已经登录
		token, expireTime, err = service.GetTokenByUserId(ctx, userInfo.UserId)
		liberr.ErrIsNil(ctx, err, "登录失败")
		if token == "" {
			// 没有其他登录，生成TOKEN
			token, expireTime, err = service.CreateWechatToken(ctx, userInfo)
			liberr.ErrIsNil(ctx, err, "登录失败")
		}
	})
	return
}

// WechatPhoneLogin 手机号授权登录
func (s *wechatLoginImpl) WechatPhoneLogin(ctx context.Context, code string, maOpenid string, unionId string, miniAppId string, nickName string, avatar string) (userInfo *entity.BaseUserInfo, token string, expireTime int64, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			userInfo = new(entity.BaseUserInfo)
			userInfo.Phone, err = WechatApp().GetWxUserMobile(ctx, code, false)
			if err != nil {
				if err.Error() == "40001" {
					// token失效
					userInfo.Phone, err = WechatApp().GetWxUserMobile(ctx, code, true)
				}
				liberr.ErrIsNil(ctx, err, "登录失败")
			}
			// 获取登录信息
			ip, deviceInfo := utils.GetRequestInfo(ctx)
			err = lock.Lock(userInfo.Phone, func() { //加锁，防止重复插入
				// 查询该手机号是否存在账号
				var baseUser = new(entity.BaseUser)
				baseUser, err = baseService.BaseUser().GetByMobile(ctx, userInfo.Phone)
				liberr.ErrIsNil(ctx, err, "登录失败")
				if baseUser == nil {
					userA, err1 := baseService.BaseUser().GetByMaOpenId(ctx, maOpenid)
					liberr.ErrIsNil(ctx, err1, "登录失败")
					if userA != nil && userA.UserId > 0 {
						// 登陆成功
						baseUser.LastLoginIp = ip
						baseUser.LastLoginInfo = deviceInfo
						baseUser.LastLoginTime = gtime.Now()
						baseUser.Phone = userInfo.Phone
						err = baseService.BaseUser().Edit(ctx, baseUser)
						liberr.ErrIsNil(ctx, err, "登录失败")
					}
				}
				if baseUser == nil {
					// 没有账号，注册账号
					baseUser = new(entity.BaseUser)
					baseUser.Uid = guid.S() //唯一UUID
					baseUser.Phone = userInfo.Phone
					if nickName != "" {
						baseUser.NickName = nickName
					} else {
						baseUser.NickName = userInfo.Phone
					}
					baseUser.Avatar = avatar
					baseUser.CreatedAt = gtime.Now()
					baseUser.LastLoginIp = ip
					baseUser.MaOpenId = maOpenid
					baseUser.UnionId = unionId
					baseUser.LastLoginInfo = deviceInfo
					baseUser.LastLoginTime = gtime.Now()
					var id int64
					id, err = baseService.BaseUser().Add(ctx, baseUser)
					liberr.ErrIsNil(ctx, err, "用户注册失败")
					baseUser.UserId = uint64(id)
				}

				// 登陆成功
				_ = gconv.Struct(baseUser, &userInfo)
				userInfo.MaOpenId = maOpenid
				userInfo.UnionId = unionId
				token, expireTime, err = service.CreateWechatToken(ctx, userInfo)
			})
			liberr.ErrIsNil(ctx, err)
		})
		return err
	})
	return
}

// GetUserVerify 获取指定用户ID的实名认证信息
func (s *wechatLoginImpl) GetUserVerify(ctx context.Context, uid string) (userVerifyInfo *entity.BaseUserVerifyInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, _ := baseService.BaseUser().GetInfoByUid(ctx, uid)
		if user == nil {
			liberr.ErrIsNil(ctx, gerror.New("用户信息有误"))
		}
		liberr.ErrIsNil(ctx, err, `系统异常`)
		userVerifyInfo = &entity.BaseUserVerifyInfo{
			Mobile:            user.Phone,
			RelName:           user.RelName,
			CredentialType:    user.CredentialType,
			CredentialCode:    user.CredentialCode,
			UserImgUrl:        user.UserImgUrl,
			IdCardFrontImgUrl: user.IdCardFrontImgUrl,
			IdCardBackImgUrl:  user.IdCardBackImgUrl,
			VerifyState:       user.VerifyState,
		}
	})
	return
}

// UserVerifyCommit 用户提交实名认证
func (s *wechatLoginImpl) UserVerifyCommit(ctx context.Context, uid string, relName string,
	credentialType int, credentialCode string, userImgUrl string,
	idCardFrontImgUrl string, idCardBackImgUrl string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, _ := baseService.BaseUser().GetInfoByUid(ctx, uid)
		if user == nil {
			liberr.ErrIsNil(ctx, gerror.New("用户信息有误"))
		}
		if user.VerifyState == consts.UserVerifyStateIs认证中 {
			liberr.ErrIsNil(ctx, gerror.New(`信息已在认证中`))
		}
		if user.VerifyState == consts.UserVerifyStateIs认证通过 {
			liberr.ErrIsNil(ctx, gerror.New(`用户已认证通过`))
		}
		// 查询该渠道下证件号码是否重复
		baseUser, err1 := baseService.BaseUser().GetByCredentialCode(ctx, credentialCode, uid)
		liberr.ErrIsNil(ctx, err1, "校验失败")
		if baseUser != nil {
			liberr.ErrIsNil(ctx, gerror.New("该证件号已有实名认证信息！"))
		}
		_, err = dao.BaseUser.Ctx(ctx).WherePri(user.UserId).Data(g.Map{
			dao.BaseUser.Columns().RelName:           relName,
			dao.BaseUser.Columns().CredentialType:    credentialType,
			dao.BaseUser.Columns().CredentialCode:    credentialCode,
			dao.BaseUser.Columns().UserImgUrl:        userImgUrl,
			dao.BaseUser.Columns().IdCardFrontImgUrl: idCardFrontImgUrl,
			dao.BaseUser.Columns().IdCardBackImgUrl:  idCardBackImgUrl,
			dao.BaseUser.Columns().VerifyTime:        gtime.Now(),
			dao.BaseUser.Columns().VerifyState:       consts.UserVerifyStateIs认证中,
		},
		).Update()
		liberr.ErrIsNil(ctx, err, `系统异常`)
	})
	return
}
