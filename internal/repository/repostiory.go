package repository

import (
	"receipt-processor/internal/errors"

	"github.com/google/uuid"
)

// KVRepo is a simple in-memory key-value store for receipt points.
type KVRepo struct {
	store map[string]int64
}

// NewKVRepo creates a new instance of KVRepo.
func NewKVRepo() *KVRepo {
	return &KVRepo{
		store: make(map[string]int64),
	}
}

// GetById retrieves the points for a given receipt ID.
func (repo *KVRepo) GetById(id string) (int64, error) {
	points, exists := repo.store[id]
	if !exists {
		return -1, errors.NewKeyNotFoundError(id)
	}
	return points, nil
}

// AddEntry adds a new entry to the repository.
func (repo *KVRepo) AddEntry(points int64) (string, error) {
	id := uuid.New().String()
	if _, exists := repo.store[id]; exists {
		return "", errors.NewKeyAlreadyExistsError(id)
	}
	repo.store[id] = points
	return id, nil
}
