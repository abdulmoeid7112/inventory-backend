package inventory

import (
	"github.com/abdulmoeid7112/inventory-backend/server/services"
	"github.com/pkg/errors"

	"github.com/abdulmoeid7112/inventory-backend/db"
	"github.com/abdulmoeid7112/inventory-backend/db/mongo"
)

// Runtime initializes values for entry point to our application.
type Runtime struct {
	svc *services.Service
}

// NewRuntime creates a new database service.
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create mongo client")
	}

	return &Runtime{svc: services.NewService(store)}, nil
}

// Service returns pointer to our service variable.
func (r Runtime) Service() *services.Service {
	return r.svc
}
