package response

type GetAddressDeitalsByCEPResponse struct {
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}
