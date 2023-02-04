package action

import (
	"net/http"

	"SimpleDouYin/app/service/action/api/internal/logic/action"
	"SimpleDouYin/app/service/action/api/internal/svc"
	"SimpleDouYin/app/service/action/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := action.NewFollowLogic(r.Context(), svcCtx)
		resp, err := l.Follow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
