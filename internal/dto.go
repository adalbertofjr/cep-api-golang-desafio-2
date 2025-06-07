package dto

type CepDTO struct {
	Cep    string `json:"cep"`
	Estado string `json:"state"`
	Cidade string `json:"city"`
	Bairro string `json:"neighborhood"`
	Rua    string `json:"street"`
	Api    string `json:"api"`
}
