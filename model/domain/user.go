package domain

// TODO: All result from database should be here
type Users struct {
	Id 				string 		`json:"id" bson:"_id"`
	Email 			string 		`json:"emaik" bson:"email"`
	Username		string 		`json:"username" bson:"username"`
	FirstName		string 		`json:"firstname" bson:"firstname"`
	LastName		string 		`json:"lastname" bson:"lastname"`
}