package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
	"testing"

	"gin-test/handlers"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/producttype/health", nil)
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		healthHandler := handlers.NewHealthHandler()
		healthHandler.HealthCheck(c)

		expectedBody := `{"message":"ProductType Service : OK"}`
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedBody, rec.Body.String())
	})
}