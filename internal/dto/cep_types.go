package dto

import (
	"fmt"
	"reflect"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/utils"
)

type Cep_types interface {
	Viacep | BrasilApi | CdnApiCep
}

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro" type:"Street"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro" type:"Neighborhood"`
	Localidade  string `json:"localidade" type:"City"`
	Uf          string `json:"uf" type:"State"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state" type:"State"`
	City         string `json:"city" type:"City"`
	Neighborhood string `json:"neighborhood" type:"Neighborhood"`
	Street       string `json:"street" type:"Street"`
	Service      string `json:"service"`
}

type CdnApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state" type:"State"`
	City       string `json:"city" type:"City"`
	District   string `json:"district" type:"Neighborhood"`
	Address    string `json:"address" type:"Street"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func MapperToCepResponse[T Cep_types](cepType T) response.GetAddressDeitalsByCEPResponse {
	response := response.GetAddressDeitalsByCEPResponse{}

	structFieldsNames := utils.GetStructFieldsNames(response)

	responseReflected := reflect.ValueOf(&response).Elem()

	cepTypeObj := reflect.ValueOf(&cepType).Elem()

	cepTypeObjType := reflect.TypeOf(cepType)
	if cepTypeObjType.Kind() != reflect.Struct {
		fmt.Println("populatingGetNetFileLinesByTag / cepTypeObjType.Kind() != reflect.Struct")
		return response
	}

	numFields := cepTypeObj.NumField()

	for _, fieldName := range structFieldsNames {
		for i := 0; i < numFields; i++ {
			field := cepTypeObjType.Field(i)
			fieldValue := cepTypeObj.Field(i)

			tagValue := utils.GetTagValueByField("type", field)

			if tagValue == fieldName {
				responseField := responseReflected.FieldByName(fieldName)

				utils.SetFieldValue(responseField, fieldValue.String())
			}
		}
	}

	return response
}
