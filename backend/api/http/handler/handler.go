package handler

import (
	"balanceyourportfolio/api/contracts"
)

type Handler struct {
	fundService      contracts.FundService
	portfolioService contracts.PortfolioService
	userService      contracts.UserService
}

func NewHandler(
	fundService contracts.FundService,
	portfolioService contracts.PortfolioService,
	userService contracts.UserService,
) *Handler {
	return &Handler{
		fundService:      fundService,
		portfolioService: portfolioService,
		userService:      userService,
	}
}
