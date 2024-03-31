package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
	"os"

	"gin-test/configs"
	"gin-test/models"
	"gin-test/routes"
)

func main() {
	configs.LoadEnv()

	r := gin.New()

	//middleware
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	configs.DatabaseInit()
	defer configs.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := configs.GetDB()
	err := db.AutoMigrate(&models.ProductTypeEntity{})
	if err != nil {
		panic(err)
	}

	//routes
	productTypeGroup := r.Group("/producttype")
	routes.SetupProductTypeRoutes(productTypeGroup)

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "Service : OK",
		})
	})

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
