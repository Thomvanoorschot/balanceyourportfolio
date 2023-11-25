package fund

import (
	"time"

	"etfinsight/api/contracts"

	"github.com/google/uuid"
)

type Holdings []Holding
type Holding struct {
	Ticker            string        `db:"holding.ticker"`
	Name              string        `db:"holding.name"`
	Type              IssueTypeName `db:"holding.type"`
	Sector            SectorName    `db:"holding.sector"`
	Amount            float64       `db:"fund_holding.amount"`
	PercentageOfTotal float64       `db:"fund_holding.percentage_of_total"`
	MarketValue       float64       `db:"fund_holding.market_value"`
}

type Funds []Fund
type Fund struct {
	ID      uuid.UUID `db:"fund.id"`
	Name    string    `db:"fund.name"`
	Tickers []string
}

type InformationList []Information
type Information struct {
	ID                uuid.UUID `db:"fund.id"`
	Name              string    `db:"fund.name"`
	OutstandingShares string    `db:"fund.outstanding_shares"`
	EffectiveDate     time.Time `db:"fund.effective_date"`
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
type SectorWeightings []SectorWeighting
type SectorWeighting struct {
	SectorName SectorName `db:"holding.sector"`
	Percentage float64    `db:"percentage_sum"`
}

type SectorNames []SectorName
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

type HoldingsFilter struct {
	FundID     uuid.UUID
	SearchTerm string
	SectorName string
	Limit      int64
	Offset     int64
}

func ConvertToHoldingsFilter(f contracts.FundHoldingsFilter) HoldingsFilter {
	return HoldingsFilter{
		FundID:     f.FundID,
		SearchTerm: f.SearchTerm,
		SectorName: f.SectorName,
		Limit:      f.Limit,
		Offset:     f.Offset,
	}
}
func (il InformationList) ConvertToResponse() []contracts.FundInformation {
	fi := make([]contracts.FundInformation, len(il))
	for i := range il {
		fi[i] = il[i].ConvertToResponse()
	}
	return fi
}
func (i Information) ConvertToResponse() contracts.FundInformation {
	return contracts.FundInformation{
		ID:                i.ID,
		Name:              i.Name,
		OutstandingShares: i.OutstandingShares,
		EffectiveDate:     i.EffectiveDate,
	}
}
func (f Funds) ConvertToResponse() []contracts.Fund {
	cf := make([]contracts.Fund, len(f))
	for i := range f {
		cf[i] = f[i].ConvertToResponse()
	}
	return cf
}
func (f Fund) ConvertToResponse() contracts.Fund {
	return contracts.Fund{
		ID:      f.ID,
		Name:    f.Name,
		Tickers: f.Tickers,
	}
}
func (h Holdings) ConvertToResponse() []contracts.FundHolding {
	fh := make([]contracts.FundHolding, len(h))
	for i := range h {
		fh[i] = h[i].ConvertToResponse()
	}
	return fh
}

func (h Holding) ConvertToResponse() contracts.FundHolding {
	return contracts.FundHolding{
		Ticker:            h.Ticker,
		Name:              h.Name,
		Type:              string(h.Type),
		Sector:            string(h.Sector),
		Amount:            h.Amount,
		PercentageOfTotal: h.PercentageOfTotal,
		MarketValue:       h.MarketValue,
	}
}

func (sw SectorWeightings) ConvertToResponse() []contracts.FundSectorWeighting {
	fsw := make([]contracts.FundSectorWeighting, len(sw))
	for i := range sw {
		fsw[i] = sw[i].ConvertToResponse()
	}
	return fsw
}
func (sw SectorWeighting) ConvertToResponse() contracts.FundSectorWeighting {
	return contracts.FundSectorWeighting{
		SectorName: string(sw.SectorName),
		Percentage: sw.Percentage,
	}
}

func (sn SectorNames) ConvertToResponse() []string {
	res := make([]string, len(sn))
	for i := range sn {
		res[i] = string(sn[i])
	}
	return res
}
