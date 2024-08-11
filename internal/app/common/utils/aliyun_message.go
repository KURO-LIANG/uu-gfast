package utils

import (
	"context"
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	dyvmsapi20170525 "github.com/alibabacloud-go/dyvmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var aliyunSmsClient *dysmsapi20170525.Client
var aliyunVmsClient *dyvmsapi20170525.Client

func init() {
	AccessKeyId := g.Cfg().MustGet(gctx.New(), "sms.aliyun.accessKeyId").String()
	accessSecret := g.Cfg().MustGet(gctx.New(), "sms.aliyun.accessSecret").String()
	endpoint := g.Cfg().MustGet(gctx.New(), "sms.aliyun.endpoint").String()
	vmsEndpoint := g.Cfg().MustGet(gctx.New(), "sms.aliyun.vmsEndpoint").String()
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &accessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	var err error
	aliyunSmsClient, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		panic(`启动失败 阿里云SMS配置错误`)
	}

	// 访问的域名
	config.Endpoint = tea.String(vmsEndpoint)
	aliyunVmsClient, err = dyvmsapi20170525.NewClient(config)
	if err != nil {
		panic(`启动失败 阿里云VMS配置错误`)
	}
}

// SendSms 发送短信
func SendSms(mobile string, templateCode string, signName string, templateParam map[string]interface{}) (err error) {
	var param string
	if templateParam != nil {
		bytes, _ := json.Marshal(templateParam)
		param = string(bytes)
	}
	result, err := aliyunSmsClient.SendSms(&dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &mobile,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &param,
	})
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	if result != nil && *result.Body.Code == `OK` {
		return nil
	} else {
		if result != nil {
			g.Log().Error(context.TODO(), fmt.Sprintf("发送短信失败:%s", *result.Body.Message))
		}
		return gerror.New(`发送短信失败`)
	}
}

// SendVoiceNotice 发送语音通知
func SendVoiceNotice(mobile string, templateCode string, templateParam map[string]interface{}) (err error) {
	var param string
	if templateParam != nil {
		bytes, _ := json.Marshal(templateParam)
		param = string(bytes)
	}
	result, err := aliyunVmsClient.SingleCallByTts(&dyvmsapi20170525.SingleCallByTtsRequest{
		CalledNumber: &mobile,
		TtsCode:      &templateCode,
		TtsParam:     &param,
		//Volume:               nil,
	})
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	g.Log().Info(context.TODO(), "语音通知请求结果：", result)
	if result != nil && *result.Body.Code == `OK` {
		return nil
	} else {
		if result != nil {
			g.Log().Error(context.TODO(), fmt.Sprintf("发送语音失败:%s", *result.Body.Message))
		}
		return gerror.New(*result.Body.Message)
	}
}
