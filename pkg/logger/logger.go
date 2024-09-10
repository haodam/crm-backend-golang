package logger

import (
	"github.com/haodam/user-backend-golang/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *ZapLogger {

	logLevel := config.Log_level
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		level)

	//logger := zap.New(core, zap.AddCaller())
	return &ZapLogger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1725434785.4335027 -> 2024-09-04T14:26:25.431+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": -> cmd/main.go:29 ->
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

//func getWriterSync() zapcore.WriteSyncer {
//	// Open the file with the correct flags
//	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	//if err != nil {
//	//	// Handle error (e.g., log it or panic)
//	//	panic(err)
//	//}
//	syncFile := zapcore.AddSync(file)
//	syncConsole := zapcore.AddSync(os.Stderr)
//	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
//}
