package dto

import "github.com/Guilherme415/cep-api/internal/api/response"

type CepServiceResponse struct {
	response.GetAddressDeitalsByCEPResponse
	Error error
}
