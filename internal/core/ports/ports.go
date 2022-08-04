package ports

import (
	"net/http"

	"github.com/luispfcanales/sse-go/internal/core/domain"
)

//UserService is port to user
type UserService interface {
	Save(*domain.User) error
	UserIsValid(*domain.User) bool
	GetUser(string) *domain.User
}

//UserRepository is storage to user
type UserRepository interface {
	Create(*domain.User) bool
	GetByEmail(string) domain.User
}

//LoginService is port to Login
type LoginService interface {
	Oauth(http.ResponseWriter, *http.Request)
	OauthCallback(http.ResponseWriter, *http.Request) (*domain.User, error)
}

//TourService is port to tour
type TourService interface {
	//GetAll return all Tours of TourService
	GetAll() []domain.Tour
	//GetOneTour return Tour of TourRepository
	GetOneTour(string) (domain.Tour, error)
}

//TourRepository is storage to tour
type TourRepository interface {
	//Get return all Tours of TourRepository
	Get() []domain.Tour
	//GetByIDTour return Tour of TourRepository
	GetByIDTour(string) (domain.Tour, error)
}
