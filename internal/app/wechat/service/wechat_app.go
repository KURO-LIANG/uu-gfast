package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"strconv"
	"time"
	"uu-gfast/internal/app/common/service"
	"uu-gfast/library/liberr"
)

type IWechatApp interface {
	GetWxUserOpenid(ctx context.Context, code string) (openid string, unionId string, err error)
	GetWxUserMobile(ctx context.Context, code string, refresh bool) (mobile string, err error)
	GetQRCode(ctx context.Context, chatId uint64) (qrCodeBytes []byte, err error)
}

type wechatAppImpl struct {
}

func (w wechatAppImpl) GetQRCode(ctx context.Context, chatId uint64) (qrCodeBytes []byte, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 生成小程序码的参数
		scene := strconv.FormatUint(chatId, 10)
		page := "pages/index/detail/detail"
		width := 430

		appId := g.Cfg().MustGet(ctx, "wx.appId").String()
		appSecret := g.Cfg().MustGet(ctx, "wx.appSecret").String()

		// 获取小程序 access token
		accessToken, e := getAccessToken(ctx, appId, appSecret)
		if e != nil {
			liberr.ErrIsNil(ctx, e, "获取token失败")
		}

		// 获取小程序码
		qrCodeURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", accessToken)
		qrCodeBytes, err = downloadQRCode(qrCodeURL, scene, page, width, "qrcode.jpg")
		liberr.ErrIsNil(ctx, err, "下载失败")
	})
	return
}

var (
	wechatAppService = wechatAppImpl{}
)

func WechatApp() IWechatApp {
	return &wechatAppService
}

// 下载小程序码
func downloadQRCode(url, scene, page string, width int, filename string) ([]byte, error) {
	payload := struct {
		Scene string `json:"scene"`
		Page  string `json:"page"`
		Width int    `json:"width"`
	}{
		Scene: scene,
		Page:  page,
		Width: width,
	}

	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	qrCodeBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return qrCodeBytes, nil
}

func (w wechatAppImpl) GetWxUserOpenid(ctx context.Context, code string) (openid string, unionId string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		miniAppId := g.Cfg().MustGet(ctx, "wx.miniAppId").String()
		miniAppSecret := g.Cfg().MustGet(ctx, "wx.miniAppSecret").String()
		startTime := time.Now()
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", miniAppId, miniAppSecret, code)
		g.Log().Info(ctx, fmt.Sprintf("请求小程序 jscode2session url = %s", url))
		r, err1 := g.Client().Get(ctx, url)
		if err1 != nil {
			g.Log().Error(ctx, " 请求小程序openid err ", err1)
			liberr.ErrIsNil(ctx, gerror.New("获取openid失败"))
		}
		defer func() {
			if r != nil {
				r.Close()
			}
		}()
		str := r.ReadAllString()
		if str != `` {
			var data map[string]interface{}
			err1 = json.Unmarshal([]byte(str), &data)
			if err1 != nil {
				g.Log().Error(ctx, "解析获取openid数据失败", err1)
				liberr.ErrIsNil(ctx, gerror.New("登录失败"))
				return
			}
			errMsg := data["errmsg"]
			errCode := gconv.Int(fmt.Sprintf("%1.0f", data["errcode"]))
			if errCode != 0 {
				g.Log().Error(ctx, fmt.Sprintf("获取小程序openid失败 errMsg = %s", errMsg))
				liberr.ErrIsNil(ctx, gerror.New("登录失败"))
				return
			}
			openid = gconv.String(data["openid"])
			unionId = gconv.String(data["unionid"])
			g.Log().Info(ctx, fmt.Sprintf("<== 请求结束，本次请求耗时: %s\n", time.Since(startTime)))
			return
		}
		g.Log().Error(ctx, fmt.Sprintf("获取小程序openid失败 str = %s", str))
		liberr.ErrIsNil(ctx, gerror.New("登录失败"))
		return
	})
	return
}

// getAccessToken 小程序全局唯一后台接口调用凭据，token有效期为7200s
func getAccessToken(ctx context.Context, appID string, appSecret string) (accessToken string, err error) {
	if appID == `` || appSecret == `` {
		return ``, gerror.New("小程序配置信息不正确")
	}
	v := service.Cache().Get(ctx, appID)
	if v != nil && v.String() != `` {
		g.Log().Info(ctx, "获取缓存的 access_token ")
		return v.String(), err
	}
	params := g.Map{
		"grant_type":    "client_credential",
		"appid":         appID,
		"secret":        appSecret,
		"force_refresh": false, // false 普通调用模式， access_token有效期内重复调用不会更新token，true时为强制刷新，会导致上次获取的token失效并返回新的token
	}
	// 请求接口获取小程序token
	r, err := g.Client().SetHeader("Content-Type", "application/json").
		Post(ctx, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/stable_token"), params)
	if err != nil {
		g.Log().Error(ctx, " 请求小程序获取token err ", err)
		return ``, gerror.New("登录获取token失败")
	}
	r.RawDump()
	defer func() {
		if r != nil {
			r.Close()
		}
	}()
	str := r.ReadAllString()
	if str != `` {
		var data map[string]interface{}
		err = json.Unmarshal([]byte(str), &data)
		if err != nil {
			return ``, gerror.New("解析token失败")
		}
		token := data["access_token"]
		expires := data["expires_in"]
		if token != nil && expires != nil {
			accessToken = token.(string)
			t := gconv.Int64(fmt.Sprintf("%1.0f", expires))
			//缓存access_token
			service.Cache().Set(ctx, appID, accessToken, time.Duration(t)*time.Second)
			return accessToken, nil
		}
	}
	g.Log().Error(ctx, fmt.Sprintf("请求小程序获取token 结果 = %s", str))
	return ``, gerror.New("获取token 失败")
}

// GetWxUserMobile 小程序接口获取access_token
func (w wechatAppImpl) GetWxUserMobile(ctx context.Context, code string, refresh bool) (mobile string, err error) {
	miniAppId := g.Cfg().MustGet(ctx, "wx.miniAppId").String()
	miniAppSecret := g.Cfg().MustGet(ctx, "wx.miniAppSecret").String()
	if refresh {
		service.Cache().Remove(ctx, miniAppId)
	}
	accessToken, e := getAccessToken(ctx, miniAppId, miniAppSecret)
	if e != nil {
		liberr.ErrIsNil(ctx, e)
		return
	}
	r, err := g.Client().SetHeader("content-type", "application/json").
		Post(ctx, fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", accessToken), fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		g.Log().Error(ctx, " 请求小程序获取手机号码 err ", err)
		liberr.ErrIsNil(ctx, gerror.New("登录失败"))
	}
	r.RawDump()
	defer func() {
		if r != nil {
			r.Close()
		}
	}()
	str := r.ReadAllString()
	if str != `` {
		var data map[string]interface{}
		err = json.Unmarshal([]byte(str), &data)
		if err != nil {
			return ``, gerror.New("解析获取手机号码数据失败")
		}
		errMsg := data["errmsg"]
		errCode := gconv.Int(fmt.Sprintf("%1.0f", data["errcode"]))
		phoneInfo := data["phone_info"]
		if errCode == 40001 { //token 过期重新获取
			return ``, gerror.New("40001")
		}
		if errCode != 0 {
			g.Log().Error(ctx, fmt.Sprintf("获取小程序绑定手机号码失败 errMsg = %s", errMsg))
			liberr.ErrIsNil(ctx, gerror.New("获取手机号码失败"))
		}
		info := phoneInfo.(map[string]interface{})
		return info["phoneNumber"].(string), nil
	}
	g.Log().Error(ctx, fmt.Sprintf("获取小程序绑定手机号码失败 str = %s", str))
	liberr.ErrIsNil(ctx, gerror.New("获取手机号码失败"))
	return
}
