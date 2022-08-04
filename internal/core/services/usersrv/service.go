package usersrv

import (
	"errors"
	"log"

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

//Save return error
func (s *service) Save(user *domain.User) error {
	ok := s.storage.Create(user)
	if !ok {
		return errors.New("User already exists")
	}
	return nil
}
func (s *service) UserIsValid(user *domain.User) bool {
	validUser := s.storage.GetByEmail(user.Email)
	if validUser.ID == "" {
		return false
	}
	return true
}

//GetUser return user exits of database
func (s *service) GetUser(idToken string) *domain.User {
	log.Println(idToken, ": ", "getting user...")
	return &domain.User{}
}
