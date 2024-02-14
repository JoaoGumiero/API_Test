package api

import (
	"net/http"
)

func UploadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", getAllProductsHandler)
	mux.HandleFunc("/newProduct", NewProductHandler) //handler for only the newproduct page
	mux.HandleFunc("POST /products", InsertProductHandler)
	mux.HandleFunc("DELETE /products/{id}", DeleteProductHandler)
	mux.HandleFunc("PATCH /products/{id}", EditProductHandler)
	mux.HandleFunc("PUT /products", UpdateProductHandler)

	return mux
}
