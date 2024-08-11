/*
* @desc:登录日志
* @company:深圳慢云智能科技有限公司
* @Author: KURO
* @Date:   2023/8/223/8 11:43
 */

package model

// LoginLogParams 登录日志写入参数
type LoginLogParams struct {
	Status    int
	Username  string
	Ip        string
	UserAgent string
	Msg       string
	Module    string
}
