package config

import (
	"mading/utils"
	"strconv"
)

var (
	// Server
	Host = utils.MustGetEnv("SERVER_URI")
	Port = utils.MustGetEnv("SERVER_PORT")
	// EndpointPrefix 		= utils.MustGetEnv("ENDPOINT_PREFIX")

	// Database
	Dsn        = utils.MustGetEnv("DB_DSN_PROD")
	DBHost     = utils.MustGetEnv("DB_HOST")
	DBPort     = utils.MustGetEnv("DB_PORT")
	DBUsername = utils.MustGetEnv("DB_USER")
	DBPassword = utils.MustGetEnv("DB_PASS")
	DBName     = utils.MustGetEnv("DB_NAME")

	// if use db pooling
	// mongoPoolMin, _     = strconv.Atoi(utils.MustGetEnv("DB_POOL_MIN"))
	// mongoPoolMax, _     = strconv.Atoi(utils.MustGetEnv("DB_POOL_MAX"))
	// mongoMaxIdleTime, _ = strconv.Atoi(utils.MustGetEnv("DB_MAX_IDLE_TIME_SECOND"))
	setTimeout, _ = strconv.Atoi(utils.MustGetEnv("DB_TIMEOUT"))

	// Variables
	LogLevel = utils.MustGetEnv("LOG_LEVEL")
	Prod     = utils.MustGetEnv("PROD")

	// Another env variables here
)
