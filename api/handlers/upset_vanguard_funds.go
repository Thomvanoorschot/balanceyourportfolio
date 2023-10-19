package handlers

import (
	"etfinsight/clients"
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/vanguard"

	"github.com/gofiber/fiber/v2"
)

func UpsertVanguardFunds(c *fiber.Ctx) error {
	repo := pgrepo.NewRepository(config.Load())
	vc := clients.NewVanguard(config.Load())
	vs := vanguard.NewService(vc, repo)

	return vs.UpsertFunds(c.Context())
}
