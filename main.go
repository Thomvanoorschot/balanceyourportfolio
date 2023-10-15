package main

import (
	"fmt"
	"time"

	"etfinsight/handlers"
	"etfinsight/services/fund"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		PassLocalsToViews:     true,
		Views:                 createEngine(),
		ViewsLayout:           "layouts/main",
	})

	app.Get("/upsert-vanguard-funds", handlers.UpsertVanguardFunds)

	search := app.Group("/search")
	search.Get("/", handlers.SearchPage)
	search.Post("/", handlers.SearchFunds)

	fundDetails := app.Group("/fund-details")
	fundDetails.Get("/", handlers.FundDetails)
	fundDetails.Post("/filter", handlers.FilterHoldings)

	portfolios := app.Group("/portfolios")
	portfolios.Get("/", handlers.PortfoliosPage)

	app.Static(
		"/static", // mount address
		"./www",
	)
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createEngine() *handlebars.Engine {
	engine := handlebars.New("./www", ".html")
	engine.AddFunc("equals", func(x, y string) bool {
		return x == y
	})
	engine.AddFunc("formatFloat", func(f float64) string {
		return fmt.Sprintf("%.2f%%", f)
	})
	engine.AddFunc("formatDate", func(t time.Time) string {
		return t.Format("2006-02-01")
	})
	engine.AddFunc("formatRelativePercentage", func(percentage float64, sectorWeightings []fund.SectorWeighting) string {
		return fmt.Sprintf("%.2f%%", percentage/sectorWeightings[0].Percentage*100)
	})
	engine.Reload(true)
	return engine
}
