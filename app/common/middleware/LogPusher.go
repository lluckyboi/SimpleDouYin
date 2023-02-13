package middleware

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/tool"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var logger = tool.InitLogger()

type LoggerPusher struct {
	MsgQHost string
	MsgQUser string
	MsgQPass string
}
type Ch *amqp.Channel

func NewLoggerPusher(host, user, pass string) *LoggerPusher {
	return &LoggerPusher{
		MsgQHost: host,
		MsgQUser: user,
		MsgQPass: pass,
	}
}

func (l *LoggerPusher) linkMsgQ() Ch {
	//连接RMQ
	RMQConn, err := amqp.Dial("amqp://" + l.MsgQUser + ":" + l.MsgQPass + "@" + l.MsgQHost + "/")
	if err != nil {
		logger.Error("link RMQ err", zap.Error(err))
		return nil
	}
	//创建一个通道，大多数API都是用过该通道操作的。
	ch, err := RMQConn.Channel()
	if err != nil {
		logger.Error("create RMQ channel err", zap.Error(err))
		return nil
	}
	return ch
}

func (l *LoggerPusher) WithConsole(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer logger.Sync()

		if ch := l.linkMsgQ(); ch != nil {
			next(w, r)
		} else {
			start := time.Now()
			path := r.URL.Path
			query := r.URL.RawQuery

			msgs := make([]byte, key.LogReadBuffer)
			_, err := r.Response.Body.Read(msgs)
			if err != nil {
				logger.Error("read request msg err", zap.Error(err))
				return
			}

			next(w, r)

			cost := time.Since(start)
			logger.Info(path,
				zap.Int("status", r.Response.StatusCode),
				zap.String("method", r.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("msgs", string(msgs)),
				zap.String("user-agent", r.UserAgent()),
				zap.Duration("cost", cost),
			)
		}
	}
}
