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

// @Summary Check health of the service
// @Description Check if the service is up and running
// @Tags health
// @Accept  json
// @Produce  json
// @response 200 {string} string "ProductType Service : OK"
// @Router /producttype/health [get]
func (h *healthHandler) CheckHealth(ctx *gin.Context){
	logs.Info("Handler: Create ProductType Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ProductType Service : OK",
	})
}
