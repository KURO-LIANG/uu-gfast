package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"uu-gfast/internal/app/common/utils"
)

type IAliyunSms interface {
	// SendDeviceNotice 发送设备报警短信
	SendDeviceNotice(ctx context.Context, phone string, msgType string, msgContent string, time string, familyName string) (err error)
	// SendVerCode 发送验证码
	SendVerCode(ctx context.Context, phone string, templateId string, code string) (err error)
}

type aliyunSmsImpl struct {
}

var (
	aliyunSmsService = aliyunSmsImpl{}
)

func AliyunSms() IAliyunSms {
	return &aliyunSmsService
}

// SendDeviceNotice 发送设备报警短信
func (s *aliyunSmsImpl) SendDeviceNotice(ctx context.Context, phone string, msgType string, msgContent string, time string, familyName string) (err error) {
	if !utils.CheckMobile(phone) {
		return gerror.New("手机号格式错误")
	}
	templateCode := g.Cfg().MustGet(ctx, "sms.aLiYun.template.deviceNoticeBySms").String()
	signName := g.Cfg().MustGet(ctx, "sms.aLiYun.signName").String()
	err = utils.SendSms(phone, templateCode, signName, map[string]interface{}{"msgType": msgType, "msgContent": msgContent, "time": time, "familyName": familyName})
	return
}

// SendVerCode 发送验证码
func (s *aliyunSmsImpl) SendVerCode(ctx context.Context, phone string, templateId string, code string) (err error) {
	if !utils.CheckMobile(phone) {
		return gerror.New("手机号格式错误")
	}
	signName := g.Cfg().MustGet(ctx, "sms.aLiYun.signName").String()
	err = utils.SendSms(phone, templateId, signName, map[string]interface{}{"code": code})
	return
}
