package validation

import (
	"receipt-processor/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitValidator(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Sample Retailer",
		PurchaseDate: "2023-10-01",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.NoError(t, err, "Expected no validation errors, but got: %v", err)
}

func TestInvalidRetailer(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Invalid@Retailer!",
		PurchaseDate: "2023-10-01",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for invalid retailer, but got none")
}

func TestInvalidDate(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Sample Retailer",
		PurchaseDate: "2023-02-30",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for invalid date, but got none")
}

func TestInvalidTime(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Sample Retailer",
		PurchaseDate: "2023-10-01",
		PurchaseTime: "25:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for invalid time, but got none")
}

func TestInvalidTotal(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Simple retailer",
		PurchaseDate: "2023-10-10",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.0",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for invalid total, but got none")

}

func TestShortDescritpion(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Simple retailer",
		PurchaseDate: "2023-10-10",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Dadf&@#@",
				Price:            "10.0",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for empty short description, but got none")
}

func TestItemPrice(t *testing.T) {
	receipt := models.ReceiptRequest{

		Retailer:     "Simple retailer",
		PurchaseDate: "2023-10-10",
		PurchaseTime: "14:00",
		Items: []models.Item{
			{
				ShortDescription: "Item 1",
				Price:            "10.0",
			},
		},
		Total: "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for empty short description, but got none")

}

func TestEmptyItems(t *testing.T) {
	receipt := models.ReceiptRequest{
		Retailer:     "Simple retailer",
		PurchaseDate: "2023-10-10",
		PurchaseTime: "14:00",
		Items:        []models.Item{},
		Total:        "30.00",
	}
	err := ValidateStruct(receipt)
	assert.Error(t, err, "Expected validation error for empty items, but got none")
}
