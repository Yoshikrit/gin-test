package routes

import (
	"gin-test/repositories"
	"gin-test/services"
	"gin-test/handlers"
	
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func SetupProductTypeRoutes(g *gin.RouterGroup, db *gorm.DB) {
	productTypeRepository := repositories.NewProductTypeRepositoryDB(db)
	productTypeService := services.NewProductTypeService(productTypeRepository)
	productTypeHandler := handlers.NewProductTypeHandler(productTypeService)

	g.GET("/", productTypeHandler.GetAllProductTypes)
	g.GET("/:id", productTypeHandler.GetProductTypeByID)
	g.POST("/", productTypeHandler.CreateProductType)
	g.PUT("/:id", productTypeHandler.UpdateProductTypeByID)
	g.DELETE("/:id", productTypeHandler.DeleteProductTypeByID)
	g.GET("/count", productTypeHandler.GetProductTypeCount)
}

func SetupHealth(g *gin.RouterGroup) {
	healthHandler := handlers.NewHealthHandler()
	
	g.GET("/health", healthHandler.CheckHealth)
}