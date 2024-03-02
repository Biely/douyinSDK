package certificate

import (
	"fmt"

	"github.com/Biely/douyinSDK/response"
	"github.com/Biely/douyinSDK/util"
)

const (
	certificateVerifyUrl = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/verify/"
)

type CertVerifyRequest struct {
	VerifyToken    string   `json:"verify_token" form:"verify_token"`
	EncryptedCodes []string `json:"encrypted_codes" form:"encrypted_codes[]"`
	PoiID          string   `json:"poi_id" form:"poi_id"`
}

func (certificate *Certificate) CertificateVerify(in *CertVerifyRequest) (rep interface{}, err error) {
	accessToken, err := certificate.GetAccessToken()
	// fmt.Println(accessToken)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"access-token": accessToken,
	}
	res, err := util.PostJSON(certificateVerifyUrl, in, header)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	rep = response.Response{}
	err = util.DecodeWithError(res, &rep, "CertVerify")
	if err != nil {
		return nil, fmt.Errorf("decodeWithError is invalid %v", err)
	}
	return
}
