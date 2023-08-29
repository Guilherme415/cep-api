package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
)

type ICepService interface {
	GetAddressDeitalsByCEP(cep string, responseChan chan<- response.GetAddressDeitalsByCEPResponse)
}

type CepService[T dto.Cep_types] struct {
	url string
}

func NewCepService[T dto.Cep_types](url string) ICepService {
	return &CepService[T]{url}
}

func (c *CepService[T]) GetAddressDeitalsByCEP(cep string, responseChan chan<- response.GetAddressDeitalsByCEPResponse) {
	response := response.GetAddressDeitalsByCEPResponse{}

	url := strings.Replace(c.url, "?", cep, -1)

	request, err := http.NewRequest(http.MethodGet, url, nil)
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

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	responseChan <- response
}
