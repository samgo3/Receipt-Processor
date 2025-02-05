package validation

import (
	"errors"
	"fmt"
	"receipt-processor/internal/models"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

func validateRetailer(fl validator.FieldLevel) bool {
	retailer := fl.Field().String()
	return regexp.MustCompile(`^[\w\s\-&]+$`).MatchString(retailer)
}
func validateDescription(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	return regexp.MustCompile(`^[\w\s\-]+$`).MatchString(description)
}

func ValidateDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	// Assuming multiple date formats are possible, we try to parse the date in multiple formats.
	layout := []string{"2006-01-02", "2006/01/02", "01-02-2006", "01/02/2006", "Jan 2, 2006"}
	var err error
	for _, l := range layout {
		_, err = time.Parse(l, date)
		if err == nil {
			return true
		}
	}
	return false

}

func ValidateTime(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	// Assuming multiple time formats are possible, we try to parse the time in multiple formats.
	layouts := []string{"15:04", "03:04 PM", "3:04 PM", "03:04pm", "3:04pm"}
	var err error
	for _, layout := range layouts {
		_, err = time.Parse(layout, timeStr)
		if err == nil {
			return true
		}
	}
	return false
}

func ValidateNum(fl validator.FieldLevel) bool {
	total := fl.Field().String()
	return regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(total)
}

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

var customErrorMessages = map[string]string{
	"validateRetailer":    "The string contains invalid characters.",
	"validateDate":        "Invalid format or the value is invalid.",
	"validateTime":        "Invalid format or the value is invalid.",
	"validateNum":         "The number format is invalid. It should be a decimal with two decimal places.",
	"validateDescription": "The string contains invalid characters.",
	"required":            "The field is required.",
}

func GetErrorMessage(tag string) string {
	if msg, exists := customErrorMessages[tag]; exists {
		return msg
	}
	return "Validation failed."
}

func initValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("validateRetailer", validateRetailer)
	validate.RegisterValidation("validateDescription", validateDescription)
	validate.RegisterValidation("validateDate", ValidateDate)
	validate.RegisterValidation("validateTime", ValidateTime)
	validate.RegisterValidation("validateNum", ValidateNum)
}

func ValidateStruct(receipt models.ReceiptRequest) error {
	validateOnce.Do(initValidator)
	if err := validate.Struct(receipt); err != nil {
		var errorStr []string
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Field '%s': %s", e.Field(), GetErrorMessage(e.Tag()))
			errorStr = append(errorStr, message)
		}
		return errors.New(strings.Join(errorStr, "\n"))
	}
	return nil
}
