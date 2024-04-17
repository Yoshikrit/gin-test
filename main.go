package main

import (
	"github.com/gin-gonic/gin"
	"os"

    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
	_ "gin-test/docs"
	
	"gin-test/configs"
	"gin-test/models"
	"gin-test/routes"
	"gin-test/middlewares"
)

// @title Gin-Test - ProductType API
// @description Gin-Test - Teletubbie's ProductType API.
// @version 1.0

// @contact.name   Walter White
// @contact.url    https://twitter.com/example
// @contact.email  example@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /

// @schemes http https
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	configs.LoadEnv()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.SetTrustedProxies(nil)

	// Middleware
	m := middlewares.Metrics()
	m.Use(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.Cors())

	configs.DatabaseInit()
	defer configs.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := configs.GetDB()
	err := db.AutoMigrate(&models.ProductTypeEntity{})
	if err != nil {
		panic(err)
	}

	// Routes
	productTypeGroup := r.Group("/producttype")
	routes.SetupProductTypeRoutes(productTypeGroup, db)
	routes.SetupHealth(productTypeGroup)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
