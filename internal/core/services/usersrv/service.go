package usersrv

import (
	"errors"
	"log"
	"net/http"

	"github.com/luispfcanales/sse-go/internal/core/domain"
	"github.com/luispfcanales/sse-go/internal/core/ports"
)

//service is template to user Service
type service struct {
	storage ports.UserRepository
}

//revive:disable:unexported-return
//New return new user service
func New(repo ports.UserRepository) *service {
	return &service{
		storage: repo,
	}
}

//revive:enable:unexported-return
//ExistSession Valid Session exits
func (s *service) ExistSession(r *http.Request) (bool,string) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false,""
		}
		log.Println("Error to get cookie")
		return false,""
	}
	return true,c.Value
}

//Save return error
func (s *service) Save(user *domain.User) error {
	ok := s.storage.Create(user)
	if !ok {
		return errors.New("User already exists")
	}
	return nil
}

//UserIsValid verify user
func (s *service) GmailIsValid(email string ) bool {
	validUser := s.storage.GetByEmail(email)
	if validUser.ID == "" {
		return false
	}
	return true
}

//GetUser return user exits of database
func (s *service) GetUserWithCredentials(email ,password string) *domain.User {
	user := s.storage.GetByEmail(email)
	if user.Password == password {
		return &user
	}
	return &domain.User{}
}
