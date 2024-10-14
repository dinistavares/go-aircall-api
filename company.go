package aircall

// Company service
type CompaniesService service

type CompanyResponse struct {
	Company *Company `json:"company,omitempty"`
}

type Company struct {
	Name         string `json:"name,omitempty"`
	UsersCount   int    `json:"users_count,omitempty"`
	NumbersCount int    `json:"numbers_count,omitempty"`
}

//  ***********************************************************************************
//  GET COMPANY
//  https://developer.aircall.io/api-references/#retrieve-company-information
//  ***********************************************************************************

// Get a company. Reference: https://developer.aircall.io/api-references/#retrieve-company-information
func (service *CompaniesService) Get() (*CompanyResponse, *Response, error) {
	_url := "company"

	responseBody := new(CompanyResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
