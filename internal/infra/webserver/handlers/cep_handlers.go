package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetCEP(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to get CEP from Brasil API
	// ctx := context.Background()

	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}

	brasilChannel := make(chan string)
	viaCepChannel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Simulate a delay for the Brasil API call
		// Call the Brasil API
		apiResponse := GetCepBrasilApi(w, r)
		brasilChannel <- apiResponse
	}()

	go func() {
		// If Brasil API fails, call ViaCEP
		time.Sleep(3 * time.Second) // Simulate a delay for the ViaCEP API call
		apiResponse := GetCepViaCep(w, r)
		viaCepChannel <- apiResponse
	}()

	for {
		select {
		case apiResponse := <-brasilChannel:
			if apiResponse != "" {
				println("Response from Brasil API")
				fmt.Fprint(w, apiResponse)
				return
			}
		case viaCepResponse := <-viaCepChannel:
			if viaCepResponse != "" {
				println("Response from ViaCEP")
				fmt.Fprint(w, viaCepResponse)
				return
			}
		default:
			// If no response is received, continue waiting
			continue

		}
	}
}

func GetCepBrasilApi(w http.ResponseWriter, r *http.Request) string {
	// Implement the logic to get CEP from Brasil API
	ctx := context.Background()

	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return "CEP is required"
	}
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "404 not found", http.StatusNoContent)
		return "Content not found"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "Internal Server Error"
	}

	w.Header().Set("Content-Type", "application/json") // Define o tipo de conteúdo como JSON
	w.WriteHeader(http.StatusOK)                       // Retorna o status 200 OK
	// json.NewEncoder(w).Encode(body)
	// fmt.Fprint(w, string(body))
	return string(body)
}

func GetCepViaCep(w http.ResponseWriter, r *http.Request) string {
	// Implement the logic to get CEP from Brasil API
	ctx := context.Background()

	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return "CEP is required"
	}
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return "404 not found"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "404 not found", http.StatusNoContent)
		return "Content not found"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "Internal Server Error"
	}

	w.Header().Set("Content-Type", "application/json") // Define o tipo de conteúdo como JSON
	w.WriteHeader(http.StatusOK)                       // Retorna o status 200 OK
	// json.NewEncoder(w).Encode(body)
	// fmt.Fprint(w, string(body))
	return string(body)
}
