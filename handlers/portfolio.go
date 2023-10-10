package handlers

import (
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/portfolio"

	"github.com/gofiber/fiber/v2"
)

func PortfoliosPage(c *fiber.Ctx) error {
	return c.Render("portfolios", fiber.Map{})
}

func CreatePortfolio(c *fiber.Ctx) error {
	req := portfolio.CreatePortfolioRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	repo := pgrepo.NewRepository(config.Load())
	hs := portfolio.NewService(repo)
	err := hs.CreatePortfolio(c.Context(), req)
	if err != nil {
		return err
	}
	return c.Render("partials/portfolioList", fiber.Map{
		"itemList": nil,
	}, "")
}
