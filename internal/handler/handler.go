package handler

import (
	"go-native-http/internal/store"
	"net/http"
)

type ProductStore interface {
	// Create a new product
	Create(record store.Product) (store.Product, error)
	// Fetch all products
	GetAll() (store.Products, error)
	// Fetch a product using its id
	GetById(id string) (store.Product, error)
	// Update a product using its id
	UpdateById(id string, record store.Product) (store.Product, error)
	// Delete a product using its id
	DeleteById(id string) error
}

func GetAllProducts(ps ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetProductById(ps ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func AddProduct(ps ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func UpdateProductById(ps ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func DeleteProductById(ps ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
