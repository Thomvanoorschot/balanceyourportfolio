package handlers

import (
	"context"

	"etfinsight/generated/proto"
	"etfinsight/utils/stringutils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PortfolioService interface {
	GetPortfolios(ctx context.Context, userId uuid.UUID) (*proto.PortfoliosResponse, error)
	UpsertPortfolio(ctx context.Context, userId uuid.UUID, req *proto.Portfolio) (*proto.UpsertPortfolioResponse, error)
	GetPortfolioDetails(ctx context.Context, userId uuid.UUID, portfolioId uuid.UUID) (*proto.PortfolioDetailsResponse, error)
	FilterPortfolioHoldings(ctx context.Context, req *proto.FilterPortfolioFundHoldingsRequest) (*proto.FilterPortfolioFundHoldingsResponse, error)
	UpdatePortfolioFundAmount(ctx context.Context, req *proto.UpdatePortfolioFundAmountRequest) (*proto.Empty, error)
}

type PortfolioHandler struct {
	proto.UnimplementedPortfolioServiceServer
	portfolioService PortfolioService
}

func NewPortfolioHandler(portfolioService PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioService: portfolioService,
	}
}

func (h *PortfolioHandler) GetPortfolios(ctx context.Context, _ *proto.PortfoliosRequest) (*proto.PortfoliosResponse, error) {
	resp, err := h.portfolioService.GetPortfolios(ctx, stringutils.ConvertToUUID("b21b14c9-70bb-4336-a35c-7a69396ffbd8"))
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
func (h *PortfolioHandler) UpsertPortfolio(ctx context.Context, req *proto.UpsertPortfolioRequest) (*proto.UpsertPortfolioResponse, error) {
	resp, err := h.portfolioService.UpsertPortfolio(ctx, stringutils.ConvertToUUID("b21b14c9-70bb-4336-a35c-7a69396ffbd8"), req.Portfolio)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
func (h *PortfolioHandler) UpdatePortfolioFundAmount(ctx context.Context, req *proto.UpdatePortfolioFundAmountRequest) (*proto.Empty, error) {
	resp, err := h.portfolioService.UpdatePortfolioFundAmount(ctx, req)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
func (h *PortfolioHandler) GetPortfolioDetails(ctx context.Context, req *proto.PortfolioDetailsRequest) (*proto.PortfolioDetailsResponse, error) {
	resp, err := h.portfolioService.GetPortfolioDetails(ctx,
		stringutils.ConvertToUUID("b21b14c9-70bb-4336-a35c-7a69396ffbd8"),
		stringutils.ConvertToUUID(req.PortfolioId),
	)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}

func (h *PortfolioHandler) FilterPortfolioHoldings(ctx context.Context, req *proto.FilterPortfolioFundHoldingsRequest) (*proto.FilterPortfolioFundHoldingsResponse, error) {
	resp, err := h.portfolioService.FilterPortfolioHoldings(ctx, req)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
