package models

import (
	"log"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func SearchAllProducts() []Produto {
	db := DbConnection()
	sql, err := db.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}

	p := Produto{}
	produtos := []Produto{}

	for sql.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := sql.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Fatal(err)
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco
	}
	produtos = append(produtos, p)

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
