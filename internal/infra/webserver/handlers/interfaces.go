package handlers

import "github.com/adalbertofjr/cep-api-golang-desafio-2/internal/entity"

type CEPInterface interface {
	GetCEP(cep string) (*entity.CEP, error)
}
