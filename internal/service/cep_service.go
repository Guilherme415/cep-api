package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/utils"
)

type ICepService interface {
	GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- dto.CepServiceResponse)
}

type CepService[T dto.Cep_types] struct {
	url    string
	client IClient
}

func NewCepService[T dto.Cep_types](url string, client IClient) ICepService {
	return &CepService[T]{url, client}
}

func (c *CepService[T]) GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- dto.CepServiceResponse) {
	var requestResponse T
	url := c.formatUrlToFindCep(cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			c.processNotFoundStatus(cep, ctx, responseChan)
			return
		}

		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &requestResponse)
	if err != nil {
		return
	}

	response := dto.MapperToCepResponse[T](requestResponse)
	cepResponse := dto.CepServiceResponse{
		GetAddressDeitalsByCEPResponse: response,
	}

	responseChan <- cepResponse
}

func (c *CepService[T]) formatUrlToFindCep(cep string) string {
	return strings.Replace(c.url, "?", cep, -1)
}

func (c *CepService[T]) processNotFoundStatus(cep string, ctx context.Context, responseChan chan<- dto.CepServiceResponse) {
	if utils.HasNonZeroAndHyphenCharacter(cep) {
		cep = utils.ReplaceLastNonZeroDigitWithZero(cep)
		c.GetAddressDeitalsByCEP(cep, ctx, responseChan)
		return
	}

	cepResponse := dto.CepServiceResponse{}
	cepResponse.Error = fmt.Errorf("cep not found, cep: %s", cep)
	responseChan <- cepResponse
}
