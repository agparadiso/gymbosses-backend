package mongoDB

import (
	"fmt"

	"github.com/agparadiso/gymbosses/backend/pkg/users"
	"github.com/pkg/errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbname     = "gymbosses"
	usersTable = "users"
)

type UserSrv struct {
	session *mgo.Session
}

func NewUsersSrv(s *mgo.Session) *UserSrv {
	return &UserSrv{
		session: s,
	}
}

func (srv *UserSrv) CreateUser(name, lastname, email string) (string, error) {
	user := &users.User{}

	// Add an Id
	user.ID = bson.NewObjectId()
	user.Name = name
	user.Lastname = lastname
	user.Email = email

	// Write the food to mongo
	err := srv.session.DB(dbname).C(usersTable).Insert(user)
	if err != nil {
		return "", errors.Wrap(err, "Failed to store user in mongoDB")
	}

	return fmt.Sprintf("user %s %s with email %s created in mongoDB", name, lastname, email), nil
}
