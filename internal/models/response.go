package models

type PointsResponse struct {
	// The total points earned from the receipt.
	Points int64 `json:"points"`
}

type ProcessReceiptResponse struct {
	// The unique identifier for the receipt.
	Id string `json:"id"`
}
