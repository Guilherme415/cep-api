package mock

import (
	"context"

	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/service"
)

type ICepServiceSpy struct {
	service.ICepService
	GetAddressDeitalsByCEPResponse dto.CepServiceResponse
}

func (c ICepServiceSpy) GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- dto.CepServiceResponse) {
	responseChan <- c.GetAddressDeitalsByCEPResponse
}
