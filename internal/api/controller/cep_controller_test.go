package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/mock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCepController(t *testing.T) {
	t.Run("Success - GetAddressDeitalsByCEP should return 200 ok", func(t *testing.T) {
		expectedResponse := response.GetAddressDeitalsByCEPResponse{}
		gofakeit.Struct(&expectedResponse)

		cepUseCase := &mock.ICepUseCaseSpy{GetAddressDeitalsByCEPResponse: expectedResponse}

		fakeCep := "04726-906"

		controller := controller.NewCepController(cepUseCase)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAddressDeitalsByCEP)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		if err != nil {
			return
		}

		response := response.GetAddressDeitalsByCEPResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}

		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, expectedResponse, response)
	})

	t.Run("Fail - GetAddressDeitalsByCEP should return 400 when cep is invalid", func(t *testing.T) {
		cepUseCase := &mock.ICepUseCaseSpy{}

		fakeCep := "0472"

		controller := controller.NewCepController(cepUseCase)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAddressDeitalsByCEP)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Fail - GetAddressDeitalsByCEP should return 500 when an internal error occurred", func(t *testing.T) {
		expectedError := errors.New("can't connect DB")

		cepUseCase := &mock.ICepUseCaseSpy{GetAddressDeitalsByCEPError: expectedError}

		fakeCep := "04726-906"

		controller := controller.NewCepController(cepUseCase)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAddressDeitalsByCEP)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		if err != nil {
			return
		}

		response := response.GetAddressDeitalsByCEPResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}

		require.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
