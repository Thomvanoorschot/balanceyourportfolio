package portfolio

import (
	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/generated/proto"
	"etfinsight/services/fund"
	"etfinsight/utils/stringutils"

	"github.com/google/uuid"
)

type Models []Model

func (m Models) ConvertToResponse() *proto.PortfoliosResponse {
	resp := &proto.PortfoliosResponse{}
	for i := range m {
		resp.Entries = append(resp.Entries, m[i].ConvertToResponse())
	}
	return resp
}

type Model struct {
	Id    uuid.UUID `db:"portfolio.id"`
	Name  string    `db:"portfolio.name"`
	Items ListItems
}

func ConvertToModel(p *proto.Portfolio) Model {
	return Model{
		Id:    stringutils.ConvertToUUID(p.Id),
		Name:  p.Name,
		Items: ConvertToListItems(p.Entries),
	}
}
func (m *Model) ConvertToResponse() *proto.Portfolio {
	return &proto.Portfolio{
		Id:      m.Id.String(),
		Name:    m.Name,
		Entries: m.Items.ConvertToResponse(),
	}
}

type ListItems []ListItem

func ConvertToListItems(pli []*proto.PortfolioListItem) []ListItem {
	li := make([]ListItem, len(pli))
	for i := range pli {
		li[i] = ConvertToListItem(pli[i])
	}
	return li
}
func (li ListItems) ConvertToResponse() []*proto.PortfolioListItem {
	pli := make([]*proto.PortfolioListItem, len(li))
	for i := range li {
		pli[i] = li[i].ConvertToResponse()
	}
	return pli
}
func (li ListItems) ConvertToDbModel(portfolioId uuid.UUID) []model.PortfolioFund {
	var pfs []model.PortfolioFund

	for i := range li {
		if li[i].FundID == uuid.Nil {
			continue
		}
		if li[i].Id == uuid.Nil {
			li[i].Id = uuid.New()
		}
		pf := model.PortfolioFund{
			ID:          li[i].Id,
			PortfolioID: &portfolioId,
			FundID:      &li[i].FundID,
			Amount:      &li[i].Amount,
		}
		pfs = append(pfs, pf)
	}
	return pfs
}

type ListItem struct {
	Id     uuid.UUID `db:"portfolio_fund.id"`
	Amount float64   `db:"portfolio_fund.amount"`
	FundID uuid.UUID `db:"fund.id"`
	Name   string    `db:"fund.name"`
}

func ConvertToListItem(pli *proto.PortfolioListItem) ListItem {
	return ListItem{
		Id:     stringutils.ConvertToUUID(pli.Id),
		FundID: stringutils.ConvertToUUID(pli.FundId),
		Amount: pli.Amount,
		Name:   pli.Name,
	}
}
func (li ListItem) ConvertToResponse() *proto.PortfolioListItem {
	return &proto.PortfolioListItem{
		Id:     li.Id.String(),
		FundId: li.FundID.String(),
		Name:   li.Name,
		Amount: li.Amount,
	}
}

type RelativeSectorWeightings []RelativeSectorWeighting
type RelativeSectorWeighting struct {
	FundID           uuid.UUID
	FundName         string
	SectorWeightings SectorWeightings
}

func (rsw RelativeSectorWeightings) ConvertToResponse(ratio map[uuid.UUID]float64) map[string]*proto.PortfolioFundSectorWeighting {
	pfsw := map[string]*proto.PortfolioFundSectorWeighting{}
	for i := range rsw {
		rsw[i].ConvertToResponse(ratio, pfsw)
	}
	return pfsw
}
func (rsw RelativeSectorWeighting) ConvertToResponse(ratio map[uuid.UUID]float64, pfsw map[string]*proto.PortfolioFundSectorWeighting) {
	for _, sw := range rsw.SectorWeightings {
		ratiodPerentage := sw.Percentage * ratio[rsw.FundID]
		sector, ok := pfsw[string(sw.SectorName)]
		if !ok {
			newEntry := &proto.PortfolioFundSectorWeighting{
				TotalPercentage: ratiodPerentage,
				FundSectorWeighting: []*proto.PortfolioFundSectorWeightingEntry{
					{
						FundId:     rsw.FundID.String(),
						FundName:   rsw.FundName,
						Percentage: ratiodPerentage,
					},
				},
			}
			pfsw[string(sw.SectorName)] = newEntry
			sector = newEntry
			continue
		}
		sector.TotalPercentage += ratiodPerentage
		sector.FundSectorWeighting = append(sector.FundSectorWeighting, &proto.PortfolioFundSectorWeightingEntry{
			FundId:     rsw.FundID.String(),
			FundName:   rsw.FundName,
			Percentage: ratiodPerentage,
		})
	}
}

type SectorWeightings []SectorWeighting
type SectorWeighting struct {
	fund.SectorWeighting
}

type FundHoldings []FundHolding
type FundHolding struct {
	Ticker               string
	HoldingName          string
	HoldingId            uuid.UUID
	CumulativePercentage float64
	Funds                FundsHoldingEntries
}

func (fh FundHoldings) ConvertToResponse() []*proto.PortfolioFundHolding {
	pfh := make([]*proto.PortfolioFundHolding, len(fh))
	for i := range fh {
		pfh[i] = fh[i].ConvertToResponse()
	}
	return pfh
}
func (rsw FundHolding) ConvertToResponse() *proto.PortfolioFundHolding {
	return &proto.PortfolioFundHolding{
		Ticker:               rsw.Ticker,
		HoldingId:            rsw.HoldingId.String(),
		HoldingName:          rsw.HoldingName,
		CumulativePercentage: rsw.CumulativePercentage,
		Funds:                rsw.Funds.ConvertToResponse(),
	}
}

type FundsHoldingEntries []FundHoldingEntry
type FundHoldingEntry struct {
	FundId          uuid.UUID
	RatiodPerentage float64
}

func (fh FundsHoldingEntries) ConvertToResponse() []*proto.PortfolioFundHoldingEntry {
	pfhe := make([]*proto.PortfolioFundHoldingEntry, len(fh))
	for i := range fh {
		pfhe[i] = fh[i].ConvertToResponse()
	}
	return pfhe
}
func (fhe FundHoldingEntry) ConvertToResponse() *proto.PortfolioFundHoldingEntry {
	return &proto.PortfolioFundHoldingEntry{
		FundId:           fhe.FundId.String(),
		RatiodPercentage: fhe.RatiodPerentage,
	}
}
