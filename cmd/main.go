package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServiceConfigs struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {

	viper := viper.New()
	viper.AddConfigPath("../configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Host:", viper.GetString("security.jwt.key"))

	var config ServiceConfigs
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}

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

	//encoder := getEncoderLog()
	//sync := getWriterSync()
	//core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	//logger := zap.New(core, zap.AddCaller())
	//
	//logger.Info("Info logger", zap.Int("line", 1))
	//logger.Error("Error logger", zap.Int("line", 1))

}
