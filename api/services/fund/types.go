package fund

import (
	"time"

	"github.com/google/uuid"
)

const FundHoldingsLimit = 20

type Details struct {
	Information      Information
	Holdings         []Holding
	Sectors          []SectorName
	SectorWeightings []SectorWeighting
}

type Holding struct {
	Ticker            string        `db:"holding.ticker"`
	Name              string        `db:"holding.name"`
	Type              IssueTypeName `db:"holding.type"`
	Sector            SectorName    `db:"holding.sector"`
	Amount            float64       `db:"fund_holding.amount"`
	PercentageOfTotal float64       `db:"fund_holding.percentage_of_total"`
	MarketValue       float64       `db:"fund_holding.market_value"`
}

type Fund struct {
	Id      uuid.UUID `db:"fund.id"`
	Name    string    `db:"fund.name"`
	Tickers []string
}

type Information struct {
	Name              string    `db:"fund.name"`
	OutstandingShares string    `db:"fund.outstanding_shares"`
	EffectiveDate     time.Time `db:"fund.effective_date"`
}

type HoldingsFilter struct {
	FundId     uuid.UUID  `json:"fundId"`
	SearchTerm string     `json:"searchTerm"`
	SectorName SectorName `json:"sectorName"`
	Limit      int64      `json:"limit"`
	Offset     int64      `json:"offset"`
}

type IssueTypeName string

const (
	Currency         IssueTypeName = "Currency"
	CommonStock      IssueTypeName = "Common Stock"
	ForwardContracts IssueTypeName = "Forward Contracts"
	Unknown          IssueTypeName = "Unknown"
)

type EffectiveShare struct {
	Ticker string
	Name   string
	Amount string
}
type SectorWeighting struct {
	SectorName SectorName `db:"holding.sector"`
	Percentage float64    `db:"percentage_sum"`
}

type SectorName string

const (
	AnySector                   SectorName = "Any sector"
	UnknownSector               SectorName = "Unknown"
	TechnologySector            SectorName = "Technology"
	HealthCareSector            SectorName = "HealthCare"
	FinancialsSector            SectorName = "Financials"
	RealEstateSector            SectorName = "RealEstate"
	EnergySector                SectorName = "Energy"
	MaterialsSector             SectorName = "Materials"
	ConsumerDiscretionarySector SectorName = "Consumer Discretionary"
	IndustrialsSector           SectorName = "Industrials"
	UtilitiesSector             SectorName = "Utilities"
	ConsumerStaplesSector       SectorName = "Consumer Staples"
	TelecommunicationSector     SectorName = "Telecommunication"
)
