package user

type User struct {
	Name   string `json:"name" bson:"name"`
	UserId string `json:"user_id" bson:"user_id"`
	Email  string `json:"email" bson:"email"`
}
