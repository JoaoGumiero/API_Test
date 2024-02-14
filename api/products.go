package api

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/JoaoGumiero/Crud/postgres"
)

// trazer todos os templates da pasta
var temp = template.Must(template.ParseGlob("templates/*.html"))

// Arrumar essa função aqui junto com o HTML
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		// log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
	allProducts, err := postgres.SearchAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProductHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		// log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
	var product postgres.Product
	product.Name = r.FormValue("name")
	product.Description = r.FormValue("description")
	price := r.FormValue("price") // Here, the price and amount fields are retrieved in the string format, so we have to parse them to their respective field types.
	amount := r.FormValue("amount")
	priceFloated, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Failed to convert the price object", err)
	}
	amountInted, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to convert the amount object", err)
	}
	product.Price = priceFloated
	product.Amount = amountInted
	postgres.CreateProduct(product)
	http.Redirect(w, r, "/products", http.StatusFound)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete{
		// log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
	productID := r.PathValue("id")
	postgres.DeleteProduct(productID)
	http.Redirect(w, r, "/products", http.StatusFound)
}

func EditProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch{
		// log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
	productID := r.PathValue("id")
	product := postgres.EditProduct(productID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut{
		// log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price") // Here, the price and amount fields are retrieved in the string format, so we have to parse them to their respective field types.
	amount := r.FormValue("amount")

	idInted, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to convert the amount object", err)
	}
	priceFloated, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Failed to convert the price object", err)
	}
	amountInted, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to convert the amount object", err)
	}
	postgres.UpdateProduct(idInted, name, description, priceFloated, amountInted)
	http.Redirect(w, r, "/products", http.StatusFound)
}
