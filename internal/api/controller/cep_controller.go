package controller

import (
	"net/http"

	"github.com/Guilherme415/cep-api/internal/api/response"
	usecases "github.com/Guilherme415/cep-api/internal/use_cases"
	"github.com/gin-gonic/gin"
)

type ICepController interface {
	GetAddressDeitalsByCEP(c *gin.Context)
}

type CepController struct {
	cepUseCase usecases.ICepUseCase
}

func NewCepController(cepUseCase usecases.ICepUseCase) ICepController {
	return &CepController{cepUseCase}
}

// ShowAccount godoc
// @Summary      Get address details by CEP
// @Description  Api to get address details by cep. This Api try to get the infos in one or more APIs and return the fastest
// @Description  If the Cep is invalid, the api will replace the lasts digits to 0 until find some valid Cep
// @Description  If does not find some valid cep, it will return an internal server error
// @Tags         cep
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true  "CEP"
// @Success      200  {object}  response.GetAddressDeitalsByCEPResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /cep/{cep} [get]
// @Security ApiKeyAuth
func (cp *CepController) GetAddressDeitalsByCEP(c *gin.Context) {
	cep := c.Param("cep")
	if isCepInvalid(cep) {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "parâmetro cep é obrigatório e deve possuir de 5 a 9 dígitos",
		})
		return
	}

	addressDetails, err := cp.cepUseCase.GetAddressDeitalsByCEP(cep)
	if err != nil {
		message := "um erro interno ocorreu"
		if err.Error() == "cep not found" {
			message = "CEP inválido"
		}

		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, addressDetails)
}

func isCepInvalid(cep string) bool {
	maxCepLength := 9
	minCepLength := 5

	return cep == "" || len(cep) > maxCepLength || len(cep) < minCepLength
}
