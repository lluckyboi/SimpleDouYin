package middleware

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/tool"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var logger = tool.InitLogger()

type LoggerPusher struct {
	MsgQHost string
	MsgQUser string
	MsgQPass string
}

func NewLoggerPusher(host, user, pass string) *LoggerPusher {
	return &LoggerPusher{
		MsgQHost: host,
		MsgQUser: user,
		MsgQPass: pass,
	}
}

func (l *LoggerPusher) linkMsgQ() *amqp.Channel {
	//连接RMQ
	RMQConn, err := amqp.Dial("amqp://" + l.MsgQUser + ":" + l.MsgQPass + "@" + l.MsgQHost + "/")
	log.Println("amqp://" + l.MsgQUser + ":" + l.MsgQPass + "@" + l.MsgQHost + "/")
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
		start := time.Now()
		path := r.URL.Path
		query := r.URL.RawQuery

		next(w, r)

		cost := time.Since(start)
		logger.Info(path,
			zap.String("method", r.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("user-agent", r.UserAgent()),
			zap.Duration("cost", cost),
		)
	}
}

func (l *LoggerPusher) WithMsgQ(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ch := l.linkMsgQ(); ch == nil {
			next(w, r)
			return
		} else {
			start := time.Now()
			path := r.URL.Path
			query := r.URL.RawQuery
			next(w, r)
			cost := time.Since(start)
			log.Println("init data ok")
			err := tool.SendToMsgQ(ch, key.LogMsgQueueName, "middleware log", "",
				r.Method,
				path,
				query,
				r.UserAgent(),
				cost,
			)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
	}
}
