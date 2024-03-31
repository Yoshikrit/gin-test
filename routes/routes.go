package routes

import (
	"gin-test/repositories"
	"gin-test/services"
	"gin-test/handlers"
	"gin-test/configs"
	
	"github.com/gin-gonic/gin"
)

func SetupProductTypeRoutes(g *gin.RouterGroup) {
	productTypeRepository := repositories.NewProductTypeRepositoryDB(configs.GetDB())
	productTypeService := services.NewProductTypeService(productTypeRepository)
	productTypeHandler := handlers.NewProductTypeHandler(productTypeService)

	g.GET("/", productTypeHandler.GetAllProductTypes)
	g.GET("/:id", productTypeHandler.GetProductTypeByID)
	g.POST("/", productTypeHandler.CreateProductType)
	g.PUT("/:id", productTypeHandler.UpdateProductTypeByID)
	g.DELETE("/:id", productTypeHandler.DeleteProductTypeByID)
	g.GET("/count", productTypeHandler.GetProductTypeCount)
}