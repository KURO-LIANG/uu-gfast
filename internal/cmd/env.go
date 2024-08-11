package cmd

import "github.com/gogf/gf/v2/os/genv"

// IsProd 是否是正式环境
func IsProd() bool {
	return genv.Get("env", "").String() == `prod`
}

// IsTest 是否是测试环境
func IsTest() bool {
	return genv.Get("env", "").String() == `test`
}
