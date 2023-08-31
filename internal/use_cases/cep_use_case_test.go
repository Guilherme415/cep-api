package usecases_test

import (
	"errors"
	"testing"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/mock"
	"github.com/Guilherme415/cep-api/internal/service"
	usecases "github.com/Guilherme415/cep-api/internal/use_cases"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCepUseCase(t *testing.T) {
	t.Run("Success - GetAddressDeitalsByCEP should return no error", func(t *testing.T) {
		fakeServiceResponse := dto.CepServiceResponse{}
		gofakeit.Struct(&fakeServiceResponse)

		expectedResult := fakeServiceResponse.GetAddressDeitalsByCEPResponse

		cepService := mock.ICepServiceSpy{GetAddressDeitalsByCEPResponse: fakeServiceResponse}
		cepServices := []service.ICepService{cepService}

		fakeCep := "04726-906"

		useCase := usecases.NewCepUseCase(cepServices)
		response, err := useCase.GetAddressDeitalsByCEP(fakeCep)

		require.NoError(t, err)
		require.Equal(t, expectedResult, response)
	})

	t.Run("Fail - GetAddressDeitalsByCEP should return an timeout error", func(t *testing.T) {
		expectedError := errors.New("timeout has occured")
		expectedResult := response.GetAddressDeitalsByCEPResponse{}

		cepService := mock.ICepServiceSpy{}
		cepServices := []service.ICepService{cepService}

		fakeCep := "04726-906"

		useCase := usecases.NewCepUseCase(cepServices)
		response, err := useCase.GetAddressDeitalsByCEP(fakeCep)

		require.EqualError(t, expectedError, err.Error())
		require.Equal(t, expectedResult, response)
	})

	t.Run("Success - GetAddressDeitalsByCEP should return no error with more than one service and return empty struct in the first one", func(t *testing.T) {
		fakeServiceResponseEmpty := dto.CepServiceResponse{}
		fakeServiceResponseWithData := dto.CepServiceResponse{}
		gofakeit.Struct(&fakeServiceResponseWithData)
		expectedResult := fakeServiceResponseWithData.GetAddressDeitalsByCEPResponse

		cepService := mock.ICepServiceSpy{GetAddressDeitalsByCEPResponse: fakeServiceResponseEmpty}
		viacepService := mock.ICepServiceSpy{GetAddressDeitalsByCEPResponse: fakeServiceResponseWithData}
		cepServices := []service.ICepService{cepService, viacepService}

		fakeCep := "04726-906"

		useCase := usecases.NewCepUseCase(cepServices)
		response, err := useCase.GetAddressDeitalsByCEP(fakeCep)

		require.NoError(t, err)
		require.Equal(t, expectedResult, response)
	})
}
