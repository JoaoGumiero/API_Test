package routes

import (
	"net/http"

	"github.com/JoaoGumiero/Crud/handlers"
)

func UploadRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)

	mux.HandleFunc("/newProduct", handlers.NewProductHandler)
	mux.HandleFunc("POST /products", handlers.InsertProductHandler)
	mux.HandleFunc("DELETE /products", handlers.DeleteProductHandler)
	mux.HandleFunc("PATCH /products", handlers.EditProductHandler)
	mux.HandleFunc("PUT /products", handlers.UpdateProductHandler)

	return mux
}
