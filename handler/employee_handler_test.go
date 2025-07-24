package handler

import (
	"employee-service-gin/model"
	"employee-service-gin/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.POST("/employees/", AddEmployee)
	r.GET("/employees/", GetEmployees)
	r.DELETE("/employees/:id", DeleteEmployee)

	return r
}

func TestAddEmployeeHandler(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest("POST", "/employees/",
		strings.NewReader(`{"id":99,"name":"Test","location":"Nowhere"}`))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.Code)
	}
}

func TestGetEmployeesHandler(t *testing.T) {
	repository.AddEmployeeRepo(&model.Employee{ID: 1, Name: "Greg", Location: "SG"})

	r := setupTestRouter()

	req := httptest.NewRequest("GET", "/employees/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.Code)
	}
}

func TestDeleteEmployeeHandler(t *testing.T) {
	repository.AddEmployeeRepo(&model.Employee{ID: 123, Name: "Helen", Location: "SG"})

	r := setupTestRouter()

	req := httptest.NewRequest("DELETE", "/employees/123", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.Code)
	}
}

func TestAddEmployeeHandler_InvalidJSON(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest("POST", "/employees/",
		strings.NewReader(`{"id":"not-an-int","name":123}`)) // Invalid JSON
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", resp.Code)
	}
}

func TestDeleteEmployeeHandler_NotFound(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest("DELETE", "/employees/9999", nil) // ID doesn't exist
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200 OK with 'not found', got %d", resp.Code)
	}
}
