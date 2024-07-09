package utils

import (
	"fmt"
	"mading/helper/logger"
	"os"

	"github.com/joho/godotenv"
)

var (
	log      = logger.New("ENV")
	logLevel = os.Getenv("LOG_LEVEL")
)

func MustGetEnv(key string) string {
	if valueDotEnv, err := getDotEnv(key); err == nil && len(valueDotEnv) > 0 {
		if logLevel == "DEBUG" {
			log.Log(fmt.Sprintf("found environment variable %s: %s", key, valueDotEnv))
		}
		return valueDotEnv
	}

	log.Panic(fmt.Sprintf("missing environment variable %s", key))
	return ""
}

func getDotEnv(key string) (string, error) {
	if err := godotenv.Load(); err != nil {
		err := fmt.Errorf("missing dotenv variable %s from .env", key)
		return "", err
	}

	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("no value for dotenv variable %s from .env", key)
	}

	return value, nil
}
