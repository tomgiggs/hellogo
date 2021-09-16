package handler

import (
	"net/http"

	"pinggo/internal/logic"
	"pinggo/internal/svc"
	"pinggo/internal/types"

	"gitlab.deepwisdomai.com/infra/go-zero/rest/httpx"
)

func DataBatchUploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TrainDataBatchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.FailJson(w, err)
			return
		}

		l := logic.NewDataBatchUploadLogic(r.Context(), ctx)
		resp, err := l.DataBatchUpload(req)
		if err != nil {
			httpx.FailJson(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
