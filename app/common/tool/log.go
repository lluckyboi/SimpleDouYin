package tool

import (
	"SimpleDouYin/app/common"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
