package model

type BaseCurrency struct {
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"currency"`
}
