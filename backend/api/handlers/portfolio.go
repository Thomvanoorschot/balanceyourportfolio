package handlers

import (
	"context"

	"balanceyourportfolio/generated/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PortfolioService interface {
	GetPortfolios(ctx context.Context, req *proto.PortfoliosRequest) (*proto.PortfoliosResponse, error)
	UpsertPortfolio(ctx context.Context, req *proto.UpsertPortfolioRequest) (*proto.UpsertPortfolioResponse, error)
	GetPortfolioDetails(ctx context.Context, req *proto.PortfolioDetailsRequest) (*proto.PortfolioDetailsResponse, error)
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

func (h *PortfolioHandler) GetPortfolios(ctx context.Context, req *proto.PortfoliosRequest) (*proto.PortfoliosResponse, error) {
	resp, err := h.portfolioService.GetPortfolios(ctx, req)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
func (h *PortfolioHandler) UpsertPortfolio(ctx context.Context, req *proto.UpsertPortfolioRequest) (*proto.UpsertPortfolioResponse, error) {
	resp, err := h.portfolioService.UpsertPortfolio(ctx, req)
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
	resp, err := h.portfolioService.GetPortfolioDetails(ctx, req)
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
