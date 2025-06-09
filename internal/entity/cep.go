package entity

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adalbertofjr/cep-api-golang-desafio-2/internal/infra/api"
)

// type CEP struct {
// 	Cep          string `json:"cep"`
// 	State        string `json:"state"`
// 	City         string `json:"city"`
// 	Neighborhood string `json:"neighborhood"`
// 	Street       string `json:"street"`
// 	Api          string `json:"api"`
// }

func BuscaCep(cep string, w http.ResponseWriter, r *http.Request) string {
	brasilChannel := make(chan string)
	viaCepChannel := make(chan string)

	go func() {
		// apiResponse := api.GetCepBrasilApi(cep, w, r)
		brasilChannel <- api.GetCepBrasilApi(cep, w, r)
	}()

	go func() {
		// apiResponse := api.GetCepViaCep(cep, w, r)
		viaCepChannel <- api.GetCepViaCep(cep, w, r)
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("Erro: Timeout")
		case apiResponse := <-brasilChannel:
			println("Response from Brasil API")
			fmt.Fprint(w, apiResponse)
			// RequestToJson([]byte(apiResponse))
		case apiResponse := <-viaCepChannel:
			println("Response from ViaCEP API")
			fmt.Fprint(w, apiResponse)
			// RequestToJson([]byte(apiResponse))
		default:
			// If no response is received, continue waiting
			continue
		}
	}
}
