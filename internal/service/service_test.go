package service

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.ReceiptRequest
		expected int64
	}{
		{
			name: "Points Test 1",
			receipt: models.ReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28.00,
		},

		{
			name: "Points Test 2",
			receipt: models.ReceiptRequest{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109.00,
		},
	}

	repo := repository.KVRepo{} // Mock testing repository

	service := NewReceiptService(&repo)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			point := service.CalculatePoints(test.receipt)
			assert.Equal(t, test.expected, point, "Failed tested")
		})
	}
}
