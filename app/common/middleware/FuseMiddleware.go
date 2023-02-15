package middleware

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/tool"
	"log"
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
)

type FuseMiddleware struct {
	Rules []*circuitbreaker.Rule
}

func (m *FuseMiddleware) NewFuseMiddleware(rule []*circuitbreaker.Rule) *FuseMiddleware {
	return &FuseMiddleware{Rules: rule}
}

func (m *FuseMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := sentinel.InitDefault()
		if err != nil {
			err := tool.SendToMsgQ(
				NewLoggerPusher(common.MsgQHost, common.MsgQUser, common.MsgQPass).
					LinkMsgQ(),
				key.LogMsgQueueName,
				"fuse middleware",
				err.Error(),
			)
			if err != nil {
				log.Println(err)
			}
			next(w, r)
		}

		_, err = circuitbreaker.LoadRules(m.Rules)

		if err != nil {
			err := tool.SendToMsgQ(
				NewLoggerPusher(common.MsgQHost, common.MsgQUser, common.MsgQPass).
					LinkMsgQ(),
				key.LogMsgQueueName,
				"fuse middleware",
				err.Error(),
			)
			if err != nil {
				log.Println(err)
			}
		}

		next(w, r)
	}
}
