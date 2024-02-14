package main

import (
	"net/http"

	"github.com/JoaoGumiero/Crud/api"
)

func main() {
	http.ListenAndServe(":8000", api.UploadRoutes())
}
