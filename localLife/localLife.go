package locallife

import (
	"github.com/Biely/douyinSDK/credential"
	"github.com/Biely/douyinSDK/localLife/config"
	"github.com/Biely/douyinSDK/localLife/context"
	"github.com/Biely/douyinSDK/localLife/shop"
)

type LocalLife struct {
	ctx *context.Context
}

func NewLocalLife(cfg *config.Config) *LocalLife {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &LocalLife{ctx}
}

// SetAccessTokenHandle 自定义 access_token 获取方式
func (localLife *LocalLife) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	localLife.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (localLife *LocalLife) GetContext() *context.Context {
	return localLife.ctx
}

// 门店管理接口
func (localLife *LocalLife) GetShop() *shop.Shop {
	return shop.NewShop(localLife.ctx)
}
