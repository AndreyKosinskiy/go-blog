package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string ""
	DebugLevel int
	DbURL      string
}

func New() *Config {
	c := &Config{}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err := viper.Unmarshal(c)
	if err != nil {
		log.Fatal("Unable to decode into map, %v", err)
	}
	return &Config{Port: viper.Get("PORT").(string)}
}
