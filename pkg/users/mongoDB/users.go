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

func (srv *UserSrv) IsExistingUser(email string) (bool, error) {
	c := srv.session.DB(dbname).C(usersTable)
	query := bson.M{"email": email}
	user := &users.User{}
	err := c.Find(query).One(&user)
	if err != nil {
		return false, err
	}

	if user != nil && user.Email == email {
		return true, nil
	}

	return false, nil
}

func (srv *UserSrv) SignUp(name, email string, gymName string) (string, error) {
	user := &users.User{}
	gym := &users.Gym{}

	gym.ID = bson.NewObjectId()
	gym.Name = gymName

	// Add an Id
	user.ID = bson.NewObjectId()
	user.Name = name
	user.Email = email
	user.Gym = gym

	// Write the user to mongo
	err := srv.session.DB(dbname).C(usersTable).Insert(user)
	if err != nil {
		return "", errors.Wrap(err, "Failed to store user in mongoDB")
	}

	return fmt.Sprintf("user %s with email %s created in mongoDB", name, email), nil
}
