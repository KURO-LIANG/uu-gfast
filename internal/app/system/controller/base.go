/*
* @desc:system base controller
* @company:深圳慢云智能科技有限公司
* @Author: KURO
* @Date:   2023/8/223/4 18:12
 */

package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "uu-gfast/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
