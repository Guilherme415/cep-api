package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/mock"
	"github.com/Guilherme415/cep-api/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestGetAddressDeitalsByCEP(t *testing.T) {
	fakeError := errors.New("timeout has occured")

	fakeErrorBody, err := json.Marshal(map[string]interface{}{
		"Code":    http.StatusInternalServerError,
		"Message": "an error has occurred",
	})
	require.NoError(t, err)

	fakeGetAddressDeitalsByCEPResponse := response.GetAddressDeitalsByCEPResponse{}
	gofakeit.Struct(&fakeGetAddressDeitalsByCEPResponse)

	fakeCepResult := dto.Viacep{
		Logradouro: fakeGetAddressDeitalsByCEPResponse.Street,
		Bairro:     fakeGetAddressDeitalsByCEPResponse.Neighborhood,
		Localidade: fakeGetAddressDeitalsByCEPResponse.City,
		Uf:         fakeGetAddressDeitalsByCEPResponse.State,
	}

	fakeSuccessBody, err := json.Marshal(fakeCepResult)
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		Description      string
		Client           service.IClient
		ExpectedResponse response.GetAddressDeitalsByCEPResponse
		ExpectedError    error
	}{
		{
			Description:      "Should not return an error",
			Client:           &mock.IClientSpy{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(fakeSuccessBody))},
			ExpectedResponse: fakeGetAddressDeitalsByCEPResponse,
		},
		{
			Description:      "Should return an error when an internal server error ocurred",
			Client:           &mock.IClientSpy{StatusCode: http.StatusInternalServerError, Body: io.NopCloser(bytes.NewReader(fakeErrorBody))},
			ExpectedResponse: response.GetAddressDeitalsByCEPResponse{},
			ExpectedError:    fakeError,
		},
		{
			Description:      "Should return an error when unmarshal invalid body",
			Client:           &mock.IClientSpy{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("blabla"))},
			ExpectedResponse: response.GetAddressDeitalsByCEPResponse{},
			ExpectedError:    fakeError,
		},
		{
			Description:      "Should return an error when an error with http client",
			Client:           &mock.IClientSpy{Err: fakeError},
			ExpectedResponse: response.GetAddressDeitalsByCEPResponse{},
			ExpectedError:    fakeError,
		},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			cepService := service.NewCepService[dto.Viacep]("abc/?/json/", test.Client)

			chanResp := make(chan response.GetAddressDeitalsByCEPResponse)
			ctx := context.Background()

			go cepService.GetAddressDeitalsByCEP("04726-906", ctx, chanResp)

			resp, err := listenChan(chanResp)
			require.Equal(t, test.ExpectedResponse, resp)
			require.Equal(t, test.ExpectedError, err)
		})
	}
}

func listenChan(chanResp <-chan response.GetAddressDeitalsByCEPResponse) (response.GetAddressDeitalsByCEPResponse, error) {
	ticker := time.NewTicker(100 * time.Microsecond)
	defer ticker.Stop()
	for {
		select {
		case resp := <-chanResp:
			return resp, nil
		case <-ticker.C:
			return response.GetAddressDeitalsByCEPResponse{}, errors.New("timeout has occured")
		}
	}
}
