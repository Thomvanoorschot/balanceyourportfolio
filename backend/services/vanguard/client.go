package vanguard

type EtfIssuerClient interface {
	GetFunds(ei []string) ([]byte, error)
}
type FigiClient interface {
	GetFigi(payload []FigiPayload) (r []FigiResp, err error)
}
