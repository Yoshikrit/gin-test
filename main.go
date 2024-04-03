package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"os"
	
	"gin-test/configs"
	"gin-test/models"
	"gin-test/routes"
)

func main() {
	configs.LoadEnv()

	if (os.Getenv("APP_ENV") == "production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	// middleware
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
	routes.SetupHealth(productTypeGroup)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
