package routes_test

import (
	"net/http"
	"testing"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gin-test/routes"
)

func TestSetupProductTypeRoutes(t *testing.T) {
	// testCases := []struct {
    //     name       string
    //     path       string
    //     method     string
    // }{
    //     {name: "test case : Create pass",  path: "/", 		method: http.MethodGet},
    //     {name: "test case : GetAll pass",  path: "/:id", 	method: http.MethodGet},
    //     {name: "test case : GetById pass", path: "/", 		method: http.MethodPost},
    //     {name: "test case : Update pass",  path: "/:id", 	method: http.MethodPut},
    //     {name: "test case : Delete pass",  path: "/:id", 	method: http.MethodDelete},
    //     {name: "test case : Count pass",   path: "/count", 	method: http.MethodGet},
    // }
	
    // for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {

	// 		r := gin.New()
	// 		productTypeGroup := r.Group("/producttype")
			
	// 		routes.SetupProductTypeRoutes(productTypeGroup, mockDb.GetDB())

	// 		req, err := http.NewRequest(tc.method, "/producttype" + tc.path, nil)
	// 		assert.NoError(t, err)
			
	// 		rec := httptest.NewRecorder()
	// 		r.ServeHTTP(rec, req)

	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	})
	// }
}


func TestSetupHealth(t *testing.T) {
	t.Run("test route : health pass", func(t *testing.T) {
		r := gin.New()
		productTypeGroup := r.Group("/producttype")

		routes.SetupHealth(productTypeGroup)
		
		req, err := http.NewRequest(http.MethodGet, "/producttype/health", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}