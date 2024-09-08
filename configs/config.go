package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MySQLConfigs string `mapstructure:"mysql_configs"`
}

type MySQLConfigs struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
}

func MustLoadConfig(configPath string) *Config {

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

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
	return &Config{}

}
