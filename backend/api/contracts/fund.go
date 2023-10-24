package contracts

import (
	"time"

	"github.com/google/uuid"
)

type FundDetails struct {
	Information      FundInformation       `json:"information"`
	Sectors          []string              `json:"sectors"`
	SectorWeightings []FundSectorWeighting `json:"sectorWeightings"`
}

type FundSectorWeighting struct {
	SectorName string  `json:"sectorName"`
	Percentage float64 `json:"percentage"`
}

type FundHolding struct {
	Ticker            string  `json:"ticker"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	Sector            string  `json:"sector"`
	Amount            float64 `json:"amount"`
	PercentageOfTotal float64 `json:"percentageOfTotal"`
	MarketValue       float64 `json:"marketValue"`
}

type Fund struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Tickers []string  `json:"tickers"`
}

type FundInformation struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	OutstandingShares string    `json:"outstandingShares"`
	EffectiveDate     time.Time `json:"effectiveDate"`
}

type FundHoldingsFilter struct {
	FundID     uuid.UUID `json:"fundId"`
	SearchTerm string    `json:"searchTerm"`
	SectorName string    `json:"sectorName"`
	Limit      int64     `json:"limit"`
	Offset     int64     `json:"offset"`
}
