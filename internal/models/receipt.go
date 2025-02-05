package models

// Item represents an item in a receipt.
type Item struct {
	// The short description of the item.
	// Required: true
	// pattern: ^[\s\w-]+$
	ShortDescription string `json:"shortDescription" validate:"required,validateDescription"`
	// The price of the item.
	// Required: true
	// pattern: ^\d+\.\d{2}$
	Price string `json:"price" validate:"required,validateNum"`
}

// ReceiptRequest represents a request to process a receipt.
type ReceiptRequest struct {
	// The name of the retailer.
	// Required: true
	// pattern: ^[\w\s\-&]+$
	Retailer string `json:"retailer" validate:"required,validateRetailer"`
	// The date of purchase.
	// Required: true
	PurchaseDate string `json:"purchaseDate" validate:"required,validateDate"`
	// The time of purchase.
	// Required: true
	PurchaseTime string `json:"purchaseTime" validate:"required,validateTime"`
	// The list of items in the receipt.
	// Required: true
	Items []Item `json:"items" validate:"required,min=1,dive,required"`
	// The total amount of the receipt.
	// Required: true
	// pattern: ^\d+\.\d{2}$
	Total string `json:"total" validate:"required,validateNum"`
}
