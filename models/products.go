package models

import (
	"log"

	"github.com/JoaoGumiero/Crud/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchAllProducts() []Product {
	db := db.DbConnection()

	sql, err := db.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}
	p := Product{}
	Products := []Product{}

	for sql.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := sql.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Amount = amount
		p.Price = price
		Products = append(Products, p)
	}
	defer db.Close()
	return Products
}

func CreateProduct(name, description string, price float64, amount int) {
	db := db.DbConnection()
	insertDataIntoDB, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataIntoDB.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnection()

	deleteDataFromDB, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteDataFromDB.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DbConnection()

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
	db := db.DbConnection()
	editDataFromDB, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	editDataFromDB.Exec(name, description, price, amount, id)
	defer db.Close()
}
