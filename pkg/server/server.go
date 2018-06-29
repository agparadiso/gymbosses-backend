package server

import (
	"fmt"
	"net/http"

	"github.com/agparadiso/gymbosses/backend/pkg/users"

	"github.com/gorilla/mux"
)

type Server struct {
	userSrv users.UserSrv
}

func NewServer(user users.UserSrv) http.Handler {
	fmt.Println("Running gymbosses server...")
	s := &Server{userSrv: user}
	r := mux.NewRouter()
	r.HandleFunc("/user", s.createUser).Methods("POST")

	return r
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	response, err := s.userSrv.CreateUser(name, lastname, email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
