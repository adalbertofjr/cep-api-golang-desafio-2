package api

import (
	"fmt"
	"net/http"

	"github.com/adalbertofjr/cep-api-golang-desafio-2/pkg/net"
)

func GetCepBrasilApi(cep string, w http.ResponseWriter, r *http.Request) string {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	return net.FetchData(url, w, r)
}

func GetCepViaCep(cep string, w http.ResponseWriter, r *http.Request) string {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	return net.FetchData(url, w, r)
}
