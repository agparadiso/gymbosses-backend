package server

import (
	"fmt"
	"net/http"

	"github.com/agparadiso/gymbosses/backend/pkg/authentication"
	"github.com/agparadiso/gymbosses/backend/pkg/users"
	"github.com/rs/cors"

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
	r.HandleFunc("/", s.login)
	r.HandleFunc("/callback", s.oauthCallback)
	handler := cors.Default().Handler(r)
	return handler
}

func (s *Server) oauthCallback(w http.ResponseWriter, r *http.Request) {
	userInfo, err := s.oauthSrv.OauthCallback(w, r)
	if err != nil {
		fmt.Printf(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//pending decide either to store or not the user
	_, _ = s.userSrv.IsExistingUser(userInfo.Email)

	http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	s.oauthSrv.LoginGoogleProvider(w, r)
}
