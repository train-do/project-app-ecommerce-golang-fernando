package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog(config Configuration) *zap.Logger {
	file, _ := os.Create("app.log")
	fileWriter := zapcore.AddSync(file)
	consoleWriter := zapcore.AddSync(os.Stdout)
	if config.Debug {
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			consoleWriter,
			zap.DebugLevel,
		)
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
			fileWriter,
			zap.DebugLevel,
		)
		core := zapcore.NewTee(consoleCore, fileCore)
		logger := zap.New(core)
		return logger
	} else {
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			fileWriter,
			zap.InfoLevel,
		)
		logger := zap.New(core)
		return logger
	}
}
