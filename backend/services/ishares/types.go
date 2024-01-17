package ishares

import "time"

type FundResponse struct {
	FundName           string
	Holdings           HoldingsResponse
	Currency           string
	ISIN               string
	TotalHoldings      int64
	Price              float64
	ExternalIdentifier string
	OutstandingShares  float64
	NetAssets          float64
	EffectiveDate      time.Time
}
type HoldingsResponse struct {
	AaData [][]interface{} `json:"aaData"`
}
type NumberValue struct {
	Display string  `json:"display"`
	Raw     float64 `json:"raw"`
}
