package repository

import (
	"mading/config"
	"mading/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindTrxByUserId(userId string) ([]domain.Transaction, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"user_id": userId}
	option := options.Find().SetSort(bson.D{{Key: "created_date", Value: -1}}).SetLimit(1)

	var trx []domain.Transaction
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

// TODO: other function that need transaction table/collection
// Ex  : findTrxByIds, UpdateTrxByUserId, etc.