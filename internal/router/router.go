package router

import (
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func (r *Router) New() *http.ServeMux {
	r.Mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {})
	r.Mux.HandleFunc("GET /products/{product_id}", func(w http.ResponseWriter, r *http.Request) {})
	r.Mux.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {})
	r.Mux.HandleFunc("PUT /products/{product_id}", func(w http.ResponseWriter, r *http.Request) {})
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
