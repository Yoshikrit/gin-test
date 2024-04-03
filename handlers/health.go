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

// @BasePath /producttype

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description test health
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /health [get]
func (h *healthHandler) CheckHealth(ctx *gin.Context){
	logs.Info("Handler: Create ProductType Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Service ProductType : OK",
	})
}
