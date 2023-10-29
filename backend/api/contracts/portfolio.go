package contracts

import (
	"github.com/google/uuid"
)

type Portfolio struct {
	ID    uuid.UUID           `json:"id"`
	Name  string              `json:"name"`
	Items []PortfolioListItem `json:"items"`
}

type PortfolioListItem struct {
	ID     uuid.UUID `json:"id"`
	FundID uuid.UUID `json:"fundId"`
	Name   string    `json:"name"`
	Amount float64   `json:"amount"`
}
