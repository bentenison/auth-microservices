package config

import (
	"log"
	"os"

	viper "github.com/spf13/viper"
)

type AppConfig struct {
	HTTPPort        string `mapstructure:"HTTP_PORT"`
	GRPCPort        string `mapstructure:"GRPC_PORT"`
	AppEnv          string `mapstructure:"APP_ENV"`
	IsProd          bool   `mapstructure:"IS_PROD"`
	TokenExpiryHour int    `mapstructure:"TOKEN_EXPIRY_HOUR"`
	TokenSecret     string `mapstructure:"TOKEN_SECRET"`
}

var Config *AppConfig

func InitConfig() {
	Config = &AppConfig{}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("error while getting conf file", err)
	}
	log.Println("wd", wd)
	viper.SetConfigFile("./.env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if Config.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
	log.Println(Config)
	// return Config

}
