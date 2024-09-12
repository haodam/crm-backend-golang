package configs

import (
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/spf13/viper"
)

func MustLoadConfig() {

	viper := viper.New()
	viper.AddConfigPath("./deploy/conf/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Host:", viper.GetString("server.host"))

	//var config Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}

}
