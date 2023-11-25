package portfolio

import (
	"etfinsight/api/contracts"
	"etfinsight/services/fund"

	"github.com/google/uuid"
)

type Models []Model

func (m Models) ConvertToResponse() []contracts.Portfolio {
	p := make([]contracts.Portfolio, len(m))
	for i := range m {
		p[i] = m[i].ConvertToResponse()
	}
	return p
}

type Model struct {
	ID    uuid.UUID `db:"portfolio.id"`
	Name  string    `db:"portfolio.name"`
	Items ListItems
}

func ConvertToModel(p contracts.Portfolio) Model {
	return Model{
		ID:    p.ID,
		Name:  p.Name,
		Items: ConvertToListItems(p.Items),
	}
}
func (m *Model) ConvertToResponse() contracts.Portfolio {
	return contracts.Portfolio{
		ID:    m.ID,
		Name:  m.Name,
		Items: m.Items.ConvertToResponse(),
	}
}

type ListItems []ListItem

func ConvertToListItems(pli []contracts.PortfolioListItem) []ListItem {
	li := make([]ListItem, len(pli))
	for i := range pli {
		li[i] = ConvertToListItem(pli[i])
	}
	return li
}
func (li ListItems) ConvertToResponse() []contracts.PortfolioListItem {
	pli := make([]contracts.PortfolioListItem, len(li))
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

func ConvertToListItem(pli contracts.PortfolioListItem) ListItem {
	return ListItem{
		ID:     pli.ID,
		FundID: pli.FundID,
		Amount: pli.Amount,
		Name:   pli.Name,
	}
}
func (li ListItem) ConvertToResponse() contracts.PortfolioListItem {
	return contracts.PortfolioListItem{
		ID:     li.ID,
		FundID: li.FundID,
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
