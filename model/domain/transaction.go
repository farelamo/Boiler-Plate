package domain

// TODO: All result from database should be here
type Transaction struct {
	Id     string `json:"id" bson:"_id"`
	UserId string `json:"user_id" bson:"user_id"`
	TrxId  string `json:"trx_id" bson:"trx_id"`
	Amount int    `json:"amount" bson:"amount"`
}