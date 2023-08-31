package mapper_test

import (
	"testing"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/dto/mapper"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCepUseCase(t *testing.T) {
	t.Run("Success - MapperToCepResponse should populating GetAddressDeitalsByCEPResponse struct", func(t *testing.T) {
		var viaCepResponse dto.Viacep
		gofakeit.Struct(&viaCepResponse)

		expectedResult := response.GetAddressDeitalsByCEPResponse{
			Street:       viaCepResponse.Logradouro,
			Neighborhood: viaCepResponse.Bairro,
			City:         viaCepResponse.Localidade,
			State:        viaCepResponse.Uf,
		}

		response := mapper.MapperToCepResponse[dto.Viacep](viaCepResponse)

		require.Equal(t, expectedResult, response)
	})
}
