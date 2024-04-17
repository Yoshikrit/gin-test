package handlers

import (
	"gin-test/utils/errs"

	"github.com/gin-gonic/gin"
	// "github.com/golodash/galidator"
	"net/http"
	"strconv"
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

func GetIntID(ctx *gin.Context) (int, error) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errs.NewBadRequestError("Invalid id: " + idParam + " is not integer")
	}
	return id, nil
}

// var (
// 	g = galidator.New()
// )

// func handleValidationError(c *gin.Context, err error) {
// 	var errors []string
// 	for _, v := range err.(validator.ValidationErrors) {
// 		// Create your custom error message string
// 		errorMessage := fmt.Sprintf("Field %s %s", v.Field(), v.ActualTag())
// 		errors = append(errors, errorMessage)
// 	}
// 	// Respond with custom error messages
// 	c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
// }
