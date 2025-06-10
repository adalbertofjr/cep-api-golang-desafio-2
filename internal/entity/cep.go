package entity

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adalbertofjr/cep-api-golang-desafio-2/internal/infra/api"
)

func BuscaCep(cep string, w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	brasilChannel := make(chan string)
	viaCepChannel := make(chan string)

	go func() {
		brasilChannel <- api.GetCepBrasilApi(cep, w, r)
	}()

	go func() {
		viaCepChannel <- api.GetCepViaCep(cep, w, r)
	}()

	select {
	case <-ctx.Done():
		log.Println("Context cancelled")
		cancel()
		return
	case <-time.After(1 * time.Second):
		log.Println("Timeout reached, no response received")
		cancel()
		return
	case apiResponse := <-brasilChannel:
		println("Response from Brasil API: ", apiResponse)
		cancel()
		return
	case apiResponse := <-viaCepChannel:
		println("Response from ViaCEP API: ", apiResponse)
		cancel()
		return
	}
}
