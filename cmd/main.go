package main

import (
	"go.uber.org/zap"
)

func main() {

	//r := routers.NewRouter()
	//err := r.Run()
	//if err != nil {
	//	return
	//} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name:%s, age:%d", "TipGo", 40) // like fmt.Printf(format,a)

	// logger
	logger := zap.NewExample()
	logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 40))

}
