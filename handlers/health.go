package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gin-test/utils/logs"
)

type healthHandler struct {
}

func NewHealthHandler() healthHandler {
	return healthHandler{}
}

// HealthCheckHandler godoc
// @Summary Health Check
// @Description Health check
// @id HealthCheckHandler
// @Tags producttype
// @Accept  json
// @Produce  json
// @response 200 {object} models.Response "ProductType service is running"
// @Router /producttype/health [get]
func (h *healthHandler) HealthCheck(ctx *gin.Context){
	logs.Info("Handler: ProductType service is  running")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ProductType Service : OK",
	})
}
