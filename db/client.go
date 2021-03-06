package db

import (
	"log"

	pb "github.com/abdulmoeid7112/inventory-backend/pb"
)

// DataStore is an interface for query ops.
type DataStore interface {
	AddInventoryItem(item *pb.Item) (string, error)
	UpdateInventoryItem(id string, item *pb.Item) error
	GetInventoryItem(id string) (*pb.Item, error)
	ListInventoryItems(filter map[string]interface{}, lim int64, off int64) ([]*pb.Item, error)
	DeleteInventoryItem(id string) error
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)

		return
	}

	if _, ok := datastoreFactories[name]; ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)

		return
	}
}
