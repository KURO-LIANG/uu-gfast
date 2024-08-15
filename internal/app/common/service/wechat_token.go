package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sync"
	"time"
	"uu-gfast/internal/app/base/model/entity"
	"uu-gfast/internal/app/common/model"
	"uu-gfast/library/liberr"
)

type wechatToken struct {
	options *model.TokenOptions
	gT      IGfToken
	lock    *sync.Mutex
}

type userToken struct {
	UserId      uint64
	Token       string
	ExpiredTime int64
}

var wechatTokenService = &wechatToken{
	options: nil,
	gT:      nil,
	lock:    &sync.Mutex{},
}

func WechatTokenInstance() IGfToken {
	if wechatTokenService.gT == nil {
		wechatTokenService.lock.Lock()
		defer wechatTokenService.lock.Unlock()
		if wechatTokenService.gT == nil {
			ctx := gctx.New()
			err := g.Cfg().MustGet(ctx, "wechatToken").Struct(&wechatTokenService.options)
			liberr.ErrIsNil(ctx, err)
			wechatTokenService.gT = WechatToken(wechatTokenService.options)
		}
	}
	return wechatTokenService.gT
}

// WechatTokenTimeout 获取token的timeout信息
func WechatTokenTimeout() int64 {
	if wechatTokenService.options != nil {
		return wechatTokenService.options.Timeout
	}
	return 0
}

// CreateWechatToken 生成小程序的的token
func CreateWechatToken(ctx context.Context, userInfo *entity.BaseUserInfo) (token string, expireTime int64, err error) {
	//生成token
	key := gconv.String(userInfo.UserId) + "-" + gmd5.MustEncryptString(userInfo.NickName) + gmd5.MustEncryptString(userInfo.MaOpenId)
	token, err = WechatTokenInstance().GenerateToken(ctx, key, &userInfo)
	liberr.ErrIsNil(ctx, err)
	et := time.Duration(WechatTokenTimeout()-1000) * time.Second
	expireTime = gtime.Now().Add(et).TimestampMilli()
	// 创建用户id与token的关联缓存，便于后续通过用户id查询登录态
	key = fmt.Sprintf(wechatTokenService.options.CacheUserKey, userInfo.UserId)
	var userTokenVal = userToken{
		UserId:      userInfo.UserId,
		Token:       token,
		ExpiredTime: expireTime,
	}
	val, _ := json.Marshal(userTokenVal)
	Cache().Set(ctx, key, string(val), et)
	return
}

// GetTokenByUserId 获取小程序登录TOKEN信息
func GetTokenByUserId(ctx context.Context, userId uint64) (token string, expireTime int64, err error) {
	key := fmt.Sprintf(wechatTokenService.options.CacheUserKey, userId)
	if !Cache().Contains(ctx, key) {
		return
	}
	tokenVal := Cache().Get(ctx, key).String()
	var userTokenVal userToken
	_ = json.Unmarshal([]byte(tokenVal), &userTokenVal)
	token = userTokenVal.Token
	expireTime = userTokenVal.ExpiredTime
	return
}
