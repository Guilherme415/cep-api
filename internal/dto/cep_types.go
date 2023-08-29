package dto

import "github.com/Guilherme415/cep-api/internal/api/response"

type Cep_types interface {
	Viacep | BrasilAberto
}

func ConvertToResponse[T Cep_types]() response.GetAddressDeitalsByCEPResponse {
	return response.GetAddressDeitalsByCEPResponse{}
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

type BrasilAberto struct {
	Result BrasilAbertoResult `json:"result"`
}

type BrasilAbertoResult struct {
	Street         string `json:"street" type:"Street"`
	Complement     string `json:"complement"`
	District       string `json:"district" type:"Neighborhood"`
	DistrictID     int    `json:"districtId"`
	City           string `json:"city" type:"City"`
	CityID         int    `json:"cityId"`
	IbgeID         int    `json:"ibgeId"`
	State          string `json:"state" type:"State"`
	StateShortname string `json:"stateShortname"`
	Zipcode        string `json:"zipcode"`
}
