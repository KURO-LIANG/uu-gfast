package model

import "uu-gfast/internal/app/base/model/entity"

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*entity.BaseUserInfo
}
