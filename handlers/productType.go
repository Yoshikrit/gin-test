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

// CreateProductTypeHandler godoc
// @Summary Create ProductType
// @Description Create producttype
// @Tags producttype
// @Accept  json
// @Produce  json
// @param ProductType body models.ProductTypeCreate true "ProductType data to be create"
// @response 200 {object} models.ProductType{} "Create ProductType Successfully"
// @response 400 {object} errs.AppError "Error Bad Request"
// @response 409 {object} errs.AppError "Error Conflict Error"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/ [post]
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
	ctx.JSON(http.StatusCreated, prodTypeRes)
}

// GetAllProductTypesHandler godoc
// @Summary Get All ProductType
// @Description Get all producttype
// @Tags producttype
// @Accept  json
// @Produce  json
// @response 200 {array} models.ProductType{} "Get ProductTypes Successfully"
// @response 404 {object} errs.AppError "Error Not Found"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/ [get]
func (h *productTypeHandler) GetProductTypes(ctx *gin.Context){
	prodTypesRes, err := h.productTypeSrv.GetProductTypes()
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Get ProductTypes Successfully")
	ctx.JSON(http.StatusOK, prodTypesRes)
}

// GetProductTypeByIDHandler godoc
// @Summary Get ProductType
// @Description Get producttype by id
// @Tags producttype
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "ProductType ID"
// @response 200 {object} models.ProductType{} "Get ProductType Successfully"
// @response 400 {object} errs.AppError "Error Bad Request"
// @response 404 {object} errs.AppError "Error Not Found"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/{id} [get]
func (h *productTypeHandler) GetProductTypeByID(ctx *gin.Context){
	id, err := GetIntID(ctx)
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
	ctx.JSON(http.StatusOK, prodTypeRes)
}

// UpdateProductTypeByIDHandler godoc
// @Summary Update ProductType
// @Description Update producttype by id
// @Tags producttype
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "ProductType ID"
// @param ProductType body models.ProductTypeUpdate true "ProductType data to be update"
// @response 200 {object} models.ProductType{} "Update ProductType Successfully"
// @response 400 {object} errs.AppError "Error Bad Request"
// @response 404 {object} errs.AppError "Error Not Found"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/{id} [put]
func (h *productTypeHandler) UpdateProductTypeByID(ctx *gin.Context){
	id, err := GetIntID(ctx)
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
	ctx.JSON(http.StatusOK, prodTypeRes)
}

// DeleteProductTypeByIDHandler godoc
// @Summary Delete ProductType
// @Description Delete producttype by id
// @Tags producttype
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "ProductType ID"
// @response 200 {object} models.Response "Delete ProductType Successfully"
// @response 400 {object} errs.AppError "Error Bad Request"
// @response 404 {object} errs.AppError "Error Not Found"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/{id} [delete]
func (h *productTypeHandler) DeleteProductTypeByID(ctx *gin.Context){
	id, err := GetIntID(ctx)
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
		"message": "Delete ProductType Successfully",
	})
}

// GetProductTypeCountHandler godoc
// @Summary Get ProductType Count
// @Description Get producttype's count from database
// @Tags producttype
// @Accept  json
// @Produce  json
// @response 200 {integer} int "Get ProductType's Count Successfully"
// @response 500 {object} errs.AppError "Error Unexpected Error"
// @Router /producttype/count [get]
func (h *productTypeHandler) GetProductTypeCount(ctx *gin.Context) {
    count, err := h.productTypeSrv.GetProductTypeCount()
	if err != nil {
		logs.Error(err.Error())
		HandleError(ctx, err)
		return 
	}

	logs.Info("Handler: Get ProductType's Count Successfully")
	ctx.JSON(http.StatusOK, count)
}
