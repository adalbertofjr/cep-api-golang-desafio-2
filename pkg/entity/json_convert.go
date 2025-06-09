package entity

import (
	"encoding/json"
	"log"

	dto "github.com/adalbertofjr/cep-api-golang-desafio-2/internal"
)

func RequestToJson(body []byte) dto.CepDTO {
	var cepDTO dto.CepDTO
	err := json.Unmarshal(body, &cepDTO)
	if err != nil {
		log.Printf("Erro ao fazer o Unmarshal: %v\n", err)
		return cepDTO
	}

	return cepDTO
}
