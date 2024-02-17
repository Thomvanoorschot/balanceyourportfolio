package ishares

type EtfIssuerClient interface {
	GetFunds(limit, offset int) ([]FundResponse, error)
}
type FigiClient interface {
	GetFigi(payload []FigiPayload) (r []FigiResp, err error)
}
