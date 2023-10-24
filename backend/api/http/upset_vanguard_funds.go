package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpsertVanguardFunds(c *gin.Context) {
	//err := vs.UpsertFunds(c.Context())
	//if err != nil {
	//	return err
	//}
	c.JSON(http.StatusOK, nil)
}
