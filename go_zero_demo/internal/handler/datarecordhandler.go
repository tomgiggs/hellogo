package handler

import (
	"net/http"

	"gitlab.deepwisdomai.com/infra/go-zero/rest/httpx"
	"pinggo/internal/logic"
	"pinggo/internal/svc"
)

func DataRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewDataRecordLogic(r.Context(), ctx)
		resp, err := l.DataRecord()
		if err != nil {
			httpx.FailJson(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
