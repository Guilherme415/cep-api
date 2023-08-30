package utils_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Guilherme415/cep-api/utils"
	"github.com/stretchr/testify/require"
)

type Teste struct {
	Abc        string `teste:"abc"`
	TesteInt   int64
	TesteFloat float64
	TesteBool  bool
}

func TestGetTagValueByField(t *testing.T) {
	t.Run("Success - GetTagValueByField should return tag value", func(t *testing.T) {
		expectedValue := "abc"

		testeObj := Teste{
			Abc: expectedValue,
		}

		testObjType := reflect.TypeOf(testeObj)

		field := testObjType.Field(0)

		tagValue := "teste"
		response := utils.GetTagValueByField(tagValue, field)

		require.Equal(t, expectedValue, response)
	})
}

func TestSetFieldValue(t *testing.T) {
	t.Run("Success - SetFieldValue should set field value with string type", func(t *testing.T) {
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "Abc"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := "teste"
		err := utils.SetFieldValue(testeObjField, newValue)

		require.NoError(t, err)
		require.Equal(t, testeObj.Abc, newValue)
	})

	t.Run("Success - SetFieldValue should set field value with Int64", func(t *testing.T) {
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteInt"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		var newValue int64 = 123
		err := utils.SetFieldValue(testeObjField, fmt.Sprint(newValue))

		require.NoError(t, err)
		require.Equal(t, testeObj.TesteInt, newValue)
	})

	t.Run("Success - SetFieldValue should set field value with Float64", func(t *testing.T) {
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteFloat"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		var newValue float64 = 123
		err := utils.SetFieldValue(testeObjField, fmt.Sprint(newValue))

		require.NoError(t, err)
		require.Equal(t, testeObj.TesteFloat, newValue)
	})

	t.Run("Success - SetFieldValue should set field value with Boolean", func(t *testing.T) {
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteBool"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := true
		err := utils.SetFieldValue(testeObjField, fmt.Sprint(newValue))

		require.NoError(t, err)
		require.Equal(t, testeObj.TesteBool, newValue)
	})

	t.Run("Fail - SetFieldValue should return an error when field does not found", func(t *testing.T) {
		expectedError := errors.New("unknown type: invalid")
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "Abcd"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := "teste"
		err := utils.SetFieldValue(testeObjField, newValue)

		require.EqualError(t, expectedError, err.Error())
	})

	t.Run("Fail - SetFieldValue should return an error when parse string to int64", func(t *testing.T) {
		expectedError := errors.New("strconv.Atoi: parsing \"teste\": invalid syntax")
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteInt"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := "teste"
		err := utils.SetFieldValue(testeObjField, newValue)

		require.EqualError(t, expectedError, err.Error())
	})

	t.Run("Fail - SetFieldValue should return an error when parse string to float64", func(t *testing.T) {
		expectedError := errors.New("strconv.ParseFloat: parsing \"teste\": invalid syntax")
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteFloat"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := "teste"
		err := utils.SetFieldValue(testeObjField, newValue)

		require.EqualError(t, expectedError, err.Error())
	})

	t.Run("Fail - SetFieldValue should return an error when parse string to bool", func(t *testing.T) {
		expectedError := errors.New("strconv.ParseBool: parsing \"teste\": invalid syntax")
		testeObj := Teste{}

		testeObjReflected := reflect.ValueOf(&testeObj).Elem()
		fieldName := "TesteBool"
		testeObjField := testeObjReflected.FieldByName(fieldName)

		newValue := "teste"
		err := utils.SetFieldValue(testeObjField, newValue)

		require.EqualError(t, expectedError, err.Error())
	})
}

func TestGetStructFieldsNames(t *testing.T) {
	t.Run("Success - GetStructFieldsNames should return each stuct field name", func(t *testing.T) {
		expectedFieldNames := []string{
			"Abc",
			"TesteInt",
			"TesteFloat",
			"TesteBool",
		}

		testeObj := Teste{}

		fieldNames := utils.GetStructFieldsNames(testeObj)

		for i, fieldName := range fieldNames {
			require.Equal(t, fieldName, expectedFieldNames[i])
		}

		require.Equal(t, len(expectedFieldNames), len(fieldNames))
	})
}

func TestReplaceLastNonZeroDigitWithZero(t *testing.T) {
	t.Run("Success - ReplaceLastNonZeroDigitWithZero should change last digit to zero", func(t *testing.T) {
		expectedResult := "04726-900"
		stringTest := "04726-906"

		result := utils.ReplaceLastNonZeroDigitWithZero(stringTest)

		require.Equal(t, expectedResult, result)
	})

	t.Run("Success - ReplaceLastNonZeroDigitWithZero should return lasts digits to zero and not change '-'", func(t *testing.T) {
		expectedResult := "04700-000"
		stringTest := "04720-000"

		result := utils.ReplaceLastNonZeroDigitWithZero(stringTest)

		require.Equal(t, expectedResult, result)
	})
}

func TestHasNonZeroAndHyphenCharacter(t *testing.T) {
	t.Run("Success - HasNonZeroAndHyphenCharacter should return true when has non zero or hyphens caracters", func(t *testing.T) {
		stringTest := "04726-906"

		result := utils.HasNonZeroAndHyphenCharacter(stringTest)

		require.True(t, result)
	})

	t.Run("Success - HasNonZeroAndHyphenCharacter should return false when only has zero caracters", func(t *testing.T) {
		stringTest := "0000000"

		result := utils.HasNonZeroAndHyphenCharacter(stringTest)

		require.False(t, result)
	})

	t.Run("Success - HasNonZeroAndHyphenCharacter should return false when only has zero and hyphens caracters", func(t *testing.T) {
		stringTest := "00000-000"

		result := utils.HasNonZeroAndHyphenCharacter(stringTest)

		require.False(t, result)
	})
}
