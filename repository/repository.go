package repository

import (
	"mading/helper/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	trxCol *mongo.Collection
	log    *logger.Log
}

func New(database *mongo.Database) *Repository {
	return &Repository{
		trxCol: database.Collection("transaction"),
		log:    logger.New("Repository"),
	}
}
