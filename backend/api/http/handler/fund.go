package handler

import (
	"net/http"

	"etfinsight/api/contracts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) FundDetails(c *gin.Context) {
	fundID, err := uuid.Parse(c.Param("fundID"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fundDetails, err := h.fundService.GetDetails(c.Request.Context(), fundID)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, fundDetails)
}
func (h *Handler) SearchFunds(c *gin.Context) {
	searchTerm := c.Query("searchTerm")
	funds, err := h.fundService.GetFundsWithTickers(c.Request.Context(), searchTerm)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, funds)
}
func (h *Handler) FilterHoldings(c *gin.Context) {
	var filter contracts.FundHoldingsFilter
	err := c.BindJSON(&filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	funds, err := h.fundService.FilterHoldings(c.Request.Context(), filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, funds)
}
