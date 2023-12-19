package handlers

import (
	"context"

	"etfinsight/generated/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FundService interface {
	GetDetails(ctx context.Context, fundID uuid.UUID) (*proto.FundDetailsResponse, error)
	GetFundsWithTickers(ctx context.Context, searchTerm string) (*proto.SearchFundsResponse, error)
	FilterHoldings(ctx context.Context, filter *proto.FilterHoldingsRequest) (*proto.HoldingsListResponse, error)
}

type FundHandler struct {
	proto.UnimplementedFundServiceServer
	fundService FundService
}

func NewFundHandler(fundService FundService) *FundHandler {
	return &FundHandler{
		fundService: fundService,
	}
}

func (h *FundHandler) GetDetails(ctx context.Context, req *proto.GetFundDetailsRequest) (*proto.FundDetailsResponse, error) {
	fundID, err := uuid.Parse(req.FundId)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument, "could not parse fundId",
		)
	}
	resp, err := h.fundService.GetDetails(ctx, fundID)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}
func (h *FundHandler) SearchFunds(ctx context.Context, req *proto.SearchFundsRequest) (*proto.SearchFundsResponse, error) {
	resp, err := h.fundService.GetFundsWithTickers(ctx, req.SearchTerm)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}
	return resp, nil
}
func (h *FundHandler) FilterHoldings(ctx context.Context, req *proto.FilterHoldingsRequest) (*proto.HoldingsListResponse, error) {
	resp, err := h.fundService.FilterHoldings(ctx, req)
	if err != nil {
		return nil, status.Error(
			codes.Unknown, err.Error(),
		)
	}

	return resp, nil
}