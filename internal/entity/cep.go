package entity

import (
	"context"
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
		println("Context cancelled")
		cancel()
	case <-time.After(1 * time.Second):
		println("Timeout reached, no response received")
		cancel()
	case apiResponse := <-brasilChannel:
		println("Response from Brasil API: ", apiResponse)
		cancel()
	case apiResponse := <-viaCepChannel:
		println("Response from ViaCEP API: ", apiResponse)
		cancel()
	}
}
