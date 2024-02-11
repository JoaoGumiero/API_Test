package api

import (
	"net/http"
)

func UploadRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	mux.HandleFunc("/newProduct", NewProductHandler)
	mux.HandleFunc("POST /products", InsertProductHandler)
	mux.HandleFunc("DELETE /products", DeleteProductHandler)
	mux.HandleFunc("PATCH /products", EditProductHandler)
	mux.HandleFunc("PUT /products", UpdateProductHandler)

	return mux
}
