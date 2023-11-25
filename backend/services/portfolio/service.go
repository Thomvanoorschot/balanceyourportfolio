package portfolio

import (
	"context"
	"sort"
	"time"

	"etfinsight/api/contracts"
	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"
	"etfinsight/utils/concurrencyutils"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPortfolioDetails(ctx context.Context, portfolioID uuid.UUID) (contracts.PortfolioDetails, error) {
	portfolioSectorCh := concurrencyutils.Async2(func() (fund.SectorNames, error) {
		return s.repo.GetPortfolioFundSectors(ctx, portfolioID)
	})
	informationCh := concurrencyutils.Async2(func() (fund.InformationList, error) {
		return s.repo.GetPortfolioFunds(ctx, portfolioID)
	})
	relativeWeightings := concurrencyutils.Async2(func() (RelativeSectorWeightings, error) {
		return s.repo.GetPortfolioFundRelativeWeightings(ctx, portfolioID)
	})
	portfolioSectorResult := <-portfolioSectorCh
	informationResult := <-informationCh
	relativeWeightingsResult := <-relativeWeightings
	var portfolioFundSectorWeightings []contracts.PortfolioFundSectorWeightings
	cumulativeSectorWeightings := map[string]float64{}
	var cumulativeValue float64

	for _, rw := range relativeWeightingsResult.Value {
		cumulativeValue += rw.PortfolioFundAmount * rw.FundPrice
		sw := contracts.PortfolioFundSectorWeightings{
			FundName: rw.FundName,
		}
		for _, sectorWeighting := range rw.SectorWeightings {
			sw.FundSectorWeighting = append(sw.FundSectorWeighting, contracts.FundSectorWeighting{
				SectorName: string(sectorWeighting.SectorName),
				Percentage: sectorWeighting.Percentage,
			})
		}
		portfolioFundSectorWeightings = append(portfolioFundSectorWeightings, sw)
	}
	for rwi, rw := range relativeWeightingsResult.Value {
		percentageOfTotal := (rw.PortfolioFundAmount * rw.FundPrice) / cumulativeValue
		portfolioFundSectorWeightings[rwi].PercentageOfTotal = percentageOfTotal
		for _, sectorWeighting := range rw.SectorWeightings {
			for _, weighting := range portfolioFundSectorWeightings[0].FundSectorWeighting {
				if weighting.SectorName == string(sectorWeighting.SectorName) {
					cumulativeSectorWeightings[weighting.SectorName] += sectorWeighting.Percentage * percentageOfTotal
					continue
				}
			}
		}
	}
	keys := make([]string, 0, len(cumulativeSectorWeightings))
	for key := range cumulativeSectorWeightings {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cumulativeSectorWeightings[keys[i]] > cumulativeSectorWeightings[keys[j]]
	})
	sortedKeys := map[string]int{}
	for i, k := range keys {
		sortedKeys[k] = i
	}

	portfolioSectorResult.Value = append([]fund.SectorName{fund.AnySector}, portfolioSectorResult.Value...)
	for _, weighting := range portfolioFundSectorWeightings {
		sort.Slice(weighting.FundSectorWeighting, func(i, j int) bool {
			iRank, jRank := sortedKeys[weighting.FundSectorWeighting[i].SectorName], sortedKeys[weighting.FundSectorWeighting[j].SectorName]
			return iRank < jRank
		})
	}
	return contracts.PortfolioDetails{
		FundInformation:               informationResult.Value.ConvertToResponse(),
		Sectors:                       portfolioSectorResult.Value.ConvertToResponse(),
		PortfolioFundSectorWeightings: portfolioFundSectorWeightings,
	}, nil
}
func (s *Service) GetPortfolios(ctx context.Context, userID uuid.UUID) ([]contracts.Portfolio, error) {
	p, err := s.repo.GetPortfolios(ctx, userID)
	if err != nil {
		return nil, err
	}
	return p.ConvertToResponse(), nil
}
func (s *Service) UpsertPortfolio(ctx context.Context,
	userID uuid.UUID,
	req contracts.Portfolio) (resp contracts.Portfolio, err error) {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return resp, err
	}
	defer s.repo.RollBack(tx, ctx)

	p := ConvertToModel(req)
	if p.ID != uuid.Nil {
		li, err := s.repo.GetListItems(ctx, p.ID)
		if err != nil {
			return resp, err
		}
		var itemsToDelete []uuid.UUID
		comparisonLoop := func(dbItem ListItem) bool {
			for _, newItem := range p.Items {
				if dbItem.ID == newItem.ID {
					return true
				}
			}
			return false
		}
		for _, dbItem := range li {
			match := comparisonLoop(dbItem)
			if !match {
				itemsToDelete = append(itemsToDelete, dbItem.ID)
			}
		}
		if len(itemsToDelete) > 0 {
			err = s.repo.DeleteListItems(ctx, itemsToDelete, tx)
			if err != nil {
				return resp, err
			}
		}
	}
	shouldSetCreatedAt := false
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
		shouldSetCreatedAt = true
	}
	portfolioModel := model.Portfolio{
		ID:     p.ID,
		UserID: &userID,
		Name:   &p.Name,
	}
	if shouldSetCreatedAt {
		now := time.Now()
		portfolioModel.CreatedAt = &now
	}
	err = s.repo.UpsertPortfolio(ctx, portfolioModel, tx)
	if err != nil {
		return resp, err
	}
	var pfs []model.PortfolioFund
	for i := range p.Items {
		if p.Items[i].FundID == uuid.Nil {
			continue
		}
		if p.Items[i].ID == uuid.Nil {
			p.Items[i].ID = uuid.New()
		}
		pf := model.PortfolioFund{
			ID:          p.Items[i].ID,
			PortfolioID: &p.ID,
			FundID:      &p.Items[i].FundID,
			Amount:      &p.Items[i].Amount,
		}
		pfs = append(pfs, pf)
	}
	err = s.repo.UpsertPortfolioListItems(ctx, pfs, tx)
	if err != nil {
		return resp, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return contracts.Portfolio{}, err
	}
	resp = p.ConvertToResponse()
	return resp, nil
}
