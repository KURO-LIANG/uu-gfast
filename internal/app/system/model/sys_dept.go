/*
* @desc:部门model
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/8/224/11 9:07
 */

package model

import "uu-gfast/internal/app/system/model/entity"

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}
