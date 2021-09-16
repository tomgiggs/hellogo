package handler

import (
	"net/http"

	"gitlab.deepwisdomai.com/infra/go-zero/rest/httpx"
	"pinggo/internal/logic"
	"pinggo/internal/svc"
)

func HealthCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewHealthCheckLogic(r.Context(), ctx)
		err := l.HealthCheck()
		if err != nil {
			httpx.FailJson(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
