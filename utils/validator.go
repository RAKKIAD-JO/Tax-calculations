package utils

import (
	"errors"
	"test-backend/models"
)

// ValidateTextRequest validates inputs from models.TextRequest and returns error when invalid
func ValidateTextRequest(req models.TextRequest) error {
	if req.TotalIncome < 0 {
		return errors.New("totalIncome must be non-negative")
	}
	if req.WHT < 0 {
		return errors.New("wht must be non-negative")
	}
	if req.WHT > req.TotalIncome {
		return errors.New("wht must not be greater than totalIncome")
	}
	if req.Allowances == nil {
		return errors.New("allowances must be an array")
	}
	for _, a := range req.Allowances {
		if a.Amount < 0 {
			return errors.New("allowance amounts must be non-negative")
		}
	}
	return nil
}