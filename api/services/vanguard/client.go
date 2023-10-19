package vanguard

type EtfIssuerClient interface {
	GetFunds(ei []string) ([]byte, error)
}
