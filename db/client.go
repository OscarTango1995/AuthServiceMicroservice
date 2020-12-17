package db

import (
	"log"

	"github.com/OscarTango1995/AuthServiceMicroservice/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	CreateUser(user *models.User) (string, error)
	ListUser(id string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(user *models.User) error

	CreateSession(session *models.Session) (string, error)
	ListSession(id string) (*models.Session, error)
	DeleteSession(id string) error
	UpdateSession(session *models.Session) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
