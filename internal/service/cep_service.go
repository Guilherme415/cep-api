package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/dto/mapper"
	"github.com/Guilherme415/cep-api/utils"
	"github.com/rs/zerolog/log"
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
		log.Err(err).Msgf("GetAddressDeitalsByCEP - Error to create request, url: %s", url)
		return
	}

	resp, err := c.client.Do(request)
	if err != nil {
		log.Err(err).Msgf("GetAddressDeitalsByCEP - Error to client execute, url: %s", url)
		return
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			c.processNotFoundStatus(cep, ctx, responseChan)
			return
		}

		log.Info().Msgf("GetAddressDeitalsByCEP - Status code different from 200 and 404, statusCode: %d, url: %s", resp.StatusCode, url)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msgf("GetAddressDeitalsByCEP - Error to read body, url: %s", url)
		return
	}

	err = json.Unmarshal(body, &requestResponse)
	if err != nil {
		log.Err(err).Msgf("GetAddressDeitalsByCEP - Error to unmarshal body, url: %s", url)
		return
	}

	response := mapper.MapperToCepResponse[T](requestResponse)
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
	cepResponse.Error = errors.New("cep not found")
	responseChan <- cepResponse
}
