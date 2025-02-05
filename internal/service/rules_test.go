package service

import (
	"receipt-processor/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule1_Calculate(t *testing.T) {
	receipt := models.ReceiptRequest{Retailer: "Retailer123 & rt"}
	rule := Rule1{}
	points := rule.Calculate(receipt)
	assert.Equal(t, int64(13), points)

	receipt.Retailer = "Retailer"
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(8), points)

}

func TestRule2_Calculate(t *testing.T) {
	receipt := models.ReceiptRequest{Total: "100.00"}
	rule := Rule2{}
	points := rule.Calculate(receipt)
	assert.Equal(t, int64(50), points)

	receipt.Total = "100.01"
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(0), points)
}

func TestRule3_Calculate(t *testing.T) {
	receipt := models.ReceiptRequest{Total: "100.25"}
	rule := Rule3{}
	points := rule.Calculate(receipt)
	assert.Equal(t, int64(25), points)

	receipt.Total = "100.10"
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(0), points)
}

func TestRule4_Calculate(t *testing.T) {
	receipt := models.ReceiptRequest{Items: []models.Item{{}, {}}}
	rule := Rule4{}
	points := rule.Calculate(receipt)
	assert.Equal(t, int64(5), points)

	receipt.Items = []models.Item{{}, {}, {}}
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(5), points)

	receipt.Items = []models.Item{{}, {}, {}, {}}
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(10), points)
}

func TestRule5_Calculate(t *testing.T) {
	receipt := models.ReceiptRequest{Items: []models.Item{
		{ShortDescription: "abc", Price: "1.00"},
		{ShortDescription: "abcd", Price: "2.00"},
	}}
	rule := Rule5{}
	points := rule.Calculate(receipt)
	assert.Equal(t, int64(1), points)

	receipt.Items = []models.Item{
		{ShortDescription: "abc", Price: "1.00"},
		{ShortDescription: "abcdef ", Price: "12.00"},
	}
	points = rule.Calculate(receipt)
	assert.Equal(t, int64(4), points)
}

func TestRule6_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected int64
	}{
		{
			name:     "Odd day",
			date:     "2023-10-01",
			expected: 6,
		},
		{
			name:     "Even day",
			date:     "2023-10-02",
			expected: 0,
		},
		{
			name:     "Date format with slashes",
			date:     "2023/10/02",
			expected: 0,
		},
		{
			name:     "MM-DD-YYYY format",
			date:     "10-01-2023",
			expected: 6,
		},
		{name: "MM/DD/YYYY format",

			date:     "10/01/2023",
			expected: 6,
		},
		{
			name:     "Month name format",
			date:     "Oct 1, 2023",
			expected: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			receipt := models.ReceiptRequest{PurchaseDate: test.date}
			rule := Rule6{}
			points := rule.Calculate(receipt)
			assert.Equal(t, int64(test.expected), points)
		})
	}
}

func TestRule7_Calculate(t *testing.T) {
	tests := []struct {
		name         string
		purchaseTime string
		expected     int64
	}{
		{
			name:         "Time in 24-hour format within range",
			purchaseTime: "14:30",
			expected:     10,
		},
		{
			name:         "Time in 12-hour format within range",
			purchaseTime: "03:00 PM",
			expected:     10,
		},
		{
			name:         "Time in 24-hour format out of range",
			purchaseTime: "16:00",
			expected:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			receipt := models.ReceiptRequest{PurchaseTime: test.purchaseTime}
			rule := Rule7{}
			points := rule.Calculate(receipt)
			assert.Equal(t, int64(test.expected), points)
		})
	}

}
