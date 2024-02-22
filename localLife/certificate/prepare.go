package certificate

import (
	"fmt"

	"github.com/Biely/douyinSDK/response"
	"github.com/Biely/douyinSDK/util"
	"github.com/google/go-querystring/query"
)

const (
	certificatePrepareUrl = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/prepare/"
)

type CertPrepareRequest struct {
	EncryptedData string `json:"encrypted_data,omitempty" url:"encrypted_data,omitempty"`
	Code          string `json:"code,omitempty" url:"code,omitempty"`
	PoiId         string `json:"poi_id" url:"poi_id"`
}

type Amount struct {
	CouponPayAmount  int `json:"coupon_pay_amount"`
	ListMarketAmount int `json:"list_market_amount"`
	OriginalAmount   int `json:"original_amount"`
	PayAmount        int `json:"pay_amount"`
}
type Sku struct {
	MarketPrice   int    `json:"market_price"`
	SkuID         string `json:"sku_id"`
	SoldStartTime int    `json:"sold_start_time"`
	ThirdSkuID    string `json:"third_sku_id"`
	Title         string `json:"title"`
	AccountID     string `json:"account_id"`
	GrouponType   int    `json:"groupon_type"`
}
type Certificates struct {
	StartTime     int    `json:"start_time"`
	Amount        Amount `json:"amount"`
	CertificateID int64  `json:"certificate_id"`
	EncryptedCode string `json:"encrypted_code"`
	ExpireTime    int    `json:"expire_time"`
	Sku           Sku    `json:"sku"`
}
type CertificatesV2 struct {
	StartTime     int    `json:"start_time"`
	Amount        Amount `json:"amount"`
	CertificateID int64  `json:"certificate_id"`
	EncryptedCode string `json:"encrypted_code"`
	ExpireTime    int    `json:"expire_time"`
	Sku           Sku    `json:"sku"`
}
type PrepareData struct {
	Certificates   []Certificates   `json:"certificates"`
	CertificatesV2 []CertificatesV2 `json:"certificates_v2"`
	OrderID        string           `json:"order_id"`
	VerifyToken    string           `json:"verify_token"`
	ErrorCode      int              `json:"error_code"`
	Description    string           `json:"description"`
}

func (certificate *Certificate) CertificatePrepare(in *CertPrepareRequest) (*PrepareData, error) {
	accessToken, err := certificate.GetAccessToken()
	// fmt.Println(accessToken)
	if err != nil {
		return nil, err
	}
	params, err := query.Values(in)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%v?%v", certificatePrepareUrl, params.Encode())
	fmt.Println(url)
	header := map[string]string{
		"access-token": accessToken,
	}
	res, err := util.HTTPGet(url, header)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	rep := response.Response{}
	rep.Data = PrepareData{}
	err = util.DecodeWithError(res, &rep, "CertificatePrepare")
	if err != nil {
		return nil, fmt.Errorf("decodeWithError is invalid %v", err)
	}
	fmt.Println(rep)
	return &PrepareData{}, err
}
