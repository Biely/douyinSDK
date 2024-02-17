package shop

import (
	"encoding/json"
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
	AccountID string `json:"account_id" url:"account_id,omitempty"`
	PoiID     string `json:"poi_id" url:"poi_id,omitempty"`
	Page      int64  `json:"page" url:"page,omitempty"`
	Size      int64  `json:"size" url:"size,omitempty"`
}

type ParentAccount struct {
	AccountType string `json:"account_type"`
	AccountID   string `json:"account_id"`
	AccountName string `json:"account_name"`
}
type PoiAccount struct {
	AccountID   string `json:"account_id"`
	AccountName string `json:"account_name"`
	AccountType string `json:"account_type"`
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

type Account struct {
	ParentAccount ParentAccount `json:"parent_account"`
	PoiAccount    PoiAccount    `json:"poi_account"`
}

//	type Poi struct {
//		Latitude  float64 `json:"latitude"`
//		Longitude float64 `json:"longitude"`
//		PoiID     string  `json:"poi_id"`
//		PoiName   string  `json:"poi_name"`
//		Address   string  `json:"address"`
//	}
type Pois struct {
	Account     Account     `json:"account"`
	Poi         Poi         `json:"poi"`
	RootAccount RootAccount `json:"root_account"`
}

type ShopList struct {
	util.CommonError
	Pois  []Pois `json:"pois"`
	Total int    `json:"total"`
	// ErrorCode   int    `json:"error_code"`
	// Description string `json:"description"`
}

func NewShop(context *context.Context) *Shop {
	shop := new(Shop)
	shop.Context = context
	return shop
}

func (shop *Shop) GetShopList(param *ShopQuery) (*ShopList, error) {
	accessToken, err := shop.GetAccessToken()
	fmt.Println(accessToken)
	if err != nil {
		return nil, err
	}
	params, err := query.Values(param)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%v?%v", getShopInfoListURL, params.Encode())
	fmt.Println(url)
	header := map[string]string{
		"access-token": accessToken,
	}
	res, err := util.HTTPGet(url, header)
	if err != nil {
		return nil, err
	}
	// if res != nil {
	// 	err = fmt.Errorf(string(res))
	// 	return nil, err
	// }
	rep := response.Response{}
	rep.Data = ShopList{}
	// fmt.Println(res)
	// err = util.DecodeWithError(res, &rep, "GetShopList")
	// if err != nil {
	// 	return nil, fmt.Errorf("decodeWithError is invalid %v", err)
	// }
	err = json.Unmarshal(res, &rep)
	if err != nil {
		return nil, fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	list := rep.Data.(ShopList)
	// fmt.Println(rep)

	// err = mapstructure.Decode(rep.Data, &shopList)
	// if err != nil {
	// 	return nil, fmt.Errorf("rep data decode valid %v", err)
	// }
	return &list, err
}
