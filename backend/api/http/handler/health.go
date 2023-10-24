package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz godoc
// @Summary Returns an HTTP 200 when healthy
// @Schemes
// @Tags Health
// @Accept json
// @Produce json
// @Success 200
// @Router /healthz [get]
func (h *Handler) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
