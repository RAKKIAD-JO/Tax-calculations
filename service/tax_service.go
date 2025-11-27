package service

import (
	"fmt"
	"test-backend/models"
)

const PERSONAL_DEDUCTION = 60000
const DONATION_LIMIT = 100000

var taxBrackets = []struct {
	Limit float64
	Rate  float64
}{
	{150000, 0.0},
	{500000, 0.10},
	{1000000, 0.15},
	{2000000, 0.20},
	{999999999, 0.35},
}

func CalculateTax(req models.TaxRequest) models.TaxResponse {
	taxableIncome := req.TotalIncome - PERSONAL_DEDUCTION

	// หักค่าลดหย่อน (บริจาค)
	for _, a := range req.Allowances {
		if a.AllowanceType == "donation" {
			d := a.Amount
			if d > DONATION_LIMIT {
				d = DONATION_LIMIT
			}
			taxableIncome -= d
		}
	}

	if taxableIncome < 0 {
		taxableIncome = 0
	}

	totalTax := 0.0
	remain := taxableIncome
	prev := 0.0

	var levels []models.TaxLevel

	for _, b := range taxBrackets {
		size := b.Limit - prev

		if remain <= 0 {
			levels = append(levels, zeroLevel(prev, b.Limit))
			prev = b.Limit
			continue
		}

		taxable := min(remain, size)
		tax := taxable * b.Rate

		totalTax += tax

		levels = append(levels, models.TaxLevel{
			Level: formatLevel(prev, b.Limit),
			Tax:   tax,
		})

		remain -= taxable
		prev = b.Limit
	}

	final := totalTax - req.WHT

	return models.TaxResponse{
		Tax:       final,
		TaxLevels: levels,
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func zeroLevel(a, b float64) models.TaxLevel {
	return models.TaxLevel{
		Level: formatLevel(a, b),
		Tax:   0,
	}
}

func formatLevel(a, b float64) string {
	if b >= 2000000 {
		return "2,000,001 ขึ้นไป"
	}
	return fmt.Sprintf("%.0f-%.0f", a+1, b)
}
