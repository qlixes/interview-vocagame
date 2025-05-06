package provider

import (
	"log"

	"github.com/qlixes/vocagame/config"
	"github.com/spf13/viper"
)

type Config struct {
	App *config.App
	Db  *config.Db
}

type IConfig interface {
}

var Cfg = newConfig()

func newConfig() *Config {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Config not found : %s \n", err.Error())
	}

	return &Config{
		App: &config.App{
			Host: viper.GetString("APP_HOST"),
			Port: viper.GetInt("APP_PORT"),
			Name: viper.GetString("APP_NAME"),
		},
		Db: &config.Db{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetInt("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Pass: viper.GetString("DB_PASS"),
			Name: viper.GetString("DB_NAME"),
		},
	}
}
