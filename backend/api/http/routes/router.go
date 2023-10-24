package routes

import (
	"etfinsight/api/http/handler"
	"etfinsight/api/http/middleware"
	"etfinsight/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRouter() Router {
	return Router{}
}

func (r Router) SetupRouter(engine *gin.Engine, config *config.Config, h *handler.Handler) error {
	engine.Use(
		middleware.HandleCORS(
			config.CORSAllowOrigin,
			config.CORSAllowCredentials,
			config.CORSAllowHeaders,
			config.CORSAllowMethods,
		),
	)
	engine.Use(middleware.HandleErrors())

	v1 := engine.Group("/api/v1")

	SetupFundRoutes(v1, h)

	return nil
}
