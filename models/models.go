package models

type Stock struct {
	Symbol string `json:"symbol"`
	Price  float64 `json:"price"`
}


type HistoricalData struct {
		Date  string `json:"date"`
		Price float64 `json:"price"`
}
