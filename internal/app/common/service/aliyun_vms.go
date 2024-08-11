package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"uu-gfast/internal/app/common/utils"
)

type IAliyunVms interface {
	// SendVoiceNotice 发送设备报警语音通知
	SendVoiceNotice(ctx context.Context, mobile string, content string, name string) (err error)
}

type aliyunVmsImpl struct {
}

var (
	aliyunVmsService = aliyunVmsImpl{}
)

func AliyunVms() IAliyunVms {
	return &aliyunVmsService
}

// SendVoiceNotice 发送设备报警语音通知
func (s *aliyunVmsImpl) SendVoiceNotice(ctx context.Context, mobile string, content string, name string) (err error) {
	if !utils.CheckMobile(mobile) {
		return gerror.New("手机号码不正确")
	}
	templateCode := g.Cfg().MustGet(ctx, "sms.aLiYun.template.deviceNoticeByVms").String()
	err = utils.SendVoiceNotice(mobile, templateCode, map[string]interface{}{"name": name, "msg": content})
	return
}
