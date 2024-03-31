package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	// "errors"

	"gin-test/handlers"
	// "gin-test/models"
	"gin-test/utils/errs"
)

func SetUpRouter() *gin.Engine{
    router := gin.New()
    return router
}

func TestHandleError(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
        router := SetUpRouter()
		
		appErr := errs.AppError{Code: http.StatusBadRequest, Message: "Bad Request"}
		router.GET("/", func(c *gin.Context) {
			handlers.HandleError(c, appErr)
		})
	
		req := httptest.NewRequest(http.MethodGet, "/", nil)
	
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	
		assert.Equal(t, http.StatusBadRequest, w.Code)
	
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	
		if err != nil {
			t.Errorf("Error unmarshalling response body: %v", err)
		}
	
		expectedBody := map[string]interface{}{"code": 400, "message": "Bad Request"}
		assert.Equal(t, expectedBody, responseBody)
	})

	// t.Run("test case : fail", func(t *testing.T) {
	// 	ctx := gin.New()

	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	rec := httptest.NewRecorder()

	// 	appErr := errors.New("")

	// 	handlers.HandleError(ctx.NewContext(req, rec), appErr)

	// 	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	// 	expectedBody := `{"code":500,"message":"Interval Server Error"}`
	// 	actualBody := strings.TrimSpace(rec.Body.String())

	// 	assert.Equal(t, expectedBody, actualBody)
	// })
}

// func TestGetIntId(t *testing.T) {
	// ctx := gin.New()

	// t.Run("test case : pass valid integer id", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := ctx.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("123")

	// 	id, err := handlers.GetIntId(c)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, 123, id)
	// })

	// t.Run("test case : invalid non-integer id", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := ctx.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("a")

	// 	id, err := handlers.GetIntId(c)

	// 	expectedErr := errs.NewBadRequestError("Invalid id: a is not integer")

	// 	assert.Error(t, err)
    //     assert.Equal(t, expectedErr, err)
    //     assert.Equal(t, 0, id)
	// })
// }

// func TestProductTypeValidator_Validate(t *testing.T) {
	// validator := handlers.NewProductTypeValidator()

	// t.Run("valid input", func(t *testing.T) {
	// 	productTypeCreate := models.ProductTypeCreate{
	// 		Id: 1,
	// 		Name: "burger",
	// 	}

	// 	err := validator.Validate(productTypeCreate)

	// 	assert.NoError(t, err)
	// })

	// t.Run("invalid input", func(t *testing.T) {
	// 	invalidProductTypeCreate := models.ProductTypeCreate{
	// 		Id: 1,
	// 		Name: "",
	// 	}

	// 	err := validator.Validate(invalidProductTypeCreate)

	// 	assert.Error(t, err)
	// })
// }