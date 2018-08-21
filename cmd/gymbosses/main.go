package main

import (
	"log"
	"net/http"

	"github.com/agparadiso/gymbosses/backend/pkg/authentication"

	mgo "gopkg.in/mgo.v2"

	clients "github.com/agparadiso/gymbosses/backend/pkg/clients/mock"
	"github.com/agparadiso/gymbosses/backend/pkg/server"
	users "github.com/agparadiso/gymbosses/backend/pkg/users/mongoDB"
)

func main() {
	userSrv := users.NewUsersSrv(getSession())
	oauthSrv := authentication.NewOauthSrv()
	clientsSrv := clients.NewClientsSrv()
	log.Fatal(http.ListenAndServe(":3000", server.NewServer(userSrv, oauthSrv, clientsSrv)))
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		log.Fatalf("failed to get mongo running")
		panic(err)
	}
	return s
}
