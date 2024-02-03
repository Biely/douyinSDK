package shop

import "github.com/Biely/douyinSDK/localLife/context"

const (
	getShopInfoListURL = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
)

type Shop struct {
	*context.Context
}

func NewShop(context *context.Context) *Shop {
	shop := new(Shop)
	shop.Context = context
	return shop
}

func (shop *Shop) GetShopList()
