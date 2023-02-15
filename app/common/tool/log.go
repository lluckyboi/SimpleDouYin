package tool

import (
	"SimpleDouYin/app/common"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		logger := WriteToFileInitLogger("init_error")
		defer logger.Sync()

		logger.Error(
			"Init common logger err",
			zap.Error(err),
		)
		return nil
	}
	return logger
}

func WriteToFileInitLogger(filename string) *zap.Logger {
	writeSyncer := getLogWriter(filename)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	return logger
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	var (
		file     *os.File
		err      error
		filePath = common.LogPath + "/" + filename + ".log"
	)

	_, err = os.Stat(filePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		logx.Error("init logger err:", err)
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filePath)
		if err != nil {
			logx.Error("init logger err:", err)
			return nil
		}
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			logx.Error("init logger err:", err)
			return nil
		}
	}
	return zapcore.AddSync(file)
}

func LogConstruct(info string, err string, data ...interface{}) ([]byte, error) {
	type logMsg struct {
		Time time.Time     `json:"time"`
		Info string        `json:"info"`
		Errr string        `json:"error"`
		Data []interface{} `json:"data"`
	}
	logm := logMsg{
		Time: time.Now(),
		Info: info,
		Errr: err,
	}
	for _, v := range data {
		logm.Data = append(logm.Data, v)
	}
	res, errr := json.Marshal(logm)
	log.Println("construct log", string(res))
	return res, errr
}

func SendToMsgQ(ch *amqp.Channel, queue string, info string, errr string, data ...interface{}) error {
	//构造消息
	body, err := LogConstruct(info, errr, data)
	//声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	log.Print("ready to push")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
