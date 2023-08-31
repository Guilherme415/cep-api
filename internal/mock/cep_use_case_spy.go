package mock

import (
	"github.com/Guilherme415/cep-api/internal/api/response"
	usecases "github.com/Guilherme415/cep-api/internal/use_cases"
)

type ICepUseCaseSpy struct {
	usecases.ICepUseCase
	GetAddressDeitalsByCEPResponse response.GetAddressDeitalsByCEPResponse
	GetAddressDeitalsByCEPError    error
}

func (c *ICepUseCaseSpy) GetAddressDeitalsByCEP(cep string) (response.GetAddressDeitalsByCEPResponse, error) {
	return c.GetAddressDeitalsByCEPResponse, c.GetAddressDeitalsByCEPError
}
