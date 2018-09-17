package main

import (
	"log"
	"net/http"
	"os"

	"github.com/agparadiso/gymbosses/backend/pkg/authentication"

	clients "github.com/agparadiso/gymbosses/backend/pkg/clients/mock"
	"github.com/agparadiso/gymbosses/backend/pkg/server"
	users "github.com/agparadiso/gymbosses/backend/pkg/users/mock"
)

func main() {
	userSrv := users.NewUsersSrv()
	oauthSrv := authentication.NewOauthSrv()
	clientsSrv := clients.NewClientsSrv()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), server.NewServer(userSrv, oauthSrv, clientsSrv)))
}
