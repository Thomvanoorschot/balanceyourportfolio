package contracts

import (
	"context"

	"github.com/google/uuid"
)

type FundService interface {
	GetDetails(ctx context.Context, fundID uuid.UUID) (FundDetails, error)
	GetFundsWithTickers(ctx context.Context, searchTerm string) ([]Fund, error)
	FilterHoldings(ctx context.Context, filter FundHoldingsFilter) ([]FundHolding, error)
}

type PortfolioService interface {
}

type UserService interface {
}

type VanguardService interface {
}
