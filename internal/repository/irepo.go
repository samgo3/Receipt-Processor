package repository

// Repository Interface
type Repository interface {
	GetById(id string) (int64, error)
	AddEntry(points int64) (string, error)
}
