package shop

import (
	"github.com/Biely/douyinSDK/localLife/context"
	"github.com/Biely/douyinSDK/util"
)

const (
	getShopInfoListURL = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
)

type Shop struct {
	*context.Context
}

type Poi struct {
	PoiId string `json:"poi_id"`
}

type Pois struct {
}
type ShopList struct {
	util.CommonError
	Pois
}

func NewShop(context *context.Context) *Shop {
	shop := new(Shop)
	shop.Context = context
	return shop
}

// func (shop *Shop) GetShopList() ()
