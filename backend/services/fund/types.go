package fund

import (
	"time"

	"etfinsight/generated/proto"

	"github.com/google/uuid"
)

type Holdings []Holding
type Holding struct {
	Id                string      `db:"holding.id"`
	Ticker            string      `db:"holding.ticker"`
	Name              string      `db:"holding.name"`
	Type              HoldingType `db:"holding.type"`
	Sector            SectorName  `db:"holding.sector"`
	Amount            float64     `db:"fund_holding.amount"`
	PercentageOfTotal float64     `db:"fund_holding.percentage_of_total"`
	MarketValue       float64     `db:"fund_holding.market_value"`
}

type Funds []Fund
type Fund struct {
	ID        uuid.UUID `db:"fund.id"`
	Name      string    `db:"fund.name"`
	Currency  string
	MarketCap float64
	Provider  string
	Tickers   []string
}

type InformationList []Information
type Information struct {
	ID                uuid.UUID `db:"fund.id"`
	Name              string    `db:"fund.name"`
	OutstandingShares float64   `db:"fund.outstanding_shares"`
	EffectiveDate     time.Time `db:"fund.effective_date"`
}

type HoldingType string

const (
	CashType          HoldingType = "Cash"
	Stocks            HoldingType = "Stocks"
	UnknownType       HoldingType = "Unknown"
	BondsType         HoldingType = "Bonds"
	MoneyMarketType   HoldingType = "Money Market"
	TreasuryType      HoldingType = "Treasuries"
	ClosedEndFundType HoldingType = "Closed End Fund"
	FuturesType       HoldingType = "Futures"
	NotesType         HoldingType = "Notes"
	MutualFundType    HoldingType = "Mutual Fund"
	FixedIncomeType   HoldingType = "Fixed Income"
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
	UnknownSector                SectorName = "Unknown"
	TechnologySector             SectorName = "Technology"
	HealthCareSector             SectorName = "HealthCare"
	FinancialsSector             SectorName = "Financials"
	RealEstateSector             SectorName = "RealEstate"
	EnergySector                 SectorName = "Energy"
	MaterialsSector              SectorName = "Materials"
	ConsumerDiscretionarySector  SectorName = "Consumer Discretionary"
	ConsumerCyclicalSector       SectorName = "Consumer Cyclical"
	IndustrialsSector            SectorName = "Industrials"
	UtilitiesSector              SectorName = "Utilities"
	ConsumerStaplesSector        SectorName = "Consumer Staples"
	TelecommunicationSector      SectorName = "Telecommunication"
	BondsSector                  SectorName = "Bonds"
	NotesSector                  SectorName = "Notes"
	CashSector                   SectorName = "Cash"
	InsuranceSector              SectorName = "Insurance"
	MortgageBackedSecuritySector SectorName = "Mortgage Backed Security"
)

type FundsFilter struct {
	SearchTerm string
	Providers  []string
	Limit      int64
	Offset     int64
}
type HoldingsFilter struct {
	FundId          uuid.UUID
	SearchTerm      string
	SelectedSectors []string
	Limit           int64
	Offset          int64
}

type OverlappingFunds struct {
	FundOneName                       string
	FundTwoName                       string
	TotalOverlappingPercentage        float64
	OverlappingHoldingsCount          int64
	FundOneHoldingCount               int64
	FundOneOverlappingCountPercentage float64
	FundTwoHoldingCount               int64
	FundTwoOverlappingCountPercentage float64
}

type OverlappingHoldings []OverlappingHolding
type OverlappingHolding struct {
	HoldingId             uuid.UUID
	HoldingName           string
	OverlappingPercentage float64
	FundOnePercentage     float64
	FundTwoPercentage     float64
}

func (h OverlappingHoldings) ConvertToResponse() []*proto.OverlappingHolding {
	oh := make([]*proto.OverlappingHolding, len(h))

	for i := range h {
		oh[i] = h[i].ConvertToResponse()
	}
	return oh
}
func (h OverlappingHolding) ConvertToResponse() *proto.OverlappingHolding {
	return &proto.OverlappingHolding{
		HoldingId:             h.HoldingId.String(),
		HoldingName:           h.HoldingName,
		OverlappingPercentage: h.OverlappingPercentage,
		FundOnePercentage:     h.FundOnePercentage,
		FundTwoPercentage:     h.FundTwoPercentage,
	}
}
func (il InformationList) ConvertToResponse() []*proto.FundInformation {
	fi := make([]*proto.FundInformation, len(il))
	for i := range il {
		fi[i] = il[i].ConvertToResponse()
	}
	return fi
}
func (i Information) ConvertToResponse() *proto.FundInformation {
	return &proto.FundInformation{
		Id:                i.ID.String(),
		Name:              i.Name,
		OutstandingShares: i.OutstandingShares,
		EffectiveDate:     i.EffectiveDate.String(),
	}
}
func (f Funds) ConvertToResponse() []*proto.FilterFundsResponseEntry {
	var resp []*proto.FilterFundsResponseEntry

	for i := range f {
		resp = append(resp, f[i].ConvertToResponse())
	}
	return resp
}
func (f Fund) ConvertToResponse() *proto.FilterFundsResponseEntry {
	return &proto.FilterFundsResponseEntry{
		Id:        f.ID.String(),
		Name:      f.Name,
		Tickers:   f.Tickers,
		Currency:  f.Currency,
		MarketCap: f.MarketCap,
		Provider:  f.Provider,
	}
}
func (h Holdings) ConvertToResponse() []*proto.FundHolding {
	fh := make([]*proto.FundHolding, len(h))

	for i := range h {
		fh[i] = h[i].ConvertToResponse()
	}
	return fh
}

func (h Holding) ConvertToResponse() *proto.FundHolding {
	return &proto.FundHolding{
		Ticker:               h.Ticker,
		HoldingId:            h.Id,
		HoldingName:          h.Name,
		CumulativePercentage: h.PercentageOfTotal,
	}
}

func (sw SectorWeightings) ConvertToResponse() []*proto.FundSectorWeighting {
	fsw := make([]*proto.FundSectorWeighting, len(sw))
	for i := range sw {
		fsw[i] = sw[i].ConvertToResponse()
	}
	return fsw
}
func (sw SectorWeighting) ConvertToResponse() *proto.FundSectorWeighting {
	return &proto.FundSectorWeighting{
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
