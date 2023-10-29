package handler

import (
	"net/http"

	"etfinsight/api/contracts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// func (h *Handler) Portfolios(c *gin.Context) {
//
//		fundDetails, err := h.portfolioService.GetPortfolios(c.Request.Context())
//		if err != nil {
//			_ = c.AbortWithError(http.StatusBadRequest, err)
//			return
//		}
//
//		c.JSON(http.StatusOK, fundDetails)
//	}
func (h *Handler) UpsertPortfolio(c *gin.Context) {
	var req contracts.Portfolio
	err := c.BindJSON(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fundDetails, err := h.portfolioService.UpsertPortfolio(c.Request.Context(), uuid.MustParse("b21b14c9-70bb-4336-a35c-7a69396ffbd8"), req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, fundDetails)
}
