package handlers

import (
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/fund"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Body struct {
	SearchTerm string `json:"searchTerm"`
}

func SearchPage(c *fiber.Ctx) error {
	return c.Render("pages/search", fiber.Map{})
}

func SearchFunds(c *fiber.Ctx) error {
	b := Body{}
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	if b.SearchTerm == "" {
		return c.Render("partials/search/results", fiber.Map{
			"funds": nil,
		}, "")
	}
	repo := pgrepo.NewRepository(config.Load())
	hs := fund.NewService(repo)
	funds, err := hs.GetFunds(c.Context(), b.SearchTerm)
	if err != nil {
		return err
	}
	return c.Render("partials/search/results", fiber.Map{
		"funds": funds,
	}, "")
}

func FundDetails(c *fiber.Ctx) error {
	repo := pgrepo.NewRepository(config.Load())
	hs := fund.NewService(repo)

	fundIdString := c.Query("fundId")
	fundId, err := uuid.Parse(fundIdString)
	if err != nil {
		return err
	}

	fundDetails, err := hs.GetFundDetails(c.Context(), fundId, fund.FundHoldingsLimit)
	if err != nil {
		return err
	}
	return c.Render("pages/fundDetails", fiber.Map{
		"itemList":         fundDetails.Holdings,
		"sectorList":       fundDetails.Sectors,
		"sectorWeightings": fundDetails.SectorWeightings,
		"information":      fundDetails.Information,
		"limit":            fund.FundHoldingsLimit,
		"offset":           fund.FundHoldingsLimit,
	})
}
