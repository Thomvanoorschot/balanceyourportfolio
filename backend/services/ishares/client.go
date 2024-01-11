package ishares

type EtfIssuerClient interface {
	GetFunds() ([]byte, error)
}
