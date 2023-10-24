package http

import (
	"etfinsight/api/http/handler"
	"etfinsight/api/http/routes"
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/fund"

	"github.com/gin-gonic/gin"
)

func ListenAndServe(addr string) error {
	e := gin.Default()
	cfg := config.Load()

	repo := pgrepo.NewRepository(cfg)
	h := handler.NewHandler(fund.NewService(repo), nil, nil)

	router := routes.NewRouter()
	err := router.SetupRouter(e, cfg, h)
	if err != nil {
		return err
	}

	return e.Run(addr)
}
