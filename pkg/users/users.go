package users

import "gopkg.in/mgo.v2/bson"

type Gym struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

type User struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Email string        `json:"email" bson:"email"`
	Gym   *Gym          `json:"gym" bson:"gym"`
}

type UserSrv interface {
	IsExistingUser(email string) (bool, error)
	SignUp(name, email, gymName string) (string, error)
}
