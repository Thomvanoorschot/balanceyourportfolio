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

type PortfolioDetails struct {
	FundInformation               []FundInformation               `json:"fundInformation"`
	Sectors                       []string                        `json:"sectors"`
	PortfolioFundSectorWeightings []PortfolioFundSectorWeightings `json:"portfolioFundSectorWeightings"`
}

type PortfolioFundSectorWeightings struct {
	FundName            string                `json:"fundName"`
	PercentageOfTotal   float64               `json:"percentageOfTotal"`
	FundSectorWeighting []FundSectorWeighting `json:"fundSectorWeighting"`
}
