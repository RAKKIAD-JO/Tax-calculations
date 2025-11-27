package models

type Allowance struct {
	Allowance string  `json:"allowanceType"`
	Amount    float64 `json:"amount"`
}

type TextRequest struct {
	TotalIncome float64     `json:"totalIncome"`
	WHT         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}

type TextLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type TaxResponse struct {
	Tax       float64     `json:"tax"`
	TaxLevels []TextLevel `json:"taxLevels,omitempty"`
}
