package mapper

import (
	"reflect"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/dto"
	"github.com/Guilherme415/cep-api/utils"
)

type cepStructureReflect struct {
	obj        reflect.Value
	structType reflect.Type
}

func newCepStructureReflect(cepStructure interface{}) cepStructureReflect {
	cepStructureObj := reflect.ValueOf(cepStructure)
	cepStructureType := reflect.TypeOf(cepStructure)

	return cepStructureReflect{
		obj:        cepStructureObj,
		structType: cepStructureType,
	}
}

func MapperToCepResponse[T dto.Cep_types](cepStructure T) response.GetAddressDeitalsByCEPResponse {
	getAddressDeitalsByCEPResponse := response.GetAddressDeitalsByCEPResponse{}
	getAddressDeitalsByCEPResponseFieldsNames := utils.GetStructFieldsNames(getAddressDeitalsByCEPResponse)
	getAddressDeitalsByCEPResponseReflected := reflect.ValueOf(&getAddressDeitalsByCEPResponse).Elem()

	cepStructureReflect := newCepStructureReflect(cepStructure)

	for _, fieldName := range getAddressDeitalsByCEPResponseFieldsNames {
		populateResponseField(cepStructureReflect, fieldName, getAddressDeitalsByCEPResponseReflected)
	}

	return getAddressDeitalsByCEPResponse
}

func populateResponseField(cepStructurereflect cepStructureReflect, fieldName string, responseReflected reflect.Value) {
	numFields := cepStructurereflect.obj.NumField()

	for i := 0; i < numFields; i++ {
		cepField := cepStructurereflect.structType.Field(i)
		cepFieldValue := cepStructurereflect.obj.Field(i)

		if isMatchingTypeTag(cepField, fieldName) {
			setResponseField(responseReflected, fieldName, cepFieldValue.String())
		}
	}
}

func isMatchingTypeTag(cepField reflect.StructField, fieldName string) bool {
	return utils.GetTagValueByField("type", cepField) == fieldName
}

func setResponseField(responseReflected reflect.Value, fieldName string, value string) {
	responseField := responseReflected.FieldByName(fieldName)
	if responseField.IsValid() {
		utils.SetFieldValue(responseField, value)
	}
}
