package memory

import (
	"sync"

	"github.com/luispfcanales/sse-go/internal/core/domain"
)

type repository struct {
	mem  map[string]*domain.User
	lock *sync.RWMutex
}

//revive:disable:unexported-return
//New return repository in memory
func New() *repository {
	data := make(map[string]*domain.User)
	data["luispfcanales@gmail.com"] = &domain.User{
		ID: "12312312",
		Email: "luispfcanales@gmail.com",
		Password: "luis1234",
		Name: "Luis Angel",
		FamilyName: "Pfuno canales",
	}
	return &repository{
		mem:  data,
		lock: &sync.RWMutex{},
	}
}

//revive:enable:unexported-return

//Create save to user in repository
func (r *repository) Create(user *domain.User) bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	_, ok := r.mem[user.Email]
	if !ok {
		return ok
	}
	r.mem[user.Email] = user
	return ok
}

//GetByEmail method
func (r *repository) GetByEmail(email string) domain.User {
	r.lock.RLock()
	defer r.lock.RUnlock()
	user, ok := r.mem[email]
	if !ok {
		return domain.User{}
	}
	return *user
}
