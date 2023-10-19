package handlers

import (
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/fund"
	"etfinsight/services/portfolio"

	"github.com/gofiber/fiber/v2"
)

func PortfoliosPage(c *fiber.Ctx) error {
	return c.Render("pages/portfolios", fiber.Map{})
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
	return c.Render("partials/portfolios/list", fiber.Map{
		"itemList": nil,
	}, "")
}

func SearchPortfolioFunds(c *fiber.Ctx) error {
	b := Body{}
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	if b.SearchTerm == "" {
		return c.Render("partials/portfolios/filterResults", fiber.Map{
			"funds": nil,
		}, "")
	}
	repo := pgrepo.NewRepository(config.Load())
	hs := fund.NewService(repo)
	funds, err := hs.GetFundsWithTickers(c.Context(), b.SearchTerm)
	if err != nil {
		return err
	}
	return c.Render("partials/portfolios/filterResults", fiber.Map{
		"funds": funds,
	}, "")
}
