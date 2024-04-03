package handlers

import (
	"gin-test/models"
	"gin-test/services"
	"gin-test/utils/logs"
	"gin-test/utils/errs"

	"github.com/gin-gonic/gin"
    "net/http"
)

type productTypeHandler struct {
	productTypeSrv services.ProductTypeService
}

func NewProductTypeHandler(productTypeSrv services.ProductTypeService) productTypeHandler {
	return productTypeHandler{productTypeSrv: productTypeSrv}
}

func (h *productTypeHandler) CreateProductType(ctx *gin.Context){
	var prodTypeReq models.ProductTypeCreate
	if err := ctx.ShouldBindJSON(&prodTypeReq); err != nil {
		logs.Error(err.Error())
		HandleError(ctx, errs.NewBadRequestError(err.Error()))
		return 
	}

	prodTypeRes, err := h.productTypeSrv.CreateProductType(&prodTypeReq)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Create ProductType Successfully")
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": prodTypeRes,
	})
}

func (h *productTypeHandler) GetAllProductTypes(ctx *gin.Context){
	prodTypesRes, err := h.productTypeSrv.GetProductTypes()
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Get ProductTypes Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": prodTypesRes,
	})
}

func (h *productTypeHandler) GetProductTypeByID(ctx *gin.Context){
	id, err := GetIntId(ctx)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	prodTypeRes, err := h.productTypeSrv.GetProductType(id)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Get ProductType Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": prodTypeRes,
	})
}

func (h *productTypeHandler) UpdateProductTypeByID(ctx *gin.Context){
	id, err := GetIntId(ctx)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return
	}

	var prodTypeReq models.ProductTypeUpdate
	if err := ctx.ShouldBindJSON(&prodTypeReq); err != nil {
		logs.Error(err.Error())
		HandleError(ctx, errs.NewBadRequestError(err.Error()))
		return
	}

	prodTypeRes, err := h.productTypeSrv.UpdateProductType(id, &prodTypeReq)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Update ProductType Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": prodTypeRes,
	})
}

func (h *productTypeHandler) DeleteProductTypeByID(ctx *gin.Context){
	id, err := GetIntId(ctx)
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return
	}

	if err := h.productTypeSrv.DeleteProductType(id); err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Delete ProductType Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Delete ProductType Successfully",
	})
}

func (h *productTypeHandler) GetProductTypeCount(ctx *gin.Context) {
    count, err := h.productTypeSrv.GetProductTypeCount()
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Get ProductType's Count Successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": count,
	})
}
