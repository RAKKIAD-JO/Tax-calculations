package utils

import (
	"errors"
	"test-backend/models"
)

func ValidateTextRequest(req models.TaxRequest) error {
	if req.TotalIncome < 0 {
		return errors.New("total income cannot be negative")
	}
	if req.WHT < 0 {
		return errors.New("WHT cannot be negative")
	}
	if req.WHT > req.TotalIncome {
		return errors.New("WHT cannot be greater than total income")
	}
	return nil
}
