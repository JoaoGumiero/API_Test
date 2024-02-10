package main

import (
	"github.com/JoaoGumiero/Crud/routes"
)

// Precisa de uma função principal para a nossa aplicação, dentro dela sera inserida o código
func main() {
	//Subir o servidor
	routes.UploadRoutes()

	// for {
	// 	sh.ShowMenu()
	// 	command := sh.ReadUserCommand()
	// 	switch command {
	// 	case 1:
	// 		sh.StartMonitoring()
	// 	case 2:
	// 		fmt.Println(" Op 2 - Show logs")
	// 		sh.PrintLogs()
	// 	case 0:
	// 		fmt.Println("Exiting the program")
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("Command not recognized, exiting the program")
	// 		os.Exit(-1)
	// 	}
	// }
}
