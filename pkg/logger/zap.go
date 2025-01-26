package logger

import (
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)
var Log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
    config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

    Log, err = config.Build()
    if err != nil {
        panic(err)
    }
    defer Log.Sync()
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Log.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}
