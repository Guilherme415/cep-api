package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetTagValueByField(tagName string, field reflect.StructField) string {
	fieldTag := field.Tag
	tagValue := fieldTag.Get(tagName)

	return tagValue
}

func SetFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int64:
		valueParsed := strings.TrimSpace(value)

		intValue, err := strconv.Atoi(valueParsed)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unknown type: %v", field.Kind())
	}

	return nil
}

func GetStructFieldsNames(obj interface{}) []string {
	t := reflect.TypeOf(obj)

	fieldsQuantity := t.NumField()
	fieldsName := make([]string, fieldsQuantity)

	for i := 0; i < fieldsQuantity; i++ {
		field := t.Field(i)
		fieldsName[i] = field.Name
	}

	return fieldsName
}

func ReplaceLastNonZeroDigitWithZero(input string) string {
	runes := []rune(input)
	length := len(runes) - 1

	for i := length; i >= 0; i-- {
		if runes[i] != '0' && runes[i] != '-' {
			runes[i] = '0'
			break
		}
	}

	return string(runes)
}

func HasNonZeroAndHyphenCharacter(input string) bool {
	for _, char := range input {
		if char != '0' && char != '-' {
			return true
		}
	}
	return false
}
