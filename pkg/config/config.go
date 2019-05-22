package config

import (
	"os"
	"pkg/log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ENVIRONMENT  string
	APP_PORT     string
	TLS_CERT     string
	TLS_CERT_KEY string
}

var Conf Config

func LoadConfig() {
	if env, _ := os.LookupEnv("ENVIRONMENT"); env == "DEV" {
		err := godotenv.Load()
		if err != nil {
			log.Error(err, ERR_CONFIG_NO_ENV_FILE, "")
		}
	}

	err := envconfig.Process("", &Conf)
	if err != nil {
		log.Error(err, ERR_CONFIG_LOAD_FAILURE, Conf.ENVIRONMENT)
	}
}
