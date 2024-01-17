package ishares

type EtfIssuerClient interface {
	GetFunds() ([]FundResponse, error)
	GetHoldings(url string) (HoldingsResponse, error)
}
