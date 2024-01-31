package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	_ "github.com/lib/pq"
)

func dbConnection() *sql.DB {
	connStr := "user=postgres password=123# host=localhost dbname=API_Test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

// trazer todos os templates da pasta
var temp = template.Must(template.ParseGlob("templates/*.html"))

const monitorings = 2
const delay = 5

// Precisa de uma função principal para a nossa aplicação, dentro dela sera inserida o código
func main() {
	//Subir o servidor
	http.HandleFunc("/", index)
	http.ListenAndServe(":800", nil)

	for {
		showMenu()
		command := readUserCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println(" Op 2 - Show logs")
			printLogs()
		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)
		default:
			fmt.Println("Command not recognized, exiting the program")
			os.Exit(-1)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
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

func showMenu() {
	fmt.Println("1 - Run Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Close de App")
}

func readUserCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("The chosed option is: ", commandRead)
	return commandRead
}

func startMonitoring() {
	fmt.Println("Starting ...")
	websites, _ := readWebsites()
	for i := 0; i < monitorings; i++ {
		for i, website := range websites {
			fmt.Println("Testing the following website: ", websites[i])
			testWebsite(website)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testWebsite(website string) {
	resp, err := http.Get(website)
	if err != nil {
		fmt.Println("An error ocurred when calling the website: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Test successful, code: ", resp.StatusCode)
		recordLog(website, true)
	} else {
		fmt.Println("Test not successful, code: ", resp.StatusCode)
		recordLog(website, false)
	}
}

func readWebsites() ([]string, error) {
	var websites []string

	archive, err := os.Open("webSites.txt")
	if err != nil {
		fmt.Println("An error ocurred: ", err)
	}
	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		websites = append(websites, scanner.Text())
	}
	archive.Close()
	return websites, scanner.Err()
}

func recordLog(website string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	archive.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + website + "- Online: " + strconv.FormatBool(status) + "\n")
	archive.Close()
}

func printLogs() {
	// as operações de abrir o arquivo pelo package OS (que mexe a nível de sistema OS)
	// aqui não necessida fechar o arquivo como boas práticas.
	// Porém é deprecated o io.util na versão do go, estão da para utilizar o OS também
	archive, err := os.ReadFile("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Função string ele imprime direto no nosso terminal
	fmt.Println(string(archive))
}
