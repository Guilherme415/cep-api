package internal

import (
	"github.com/Guilherme415/cep-api/internal/api/controller"
	usecases "github.com/Guilherme415/cep-api/internal/api/use_cases"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/service"
)

var (
	ViacepController controller.ICepController
)

func LoadDependencies() {
	viacepService := service.NewCepService[dto.Viacep]("https://viacep.com.br/ws/?/json/")
	brasilAbertoService := service.NewCepService[dto.BrasilAberto]("https://brasilaberto.com/api/v1/zipcode/?")

	cepServices := []service.ICepService{viacepService, brasilAbertoService}
	cepUseCase := usecases.NewCepUseCase(cepServices)

	ViacepController = controller.NewCepController(cepUseCase)
}
