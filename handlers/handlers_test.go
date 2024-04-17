package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	// "encoding/json"
	"errors"
	"strings"

	"gin-test/handlers"
	// "gin-test/models"
	"gin-test/utils/errs"
)

func TestHandleError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		appErr := errs.AppError{Code: http.StatusBadRequest, Message: "Bad Request"}
		handlers.HandleError(c, appErr)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		expectedBody := `{"code":400,"message":"Bad Request"}`
		actualBody := strings.TrimSpace(rec.Body.String())

		assert.Equal(t, expectedBody, actualBody)
	})

	t.Run("test case : fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		appErr := errors.New("")

		handlers.HandleError(c, appErr)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expectedBody := `{"code":500,"message":"Internal Server Error"}`
		actualBody := strings.TrimSpace(rec.Body.String())

		assert.Equal(t, expectedBody, actualBody)
	})
}

func TestGetIntId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("test case : pass valid integer id", func(t *testing.T) {
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Params = []gin.Param{{
			Key:   "id",
			Value: "123",
		}}

		id, err := handlers.GetIntID(c)

		assert.NoError(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("test case : invalid non-integer id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Params = []gin.Param{{
			Key:   "id",
			Value: "a",
		}}

		id, err := handlers.GetIntID(c)

		expectedErr := errs.NewBadRequestError("Invalid id: a is not integer")

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, 0, id)
	})
}

