package fund

import (
	"time"

	"etfinsight/generated/proto"
	"etfinsight/utils/stringutils"

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
	OutstandingShares float64   `db:"fund.outstanding_shares"`
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

func ConvertToHoldingsFilter(f *proto.FilterHoldingsRequest) HoldingsFilter {
	return HoldingsFilter{
		FundID:     stringutils.ConvertToUUID(f.FundId),
		SearchTerm: f.SearchTerm,
		SectorName: f.SectorName,
		Limit:      f.Limit,
		Offset:     f.Offset,
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
func (f Funds) ConvertToResponse() *proto.SearchFundsResponse {
	resp := &proto.SearchFundsResponse{}

	for i := range f {
		resp.Entries = append(resp.Entries, f[i].ConvertToResponse())
	}
	return resp
}
func (f Fund) ConvertToResponse() *proto.SearchFundsEntry {
	return &proto.SearchFundsEntry{
		Id:      f.ID.String(),
		Name:    f.Name,
		Tickers: f.Tickers,
	}
}
func (h Holdings) ConvertToResponse() *proto.HoldingsListResponse {
	resp := &proto.HoldingsListResponse{}

	for i := range h {
		resp.Entries = append(resp.Entries, h[i].ConvertToResponse())
	}
	return resp
}

func (h Holding) ConvertToResponse() *proto.HoldingsResponse {
	return &proto.HoldingsResponse{
		Ticker:            h.Ticker,
		Name:              h.Name,
		Type:              string(h.Type),
		Sector:            string(h.Sector),
		Amount:            h.Amount,
		PercentageOfTotal: h.PercentageOfTotal,
		MarketValue:       h.MarketValue,
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
