package response

type GetAddressDeitalsByCEPResponse struct {
	Street       string `json:"rua"`
	Neighborhood string `json:"bairro"`
	City         string `json:"cidade"`
	State        string `json:"estado"`
}
