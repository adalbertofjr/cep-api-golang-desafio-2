package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

const porta = ":8080"

func main() {
	startServer()
}

func startServer() {
	fmt.Println("Servidor iniciado na porta", porta)
	fmt.Println("Acesse http://localhost:8080/cep para buscar a cotação do dólar")
	fmt.Println("Pressione Ctrl+C para parar o servidor")

	http.HandleFunc("/cep", BuscaCepHandler)
	error := http.ListenAndServe(porta, nil)
	if error != nil {
		log.Fatal("Erro ao iniciar o servidor: ", error)
	}
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	// Faz a requisição HTTP
	ctx := context.Background()
	path1 := "https://brasilapi.com.br/api/cep/v1/04446160"
	// path2 := "http://viacep.com.br/ws/04446160/json/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path1, nil)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "404 not found", http.StatusNoContent)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Aqui você pode implementar a lógica para buscar a cotação do dólar
	fmt.Fprint(w, string(body))
}
