// 功能：微信用户 do
package dao

import (
	"uu-gfast/internal/app/base/dao/internal"
)

// internalBaseUserDao is internal type for wrapping internal DAO implements.
type internalBaseUserDao = *internal.BaseUserDao

// baseUserDao is the data access object for table base_user.
// You can define custom methods on it to extend its functionality as you wish.
type baseUserDao struct {
	internalBaseUserDao
}

var (
	// BaseUser is globally public accessible object for table base_user operations.
	BaseUser = baseUserDao{
		internal.NewBaseUserDao(),
	}
)

// Fill with you ideas below.
