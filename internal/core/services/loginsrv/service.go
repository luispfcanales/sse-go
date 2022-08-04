package loginsrv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/luispfcanales/sse-go/internal/core/domain"
	"github.com/luispfcanales/sse-go/internal/core/ports"
	"github.com/luispfcanales/sse-go/pkg/infoapi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//service is a template to service Login
type service struct {
	storage ports.UserRepository
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	Scopes: []string{
		infoapi.GetAPIGoogleEmail(),
		infoapi.GetAPIGoogleProfile(),
	},
	Endpoint: google.Endpoint,
}

//revive:disable:unexported-return
//New return serviceLogin
func New(repo ports.UserRepository) *service {
	return &service{
		storage: repo,
	}
}

//revive:enable:unexported-return

//Oauth validate is user
func (s *service) Oauth(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

//OauthCallback get data is authenticated
func (s *service) OauthCallback(w http.ResponseWriter, r *http.Request) (*domain.User, error) {
	user := &domain.User{}
	if r.FormValue("state") != "randomstate" {
		return user, errors.New("invalid oauth google state")
	}
	data, expiryTokenOaut, err := s.getDataFormGoogle(r.FormValue("code"))
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, user)
	if err != nil {
		return nil, errors.New("error to parse json")
	}
	user.Expiry = expiryTokenOaut
	return user, nil
}
func (s *service) getDataFormGoogle(code string) ([]byte, time.Time, error) {
	currentTime := time.Now()
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, currentTime, fmt.Errorf("code exchange wrong %s", err.Error())
	}
	response, err := http.Get(infoapi.GetAPIGoogleUserInfo(token.AccessToken))
	if err != nil {
		return nil, currentTime, fmt.Errorf("failed getting user info %s", err.Error())
	}
	defer response.Body.Close()
	userData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, currentTime, fmt.Errorf("failed read response %s", err.Error())
	}
	return userData, token.Expiry, nil
}
