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

}

func TestAddProduct(t *testing.T) {

}

func TestUpdateProduct(t *testing.T) {

}

func TestDeleteProduct(t *testing.T) {

}
