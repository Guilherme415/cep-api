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
	brasilAbertoService := service.NewCepService[dto.BrasilApi]("https://brasilapi.com.br/api/cep/v2/?")
	CdnApiCepService := service.NewCepService[dto.CdnApiCep]("https://cdn.apicep.com/file/apicep/?.json")
	OpenCepService := service.NewCepService[dto.Viacep]("https://opencep.com/v1/?")

	cepServices := []service.ICepService{viacepService, brasilAbertoService, CdnApiCepService, OpenCepService}
	cepUseCase := usecases.NewCepUseCase(cepServices)

	ViacepController = controller.NewCepController(cepUseCase)
}
