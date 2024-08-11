/*
* @desc:验证码获取
* @company:深圳慢云智能科技有限公司
* @Author: KURO
* @Date:   2023/8/223/2 17:45
 */

package controller

import (
	"context"
	"uu-gfast/api/v1/common"
	"uu-gfast/internal/app/common/service"
)

var Captcha = captchaController{}

type captchaController struct {
}

// Get 获取验证码
func (c *captchaController) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}
