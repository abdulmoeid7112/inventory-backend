package services

import (
	"github.com/abdulmoeid7112/inventory-backend/db"
)

// Service - service struct.
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database.
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
