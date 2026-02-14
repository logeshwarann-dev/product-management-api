package router

import (
	"go-native-http/internal/handler"
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func (r *Router) New(ps handler.ProductStore) *http.ServeMux {
	r.Mux.HandleFunc("GET /products", handler.GetAllProducts(ps))
	r.Mux.HandleFunc("GET /products/{product_id}", handler.GetProductById(ps))
	r.Mux.HandleFunc("POST /products", handler.AddProduct(ps))
	r.Mux.HandleFunc("PUT /products/{product_id}", handler.UpdateProductById(ps))
	r.Mux.HandleFunc("DELETE /products/{product_id}", handler.DeleteProductById(ps))
	return r.Mux
}

func (r *Router) AddCORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required
		next.ServeHTTP(w, r)
	}
}
