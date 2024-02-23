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
	CouponPayAmount  int32 `json:"coupon_pay_amount"`
	ListMarketAmount int32 `json:"list_market_amount"`
	OriginalAmount   int32 `json:"original_amount"`
	PayAmount        int32 `json:"pay_amount"`
}

type TimeCardAmount struct {
	OriginalAmount         int32 `json:"original_amount"`
	PayAmount              int32 `json:"pay_amount"`
	MerchantTicketAmount   int32 `json:"merchant_ticket_amount"`
	ListMarketAmount       int32 `json:"list_market_amount"`
	PlatformDiscountAmount int32 `json:"platform_discount_amount"`
	PaymentDiscountAmount  int32 `json:"payment_discount_amount"`
	CouponPayAmount        int32 `json:"coupon_pay_amount"`
}

type SerialAmountList struct {
	SerialNumb int32          `json:"serial_numb"`
	Amount     TimeCardAmount `json:"amount"`
}
type Sku struct {
	MarketPrice   int32  `json:"market_price"`
	SkuID         string `json:"sku_id"`
	SoldStartTime int32  `json:"sold_start_time"`
	ThirdSkuID    string `json:"third_sku_id"`
	Title         string `json:"title"`
	AccountID     string `json:"account_id"`
	GrouponType   int32  `json:"groupon_type"`
}
type Certificates struct {
	StartTime     int64  `json:"start_time"`
	Amount        Amount `json:"amount"`
	CertificateID int64  `json:"certificate_id"`
	EncryptedCode string `json:"encrypted_code"`
	ExpireTime    int64  `json:"expire_time"`
	Sku           Sku    `json:"sku"`
}
type CertificatesV2 struct {
	StartTime     int64  `json:"start_time"`
	Amount        Amount `json:"amount"`
	CertificateID int64  `json:"certificate_id"`
	EncryptedCode string `json:"encrypted_code"`
	ExpireTime    int64  `json:"expire_time"`
	Sku           Sku    `json:"sku"`
}

type TimeCard struct {
	TimesCount       int32              `json:"times_count"`
	TimesUsed        int32              `json:"times_used"`
	SerialAmountList []SerialAmountList `json:"serial_amount_list"`
}
type PrepareData struct {
	Certificates   []Certificates   `json:"certificates"`
	CertificatesV2 []CertificatesV2 `json:"certificates_v2"`
	TimeCard       TimeCard         `json:"time_card"`
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
	// fmt.Println(string(res))
	rep := response.Response{}
	rep.Data = PrepareData{}
	// nrep := rep
	err = util.DecodeWithError(res, &rep, "CertificatePrepare")
	if err != nil {
		return nil, fmt.Errorf("decodeWithError is invalid %v", err)
	}

	fmt.Println(rep)
	return &PrepareData{}, err
}
