package certificate

import "github.com/Biely/douyinSDK/localLife/context"

type Certificate struct {
	*context.Context
}

func NewCertificate(ctx *context.Context) *Certificate {
	return &Certificate{ctx}
}
