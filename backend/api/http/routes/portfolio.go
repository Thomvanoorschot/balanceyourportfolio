package routes

import (
	"etfinsight/api/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupPortfolioRoutes(routes *gin.RouterGroup, handler *handler.Handler) {
	g := routes.Group("/portfolio")

	g.PUT("/", handler.UpsertPortfolio)
}
