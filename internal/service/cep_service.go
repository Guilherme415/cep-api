package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
)

type ICepService interface {
	GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- response.GetAddressDeitalsByCEPResponse)
}

type CepService[T dto.Cep_types] struct {
	url    string
	client IClient
}

func NewCepService[T dto.Cep_types](url string, client IClient) ICepService {
	return &CepService[T]{url, client}
}

func (c *CepService[T]) GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- response.GetAddressDeitalsByCEPResponse) {
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

	responseChan <- response
}

func (c *CepService[T]) formatUrlToFindCep(cep string) string {
	return strings.Replace(c.url, "?", cep, -1)
}
