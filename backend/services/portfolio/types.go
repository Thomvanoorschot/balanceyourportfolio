package portfolio

import (
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
	ID    uuid.UUID `db:"portfolio.id"`
	Name  string    `db:"portfolio.name"`
	Items ListItems
}

func ConvertToModel(p *proto.Portfolio) Model {
	return Model{
		ID:    stringutils.ConvertToUUID(p.Id),
		Name:  p.Name,
		Items: ConvertToListItems(p.Entries),
	}
}
func (m *Model) ConvertToResponse() *proto.Portfolio {
	return &proto.Portfolio{
		Id:      m.ID.String(),
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

type ListItem struct {
	ID     uuid.UUID `db:"portfolio_fund.id"`
	Amount float64   `db:"portfolio_fund.amount"`
	FundID uuid.UUID `db:"fund.id"`
	Name   string    `db:"fund.name"`
}

func ConvertToListItem(pli *proto.PortfolioListItem) ListItem {
	return ListItem{
		ID:     stringutils.ConvertToUUID(pli.Id),
		FundID: stringutils.ConvertToUUID(pli.FundId),
		Amount: pli.Amount,
		Name:   pli.Name,
	}
}
func (li ListItem) ConvertToResponse() *proto.PortfolioListItem {
	return &proto.PortfolioListItem{
		Id:     li.ID.String(),
		FundId: li.FundID.String(),
		Name:   li.Name,
		Amount: li.Amount,
	}
}

type RelativeSectorWeighting struct {
	FundID              uuid.UUID
	FundName            string
	PortfolioFundAmount float64
	FundPrice           float64
	SectorWeightings    fund.SectorWeightings
}
type RelativeSectorWeightings []RelativeSectorWeighting
