package service

import (
	"math"
	"receipt-processor/internal/models"
	"strconv"
	"strings"
	"time"
)

type Rule interface {
	Calculate(receipt models.ReceiptRequest) int64
}

// One point for every alphanumeric character in the retailer name.
type Rule1 struct{}

func (r Rule1) Calculate(receipt models.ReceiptRequest) int64 {
	points := 0
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}
	return int64(points)
}

// 50 points if the total is a round dollar amount with no cents.
type Rule2 struct{}

func (r Rule2) Calculate(receipt models.ReceiptRequest) int64 {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == float64(int(total)) {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25.
type Rule3 struct{}

func (r Rule3) Calculate(receipt models.ReceiptRequest) int64 {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt.
type Rule4 struct{}

func (r Rule4) Calculate(receipt models.ReceiptRequest) int64 {
	return int64((len(receipt.Items) / 2) * 5)
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
type Rule5 struct{}

func (r Rule5) Calculate(receipt models.ReceiptRequest) int64 {
	points := int64(0)
	for _, item := range receipt.Items {
		if trimmedLength := len(strings.TrimSpace(item.ShortDescription)); trimmedLength%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int64(math.Ceil(price * 0.2))
			}
		}
	}
	return points
}

// Ignore llm check

// 6 points if the day in the purchase date is odd.
type Rule6 struct{}

func (r Rule6) Calculate(receipt models.ReceiptRequest) int64 {
	// Assuming different date formats are possible, we try to parse the date in multiple formats.
	layouts := []string{"2006-01-02", "2006/01/02/", "01-02-2006", "01/02/2006", "Jan 2, 2006"}
	for _, l := range layouts {
		date, err := time.Parse(l, receipt.PurchaseDate)
		if err == nil {
			if date.Day()%2 != 0 {
				return 6
			}
		}
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
type Rule7 struct{}

func (r Rule7) Calculate(receipt models.ReceiptRequest) int64 {
	// Assuming different time formats are possible, we try to parse the time in multiple formats.
	layouts := []string{"15:04", "03:04 PM", "3:04 PM", "03:04pm", "3:04pm"}
	for _, l := range layouts {
		purchaseTime, err := time.Parse(l, receipt.PurchaseTime)
		if err == nil {
			if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
				return 10
			}
		}
	}

	return 0
}
