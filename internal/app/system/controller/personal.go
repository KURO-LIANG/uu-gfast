/*
* @desc:xxxx功能描述
* @company:深圳慢云智能科技有限公司
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/2211/3 10:32
 */

package controller

import (
	"context"
	"uu-gfast/api/v1/system"
	"uu-gfast/internal/app/system/service"
)

var Personal = new(personalController)

type personalController struct {
}

func (c *personalController) GetPersonal(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error) {
	res, err = service.Personal().GetPersonalInfo(ctx, req)
	return
}

func (c *personalController) EditPersonal(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error) {
	res, err = service.Personal().EditPersonal(ctx, req)
	return
}

func (c *personalController) ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	res, err = service.Personal().ResetPwdPersonal(ctx, req)
	return
}
