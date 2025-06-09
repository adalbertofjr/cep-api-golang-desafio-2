package handlers

import (
	"net/http"

	"github.com/adalbertofjr/cep-api-golang-desafio-2/internal/entity"
)

func GetCEP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}
	entity.BuscaCep(cep, w, r)
}
