package userhdl

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luispfcanales/sse-go/internal/core/ports"
)

//userHandler is Template
type userHandler struct {
	serviceUser  ports.UserService
	serviceLogin ports.LoginService
}

//revive:disable:unexported-return
//New return handler http to user
func New(usersrv ports.UserService, loginsrv ports.LoginService) *userHandler {
	return &userHandler{
		serviceUser:  usersrv,
		serviceLogin: loginsrv,
	}
}

//revive:enable:unexported-return

//SetupRoutes initial routes to handle http requets of users
func (hdl *userHandler) SetupRoutes(m *mux.Router) {
	m.HandleFunc("/auth/login", hdl.login).Methods("GET")
	m.HandleFunc("/auth/google/callback", hdl.callback).Methods("GET")
}

func (hdl *userHandler) login(w http.ResponseWriter, r *http.Request) {
	hdl.serviceLogin.Oauth(w, r)
}
func (hdl *userHandler) callback(w http.ResponseWriter, r *http.Request) {
	user, err := hdl.serviceLogin.OauthCallback(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	//parte de registro aqui
	isValid := hdl.serviceUser.GmailIsValid(user.Email)
	if !isValid {
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		return
	}
	//http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "luispf", Expires: time.Now().Add(20 * time.Second), Path: "/"})
	//se elimina la cookie en expires
	http.Redirect(w, r, "/paquetes", http.StatusSeeOther)
}
