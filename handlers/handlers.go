package handlers

import (
	"gin-test/utils/errs"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/go-playground/validator"
)

func HandleError(ctx *gin.Context, err error) {
	if e, ok := err.(errs.AppError); ok {
		ctx.JSON(e.Code, gin.H{
			"code":    e.Code,
			"message": e.Message,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": "Internal Server Error",
	})
}

func GetIntId(ctx *gin.Context) (int, error) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errs.NewBadRequestError("Invalid id: " + idParam + " is not integer")
	}
	return id, nil
}

func NewProductTypeValidator() *ProductTypeValidator {
	return &ProductTypeValidator{validator: validator.New()}
}

type ProductTypeValidator struct {
	validator *validator.Validate
}

func (p *ProductTypeValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}