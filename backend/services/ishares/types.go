package ishares

import "time"

type FundResponse struct {
	FundName           string
	Holdings           HoldingsResponse
	Currency           string
	ISIN               string
	Price              float64
	ExternalIdentifier string
	Ticker             string
	OutstandingShares  float64
	NetAssets          float64
	EffectiveDate      time.Time
	Tickers            []string
	HoldingsTableIndex map[string]int
}
type HoldingsResponse struct {
	AaData [][]interface{} `json:"aaData"`
}
type NumberValue struct {
	Display string  `json:"display"`
	Raw     float64 `json:"raw"`
}
type FigiResp struct {
	Warning string      `json:"warning"`
	Data    []FigiValue `json:"data"`
}
type FigiValue struct {
	Figi           string  `json:"figi"`
	ShareClassFigi *string `json:"shareClassFIGI"`
	Name           string  `json:"name"`
	Ticker         string  `json:"ticker"`
}

type FigiPayload struct {
	IdType  string `json:"idType"`
	IdValue string `json:"idValue"`
}
