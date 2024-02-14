package postgres

import (
	"log"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchAllProducts() ([]Product, error ) {
	db := DbConnection()

	sql, err := db.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	Products := []Product{}

	for sql.Next() {
		p := Product{}
		err := sql.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Amount)
		if err != nil {
			panic(err.Error())
		}
		Products = append(Products, p)
	}
	return Products, nil
}

func CreateProduct(product Product) {
	db := DbConnection()
	_,err := db.Exec("insert into products(name, description, price, amount) values($1, $2, $3, $4)",
product.Name, product.Description, product.Price, product.Amount)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func DeleteProduct(id string) {
	db := DbConnection()

	deleteDataFromDB, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteDataFromDB.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := DbConnection()

	ProductInBD, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	p := Product{}

	for ProductInBD.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := ProductInBD.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Amount = amount
		p.Price = price
	}
	defer db.Close()
	return p
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := DbConnection()
	editDataFromDB, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	editDataFromDB.Exec(name, description, price, amount, id)
	defer db.Close()
}
