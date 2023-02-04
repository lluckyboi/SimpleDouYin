package chat

import (
	"net/http"

	"SimpleDouYin/app/service/chat/api/internal/logic/chat"
	"SimpleDouYin/app/service/chat/api/internal/svc"
	"SimpleDouYin/app/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MsgRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgRecordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewMsgRecordLogic(r.Context(), svcCtx)
		resp, err := l.MsgRecord(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
