package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServiceConfigs struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
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
