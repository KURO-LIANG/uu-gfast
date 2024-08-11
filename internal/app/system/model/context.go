/*
* @desc:context-model
* @company:深圳慢云智能科技有限公司
* @Author: KURO
* @Date:   2023/8/223/16 14:45
 */

package model

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*LoginUserRes
}
