package middleware

import (
	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
	"net/http"
	"pinggo/internal/constants"
	"pinggo/internal/svc"
)

func CheckToken(serverCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) bool {
	if _, exist := constants.UrlWhitelist[r.URL.Path]; exist {
		return true
	}
	if token := r.Header.Get("authorization"); token == "" {
		logx.Errorf("empty authorization")
		return false
	}

	return true
}
