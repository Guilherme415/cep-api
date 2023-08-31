package dependency

import (
	"net/http"

	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/internal/service"
	usecases "github.com/Guilherme415/cep-api/internal/use_cases"
)

var (
	ViacepController controller.ICepController
)

func LoadDependencies() {
	defaultClient := http.DefaultClient

	viacepService := service.NewCepService[dto.Viacep]("https://viacep.com.br/ws/?/json/", defaultClient)
	brasilAbertoService := service.NewCepService[dto.BrasilApi]("https://brasilapi.com.br/api/cep/v2/?", defaultClient)
	CdnApiCepService := service.NewCepService[dto.CdnApiCep]("https://cdn.apicep.com/file/apicep/?.json", defaultClient)
	OpenCepService := service.NewCepService[dto.Viacep]("https://opencep.com/v1/?", defaultClient)

	cepServices := []service.ICepService{viacepService, brasilAbertoService, CdnApiCepService, OpenCepService}
	cepUseCase := usecases.NewCepUseCase(cepServices)

	ViacepController = controller.NewCepController(cepUseCase)
}
