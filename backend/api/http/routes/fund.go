package routes

import (
	"etfinsight/api/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupFundRoutes(routes *gin.RouterGroup, handler *handler.Handler) {
	g := routes.Group("/fund")

	g.GET("/:fundID/details", handler.FundDetails)
	g.GET("/search", handler.SearchFunds)
	g.POST("/holdings/filter", handler.FilterHoldings)
}
