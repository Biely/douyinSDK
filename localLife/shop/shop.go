package shop

import (
	"fmt"

	"github.com/Biely/douyinSDK/localLife/context"
	"github.com/Biely/douyinSDK/response"
	"github.com/Biely/douyinSDK/util"
	"github.com/google/go-querystring/query"
)

const (
	getShopInfoListURL = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
)

type Shop struct {
	*context.Context
}

// 请求参数
type ShopQuery struct {
	AccountID string `json:"account_id"`
	PoiID     string `json:"poi_id"`
	Page      int64  `json:"page"`
	Size      int64  `json:"size"`
}

type Poi struct {
	PoiID     string  `json:"poi_id"`
	PoiName   string  `json:"poi_name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RootAccount struct {
	AccountID   string `json:"account_id"`
	AccountName string `json:"account_name"`
}

type Pois struct {
	Poi         Poi         `json:"poi"`
	RootAccount RootAccount `json:"root_account"`
}

type ShopList struct {
	util.CommonError
	Pois  []Pois `json:"pois"`
	Total int64  `json:"total"`
}

func NewShop(context *context.Context) *Shop {
	shop := new(Shop)
	shop.Context = context
	return shop
}

func (shop *Shop) GetShopList(param *ShopQuery) (*ShopList, error) {
	accessToken, err := shop.GetAccessToken()
	if err != nil {
		return nil, err
	}
	params, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%v?%v", getShopInfoListURL, params.Encode())
	header := map[string]string{
		"access-token": accessToken,
	}
	res, err := util.HTTPGet(url, header)
	if err != nil {
		return nil, err
	}
	shopList := ShopList{}
	rep := response.Response{}
	rep.Data = &shopList
	fmt.Println(res)
	err = util.DecodeWithError(res, &rep, "GetShopList")
	if err != nil {
		return nil, err
	}
	return &shopList, err
}
