package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServiceConfigs struct {
}

func MustLoadConfig(configPath string) *ServiceConfigs {

	viper := viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Host:", viper.GetString("server.host"))

	var config ServiceConfigs
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
	return &ServiceConfigs{}

}
