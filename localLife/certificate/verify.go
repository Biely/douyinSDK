package certificate

import "github.com/Biely/douyinSDK/util"

const (
	certificateVerifyUrl = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/verify/"
)

type CertVerifyRequest struct {
}

func (certificate *Certificate) CertificateVerify(in *CertVerifyRequest) (res interface{}, err error) {
	accessToken, err := certificate.GetAccessToken()
	// fmt.Println(accessToken)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"access-token": accessToken,
	}
	res, err = util.PostJSON(certificateVerifyUrl, in, header)
	if err != nil {
		return nil, err
	}
	return
}
