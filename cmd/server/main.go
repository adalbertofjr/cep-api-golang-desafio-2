package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adalbertofjr/cep-api-golang-desafio-2/internal/infra/webserver/handlers"
)

const porta = ":8080"

func main() {
	startServer()
}

func startServer() {
	fmt.Println("Servidor iniciado na porta", porta)
	fmt.Println("Acesse http://localhost:8080/?cep=04446160 para buscar o CEP")
	fmt.Println("Pressione Ctrl+C para parar o servidor")

	http.HandleFunc("/", handlers.GetCEP)
	error := http.ListenAndServe(porta, nil)
	if error != nil {
		log.Fatal("Erro ao iniciar o servidor: ", error)
	}
}
