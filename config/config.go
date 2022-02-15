package config

import (
	"log"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigFile(".env")

	viper.SetDefault("JWT_SECRET", "SOMETHINGSECRET")
	
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func Get(key string) string {
	value, ok := config.Get(key).(string)
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
  	return value
}