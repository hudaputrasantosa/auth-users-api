package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error
	// Konfigurasi encoder (format log)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Format waktu yang lebih mudah dibaca
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Level log dalam huruf kapital (INFO, ERROR, dll.)

	// Konfigurasi level logging
	atomicLevel := zap.NewAtomicLevelAt(zap.DebugLevel) // Atur level logging (DebugLevel untuk development)

	// Konfigurasi output (stdout dan file)
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Output dalam format JSON
			zapcore.AddSync(os.Stdout),            // Output ke console
			atomicLevel,
		),
		// Jika ingin output ke file, tambahkan core berikut:
		/*
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(getLogFile()), // Fungsi untuk membuka file log
				atomicLevel,
			),
		*/
	)
	// Build logger
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)) // Tambahkan caller dan stacktrace untuk level Error
	if err != nil {
		panic(err)
	}

	// Ganti logger global dengan logger yang baru
	zap.ReplaceGlobals(Log)
	// old
	// config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	// Log, err = config.Build()
	// if err != nil {
	// 	panic(err)
	// }
	// defer Log.Sync()
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
