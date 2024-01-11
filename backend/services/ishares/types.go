package ishares

type HoldingsResponse struct {
	AaData [][]interface{} `json:"aaData"`
}
type NumberValue struct {
	Display string  `json:"display"`
	Raw     float64 `json:"raw"`
}
