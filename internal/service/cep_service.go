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
	url string
}

func NewCepService[T dto.Cep_types](url string) ICepService {
	return &CepService[T]{url}
}

func (c *CepService[T]) GetAddressDeitalsByCEP(cep string, ctx context.Context, responseChan chan<- response.GetAddressDeitalsByCEPResponse) {
	var requestResponse T
	url := strings.Replace(c.url, "?", cep, -1)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
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
