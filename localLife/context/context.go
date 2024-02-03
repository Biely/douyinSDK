package context

import (
	"github.com/Biely/douyinSDK/credential"
	"github.com/Biely/douyinSDK/localLife/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
