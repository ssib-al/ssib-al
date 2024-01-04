package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func NewLogger() {
	var err error

	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.MessageKey = "log"
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = ""

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	log.Info("Logger initialized")
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg error, fields ...zap.Field) {
	log.Error(msg.Error(), fields...)
}

func Fatal(msg error, fields ...zap.Field) {
	log.Fatal(msg.Error(), fields...)
}

func Panic(msg error, fields ...zap.Field) {
	log.Panic(msg.Error(), fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}
