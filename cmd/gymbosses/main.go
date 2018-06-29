package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/agparadiso/gymbosses/backend/pkg/server"
	"github.com/agparadiso/gymbosses/backend/pkg/users/mongoDB"
)

func main() {
	userSrv := mongoDB.NewUsersSrv(getSession())
	log.Fatal(http.ListenAndServe(":8080", server.NewServer(userSrv)))
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
