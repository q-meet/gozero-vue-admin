package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/logic/operationRecord"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func DeleteSysOperationRecordByIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req utils.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteSysOperationRecordByIdsLogic(r.Context(), ctx)
		err := l.DeleteSysOperationRecordByIds(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
