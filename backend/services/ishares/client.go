package ishares

type EtfIssuerClient interface {
	GetFunds(limit, offset int) ([]FundResponse, error)
}
