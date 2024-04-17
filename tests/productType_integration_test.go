package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"

	"gin-test/handlers"
	"gin-test/models"
	"gin-test/services"
	"gin-test/tests/mocks/mock_repositories"
)

func TestCreateProductType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	prodTypeReqMock := &models.ProductTypeCreate{
		Id:   1,
		Name: "A",
	}

	prodTypeReqErrorMock := &models.ProductTypeCreate{
		Id:   0,
		Name: "",
	}

	prodTypeFromDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "A",
	}

	prodTypeResMock := &models.ProductType{
		Id:   1,
		Name: "A",
	}

	prodTypeReqJSON, _ := json.Marshal(prodTypeReqMock)
	prodTypeReqErrorJSON, _ := json.Marshal(prodTypeReqErrorMock)
	prodTypeResJSON, _ := json.Marshal(prodTypeResMock)

	type testCase struct {
		name           string
		body           string
		expectedStatus int
		expectedBody   string
	}

	cases := []testCase{
		{name: "test case : pass",        body: string(prodTypeReqJSON),      expectedStatus: 201, expectedBody: string(prodTypeResJSON)},
		{name: "test case : failed bind", body: string(prodTypeReqErrorJSON), expectedStatus: 400, expectedBody: `{"code":400,"message":"Key: 'ProductTypeCreate.Id' Error:Field validation for 'Id' failed on the 'required' tag\nKey: 'ProductTypeCreate.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			mockRepo.On("Create", prodTypeReqMock).Return(prodTypeFromDBMock, nil)

			prodTypeService := services.NewProductTypeService(mockRepo)
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.CreateProductType(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestGetAllProductTypes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	prodTypesDBMock := []models.ProductTypeEntity{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	prodTypesResMock := []models.ProductType {
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	prodTypesResJSON, _ := json.Marshal(&prodTypesResMock)

	type testCase struct {
		name         	string
		body            string
		expectedStatus  int
		expectedBody    string
	}

	cases := []testCase{
		{name: "test case : pass", expectedStatus: 200,	expectedBody: string(prodTypesResJSON),},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req

			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			mockRepo.On("GetAll").Return(prodTypesDBMock, nil)

			prodTypeService := services.NewProductTypeService(mockRepo)
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.GetProductTypes(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestGetProductTypeByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	prodTypeDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "A",
	}

	prodTypeResMock := models.ProductType {
		Id:   1,
		Name: "A",
	}

	prodTypeResJSON, _ := json.Marshal(&prodTypeResMock)

	type testCase struct {
		name         	string
		param           string
		expectedStatus  int
		expectedBody    string
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", expectedStatus: 200,	expectedBody: string(prodTypeResJSON),},
		{name: "test case : failed param int",  param: "a", expectedStatus: 400,	expectedBody: `{"code":400,"message":"Invalid id: a is not integer"}`,},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/producttype/", nil)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Params = []gin.Param{{
				Key:   "id",
				Value: tc.param,
			}}
			c.Request = req

			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			mockRepo.On("GetById", 1).Return(prodTypeDBMock, nil)

			prodTypeService := services.NewProductTypeService(mockRepo)
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.GetProductTypeByID(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestUpdateProductTypeByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	prodTypeReqMock := &models.ProductTypeUpdate {
		Name: "B",
	}

	prodTypeErrorReqMock := &models.ProductTypeUpdate {
		Name: "",
	}

	prodTypeServResMock := &models.ProductType {
		Id:   1,
		Name: "B",
	}

	prodTypeDBMock := &models.ProductTypeEntity{
		Id:   1,
		Name: "B",
	}

	prodTypeReqJSON, _ := json.Marshal(prodTypeReqMock)
	prodTypeErrorReqJSON, _ := json.Marshal(prodTypeErrorReqMock)
	prodTypeResJSON, _ := json.Marshal(prodTypeServResMock)

	type testCase struct {
		name         	string
		param           string
		body            string
		expectedStatus  int
		expectedBody    string
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", body: string(prodTypeReqJSON), 	  	expectedStatus: 200,	expectedBody: string(prodTypeResJSON),},
		{name: "test case : failed param int",  param: "a", body: string(prodTypeReqJSON),		expectedStatus: 400,	expectedBody: `{"code":400,"message":"Invalid id: a is not integer"}`,},
		{name: "test case : failed bind",  		param: "1", body: string(prodTypeErrorReqJSON), expectedStatus: 400,	expectedBody: `{"code":400,"message":"Key: 'ProductTypeUpdate.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Params = []gin.Param{{
				Key:   "id",
				Value: tc.param,
			}}
			c.Request = req

			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			mockRepo.On("Update", 1, prodTypeReqMock).Return(prodTypeDBMock, nil)

			prodTypeService := services.NewProductTypeService(mockRepo)
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.UpdateProductTypeByID(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestDeleteProductTypeByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	prodTypeDeleteResMock := &gin.H{
		"message": "Delete ProductType Successfully",
	}

	prodTypeResJSON, _ := json.Marshal(prodTypeDeleteResMock)

	type testCase struct {
		name         	string
		param           string
		expectedStatus  int
		expectedBody    string
		dbReturn	  	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", expectedStatus: 200,	expectedBody: string(prodTypeResJSON),									dbReturn: nil},
		{name: "test case : failed param int",  param: "a", expectedStatus: 400,	expectedBody: `{"code":400,"message":"Invalid id: a is not integer"}`,  dbReturn: errors.New("")},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/prodTypeuct/", nil)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Params = []gin.Param{{
				Key:   "id",
				Value: tc.param,
			}}
			c.Request = req

			mockRepo := mock_repositories.NewProductTypeRepositoryMock()
			mockRepo.On("DeleteById", 1).Return(nil)

         	prodTypeService := services.NewProductTypeService(mockRepo)
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.DeleteProductTypeByID(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
		})
	}
}

func TestGetProductTypeCount(t *testing.T) {
	gin.SetMode(gin.TestMode)

	num := int64(42)

	prodTypeServResMock := &num

	prodTypeResJSON, _ := json.Marshal(prodTypeServResMock)

    t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockRepo := mock_repositories.NewProductTypeRepositoryMock()
		mockRepo.On("GetCount").Return(num, nil)
		
		prodTypeService := services.NewProductTypeService(mockRepo)
		prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
		prodTypeHandler.GetProductTypeCount(c)

		expectedCode := 200
		expectedBody := string(prodTypeResJSON)
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
	})
}
