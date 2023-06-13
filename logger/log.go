package logger

import (
	util "github.com/roverliang/sword/common/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	suger *zap.SugaredLogger
	//logger *zap.Logger
)

func Init() {
	writer := zapcore.AddSync(os.Stdout)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(util.TimeLayout)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel) // 设置日志的默认级别
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	suger = logger.Sugar()
	return
}

func Infof(template string, args ...interface{}) {
	suger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	suger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	suger.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	suger.Errorf(template, args...)
}
