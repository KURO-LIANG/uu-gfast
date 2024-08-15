/*
* @desc:缓存处理
* @company:xxxx
* @Author: KURO<clarence_liang@163.com>
* @Date:   2023/2/1 18:12
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "uu-gfast/api/v1/common"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"缓存管理" method:"delete" summary:"清除缓存"`
	commonApi.Author
}

type CacheRemoveRes struct {
	commonApi.NoneRes
}
