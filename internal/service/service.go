package service

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/repository"
)

type ReceiptService struct {
	repo repository.Repository
}

func NewReceiptService(repo repository.Repository) *ReceiptService {
	return &ReceiptService{repo: repo}
}

// ProcessReceipt processes a receipt and returns the unique identifier for the receipt.
func (s *ReceiptService) ProcessReceipt(receipt models.ReceiptRequest) (string, error) {
	points := s.CalculatePoints(receipt)
	id, err := s.repo.AddEntry(points)
	if err != nil {
		return "", err
	}
	return id, nil

}

// GetPointsById retrieves the points for a given receipt ID.
func (s *ReceiptService) GetPointsById(id string) (int64, error) {
	points, err := s.repo.GetById(id)
	if err != nil {
		return 0, err
	}
	return points, nil
}

func (s *ReceiptService) CalculatePoints(receipt models.ReceiptRequest) int64 {
	points := int64(0)
	rules := []Rule{
		Rule1{},
		Rule2{},
		Rule3{},
		Rule4{},
		Rule5{},
		Rule6{},
		Rule7{},
	}

	for _, rule := range rules {
		points += rule.Calculate(receipt)
	}
	return points
}
