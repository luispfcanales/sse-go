package memtour

import (
	"errors"
	"sync"

	"github.com/luispfcanales/sse-go/internal/core/domain"
	"github.com/luispfcanales/sse-go/pkg/gendata"
)

//repository is template to memory Tour
type repository struct {
	mem  map[string]*domain.Tour
	lock *sync.RWMutex
}

//revive:disable:unexported-return
//New return new repository to Tours
func New() *repository {

	data := make(map[string]*domain.Tour)
	gendata.GenerateTours(data)
	return &repository{
		mem:  data,
		lock: &sync.RWMutex{},
	}
}

//revive:enable:unexported-return

//Get return slice of tours
func (r *repository) Get() []domain.Tour {
	r.lock.RLock()
	defer r.lock.RUnlock()
	tours := []domain.Tour{}
	for _, tour := range r.mem {
		tours = append(tours, *tour)
	}
	return tours
}

//GetByIDTour return tour
func (r *repository) GetByIDTour(id string) (domain.Tour, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	tour, ok := r.mem[id]
	if !ok {
		return domain.Tour{}, errors.New("user not found")
	}
	return *tour, nil
}
