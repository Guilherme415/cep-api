package controller

import (
	"fmt"
	"net/http"

	usecases "github.com/Guilherme415/cep-api/internal/api/use_cases"
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

func (cp *CepController) GetAddressDeitalsByCEP(c *gin.Context) {
	cep := c.Param("cep")
	if isCepInvalid(cep) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "param cep is requirer",
		})
		return
	}

	addressDetails, err := cp.cepUseCase.GetAddressDeitalsByCEP(cep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": fmt.Sprintf("An internal error occured, details: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Response": addressDetails,
	})
}

func isCepInvalid(cep string) bool {
	maxCepLength := 15
	minCepLength := 5

	return cep == "" || len(cep) > maxCepLength || len(cep) < minCepLength
}
