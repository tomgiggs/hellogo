package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"pinggo/internal/svc"
)

func MetricsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)

		//l := logic.NewMetricsLogic(r.Context(), ctx)
		//err := l.Metrics()
		//if err != nil {
		//	httpx.FailJson(w, err)
		//} else {
		//	httpx.Ok(w)
		//}
	}
}
