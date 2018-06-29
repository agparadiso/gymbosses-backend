package users

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Lastname    string        `json:"lastname" bson:"lastname"`
	Email       string        `json:"email" bson:"email"`
	AccessToken string        `json:"access_token" bson:"access_token"`
	GymID       int           `json:"gym_id" bson:"gym_id"`
}

type UserSrv interface {
	CreateUser(name, lastname, email string) (string, error)
}
