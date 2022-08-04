package toursrv

import (
	"github.com/luispfcanales/sse-go/internal/core/domain"
	"github.com/luispfcanales/sse-go/internal/core/ports"
)

//service is template to Tour Service
type service struct {
	storage ports.TourRepository
}

//revive:disable:unexported-return
//New return new Tour service
func New(repo ports.TourRepository) *service {
	return &service{
		storage: repo,
	}
}

//revive:enable:unexported-return

//GetAll return slice of tours
func (s *service) GetAll() []domain.Tour {
	tours := s.storage.Get()
	return tours
}
func (s *service) GetOneTour(id string) (domain.Tour, error) {
	tour, err := s.storage.GetByIDTour(id)
	if err != nil {
		return domain.Tour{}, err
	}
	return tour, nil
}
