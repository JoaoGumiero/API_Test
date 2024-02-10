package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaog/API_Test/handlers"
)


func UploadRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	// Define your HTTP routes using the router
	r.HandleFunc("/newProduct", handlers.NewProductHandler)
	r.HandleFunc("/insert", handlers.InsertProductHandler)
	r.HandleFunc("/delete", handlers.DeleteProductHandler)
	r.HandleFunc("/edit", handlers.EditProductHandler)
	r.HandleFunc("/update", handlers.UpdateProductHandler)

	log.Println("Server listening on:8000")
	http.ListenAndServe(":8000", r)
}
