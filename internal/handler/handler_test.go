package handler_test

import (
	"go-native-http/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	testcases := []struct {
		name           string
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "not_implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/products", nil)
			if err != nil {
				t.Fatalf("unable to create new test request: %v", err)
			}
			//Act
			handler.GetAllProducts(nil)(w, r)
			//Assert
			assert.Equal(t, tc.expectedStatus, w.Result().StatusCode)
		})
	}

}

func TestGetProductById(t *testing.T) {
	testcases := []struct {
		name           string
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "not_implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/products/123", nil)
			handler.GetProductById(nil)(w, r)
			assert.Equal(t, tc.expectedStatus, w.Result().StatusCode)
		})
	}
}

func TestAddProduct(t *testing.T) {
	testcases := []struct {
		name           string
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "not_implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/products", nil)
			handler.AddProduct(nil)(w, r)
			assert.Equal(t, tc.expectedStatus, w.Result().StatusCode)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testcases := []struct {
		name           string
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "not_implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPut, "/product/123", nil)
			handler.UpdateProductById(nil)(w, r)
			assert.Equal(t, tc.expectedStatus, w.Result().StatusCode)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	testcases := []struct {
		name           string
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "not_implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, "/products/123", nil)
			handler.DeleteProductById(nil)(w, r)
			assert.Equal(t, tc.expectedStatus, w.Result().StatusCode)
		})
	}
}
