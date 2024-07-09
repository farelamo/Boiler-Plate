package web

// TODO: All request and response should be here
type Transaction struct {
	UserId string `json:"user_id"`
	TrxId  string `json:"trx_id"`
	Amount int    `json:"amount"`
}
