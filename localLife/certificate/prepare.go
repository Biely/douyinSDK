package certificate

import (
	"encoding/json"
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

func (certificate *Certificate) CertificatePrepare(in *CertPrepareRequest) (res interface{}, err error) {
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
	data, err := util.HTTPGet(url, header)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	rep := response.Response{}
	err = json.Unmarshal(data, &rep)
	if err != nil {
		return nil, fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	return rep, err
}
