package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	DebugLevel int
	DbURL      string
}

func New() *Config {
	log.Println("Init config ....")
	c := &Config{}
	if err := loadEnvFile(c); err != nil {
		log.Fatal(err)
	}
	port := viper.GetString("PORT")
	debugLevel := viper.GetInt("LOGGER_LEVEL")
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"), viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"), viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"), viper.GetString("DB_SSL_MODE"))
	log.Println("Init config success!")
	return &Config{port, debugLevel, dbUrl}
}

func loadEnvFile(c *Config) (err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	return
}
