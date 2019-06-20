package postgres

// import (
// 	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
// 	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
// 	"sync"
// )

// type inMemoryRepository struct {
// 	m   map[entity.Customer]error
// 	lck sync.RWMutex
// }

// var c = entity.Customer{1, "test name", 1234.56, 1, "test warning"}

// func newInMemoryRepository() adding.Repository {
// 	return &inMemoryRepository{m: make(map[entity.Customer]error)}
// }

// func RepoAddingLogic(r adding.Repository) {
// 	r.AddCustomer(c)
// }

// func (i *inMemoryRepository) AddCustomer(c entity.Customer) error {
// 	i.lck.RLock()
// 	defer i.lck.RUnlock()
// 	return i.m[c]
// }

// func TestAddCustomer(t *testing.T) {
// 	repo := newInMemoryRepository()
// 	RepoAddingLogic(repo)
// 	err := repo.AddCustomer(c)
// 	assert.NoError(t, err, err)
// }
