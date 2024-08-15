package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"uu-gfast/internal/app/wechat/model"
	"uu-gfast/internal/consts"
)

type IContext interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	GetUserId(ctx context.Context) uint64
	GetLoginUser(ctx context.Context) *model.ContextUser
}

// Context 上下文管理服务
var contextService = contextServiceImpl{}

type contextServiceImpl struct{}

func Context() IContext {
	return &contextService
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextServiceImpl) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.WechatCtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *contextServiceImpl) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.WechatCtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextServiceImpl) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *contextServiceImpl) GetLoginUser(ctx context.Context) *model.ContextUser {
	context := s.Get(ctx)
	if context == nil {
		return nil
	}
	return context.User
}

// GetUserId 获取当前登录用户id
func (s *contextServiceImpl) GetUserId(ctx context.Context) uint64 {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.UserId
	}
	return 0
}

// GetMobile 获取当前登录用户手机号
func (s *contextServiceImpl) GetMobile(ctx context.Context) string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.Phone
	}
	return ``
}
