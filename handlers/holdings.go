package handlers

import (
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/fund"

	"github.com/gofiber/fiber/v2"
)

func FilterHoldings(c *fiber.Ctx) error {
	hf := fund.HoldingsFilter{}
	if err := c.BodyParser(&hf); err != nil {
		return err
	}
	if hf.Limit == 0 {
		hf.Limit = fund.FundHoldingsLimit
	}
	if hf.SectorName == fund.AnySector {
		hf.SectorName = ""
	}
	repo := pgrepo.NewRepository(config.Load())
	hs := fund.NewService(repo)
	fundHoldings, err := hs.FilterHoldings(c.Context(), hf)
	if err != nil {
		return err
	}
	return c.Render("partials/holdingList", fiber.Map{
		"itemList": fundHoldings,
		"limit":    hf.Limit,
		"offset":   hf.Offset + hf.Limit,
	}, "")
}
