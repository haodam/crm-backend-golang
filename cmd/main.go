package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	//r := routers.NewRouter()
	//err := r.Run()
	//if err != nil {
	//	return
	//} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//sugar := zap.NewExample().Sugar()
	//sugar.Infof("Hello name:%s, age:%d", "TipGo", 40) // like fmt.Printf(format,a)
	//
	//// logger
	//logger := zap.NewExample()
	//logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 40))
	//
	//// Development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("Hello NewDevelopmentLogger")
	//
	//// Product
	//logger, _ = zap.NewProduction()
	//logger.Info("Hello NewProductionLogger")

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info logger", zap.Int("line", 1))
	logger.Error("Error logger", zap.Int("line", 1))

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

func getWriterSync() zapcore.WriteSyncer {
	// Open the file with the correct flags
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	// Handle error (e.g., log it or panic)
	//	panic(err)
	//}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
