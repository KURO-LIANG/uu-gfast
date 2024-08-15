/*
* @desc:缓存处理
* @company:xxxx
* @Author: KURO
* @Date:   2023/8/223/9 11:15
 */

package service

import (
	"github.com/kuro-liang/gfast-cache/cache"
)

type ICache interface {
	cache.IGCache
}

var c ICache

func Cache() ICache {
	if c == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return c
}

func RegisterCache(che ICache) {
	c = che
}
