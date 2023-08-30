package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/service"
)

type ICepUseCase interface {
	GetAddressDeitalsByCEP(cep string) (response.GetAddressDeitalsByCEPResponse, error)
}

type CepUseCase struct {
	cepServices []service.ICepService
}

func NewCepUseCase(cepServices []service.ICepService) ICepUseCase {
	return &CepUseCase{cepServices}
}

func (cp *CepUseCase) GetAddressDeitalsByCEP(cep string) (response.GetAddressDeitalsByCEPResponse, error) {
	return cp.getFirstCepResponse(cep)
}

func (cp *CepUseCase) getFirstCepResponse(cep string) (response.GetAddressDeitalsByCEPResponse, error) {
	responseChan := make(chan response.GetAddressDeitalsByCEPResponse)

	ctx := context.Background()
	defer ctx.Done()
	for _, service := range cp.cepServices {
		go service.GetAddressDeitalsByCEP(cep, ctx, responseChan)
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case resp := <-responseChan:
			return resp, nil

		case <-ticker.C:
			return response.GetAddressDeitalsByCEPResponse{}, errors.New("timeout has occured")
		}
	}
}
