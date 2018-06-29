package server

import (
	"fmt"
	"net/http"

	"github.com/agparadiso/gymbosses/backend/pkg/authentication"
	"github.com/agparadiso/gymbosses/backend/pkg/users"

	"github.com/gorilla/mux"
)

type Server struct {
	userSrv  users.UserSrv
	oauthSrv authentication.OauthSrv
}

func NewServer(userSrv users.UserSrv, oauthSrv *authentication.OauthSrv) http.Handler {
	fmt.Println("Running gymbosses server...")
	s := &Server{userSrv: userSrv, oauthSrv: *oauthSrv}
	r := mux.NewRouter()
	r.HandleFunc("/", s.login).Methods("GET")
	r.HandleFunc("/callback", s.oauthCallback)

	return r
}

func (s *Server) oauthCallback(w http.ResponseWriter, r *http.Request) {
	userInfo, err := s.oauthSrv.OauthCallback(w, r)
	if err != nil {
		fmt.Printf(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	exists, err := s.userSrv.IsExistingUser(userInfo.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !exists {
		resp, err := s.userSrv.SignUp(userInfo.Name, userInfo.Email, "sanBorja")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ya existia"))

}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	s.oauthSrv.LoginGoogleProvider(w, r)
}
