package config

import (
	"context"
	"fmt"
	"mading/helper/logger"
	"mading/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(setTimeout) * time.Second)
}

func NewDB() *mongo.Database {
	log := logger.New("Config_New_DB")

	ctx, cancel := NewDBContext()
	defer cancel()

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", DBUsername, DBPassword, DBHost, DBPort, DBName)
	if Prod == "true" {
		dsn = utils.MustGetEnv("DB_DSN_PROD")
	}

	log.Log(dsn)

	option := options.Client().
		ApplyURI(dsn)
		// SetMinPoolSize(uint64(mongoPoolMin)).
		// SetMaxPoolSize(uint64(mongoPoolMax)).
		// SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Panic("Error when connect to database", "error", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panic("Error when ping to database", "error", err)
	}

	log.Log("Database connected",
		"database", DBName,
		"host", DBHost,
		"port", DBPort,
	)

	return client.Database(DBName)
}
