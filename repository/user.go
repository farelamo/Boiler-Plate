package repository

import (
	"mading/config"
	"mading/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) TrxFindUser(username, email string) (domain.Users, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	pattern := bson.M{"$regex": "^" + username}
	query := bson.M{"email": email, "username": pattern}
	option := options.Find().SetSort(bson.D{{Key: "name", Value: -1}}).SetLimit(1)

	var trx domain.Users
	result, err := r.trxCol.Find(ctx, query, option)
	if err != nil {
		return trx, result.Err()
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		err := result.Decode(&trx)
		if err != nil {
			return trx, err
		}
	}

	if err := result.Err(); err != nil {
		return trx, err
	}

	return trx, nil
}

// TODO: other function that need user table/collection
// Ex  : findUserByUsername, findUserByEmail, etc.