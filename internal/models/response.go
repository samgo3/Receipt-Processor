package models

// PointsResponse represents the response for points earned from a receipt.
type PointsResponse struct {
	// The total points earned from the receipt.
	Points int64 `json:"points"`
}

// ProcessReceiptResponse represents the response for a processed receipt.
type ProcessReceiptResponse struct {
	// The unique identifier for the receipt.
	Id string `json:"id"`
}
