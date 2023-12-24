package config

import (
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		App App
	}

	App struct {
		Env  string
		Port int
	}
)

func LoadConfigs() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("[ERROR] [LoadConfigs] load configs failed with error : ", err.Error())
	}

	cfg := &Config{App: App{
		Env:  viper.GetString("app.server.env"),
		Port: viper.GetInt("app.server.port"),
	}}

	return cfg

}
