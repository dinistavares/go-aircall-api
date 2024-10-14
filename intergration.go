package aircall

// Integration service
type IntegrationService service

type IntegrationResponse struct {
	Integration *Integration `json:"integration,omitempty"`
}

type Integration struct {
	Name         string `json:"name,omitempty"`
	CustomName   string `json:"custom_name,omitempty"`
	Logo         string `json:"logo,omitempty"`
	CompanyID    int    `json:"company_id,omitempty"`
	Status       string `json:"status,omitempty"`
	Active       bool   `json:"active,omitempty"`
	NumbersCount int    `json:"numbers_count,omitempty"`
	NumberIDs    []int  `json:"number_ids,omitempty"`
	User         *User  `json:"user,omitempty"`

	// Included in Webhook event only
	IntegrationID int `json:"integration_id,omitempty"`
}

//  ***********************************************************************************
//  GET INTEGRATION
//  https://developer.aircall.io/api-references/#retrieve-a-integration
//  ***********************************************************************************

// Get an integration. Reference: https://developer.aircall.io/api-references/#retrieve-a-integration
func (service *IntegrationService) Get() (*IntegrationResponse, *Response, error) {
	_url := "integrations/me"

	responseBody := new(IntegrationResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  ENABLE INTEGRATION
//  https://developer.aircall.io/api-references/#enable-integration
//  ***********************************************************************************

// Enable an integration. Reference: https://developer.aircall.io/api-references/#enable-integration
func (service *IntegrationService) Enable(install ...bool) (*Response, error) {
	_url := "integrations/enable"

	if len(install) > 0 && install[0] {
		_url += "?install=true"
	}

	return service.client.Get(_url, nil, nil)
}

//  ***********************************************************************************
//  DISABLE INTEGRATION
//  https://developer.aircall.io/api-references/#disable-integration
//  ***********************************************************************************

// Disable an integration. Reference: https://developer.aircall.io/api-references/#disable-integration
func (service *IntegrationService) Disable() (*Response, error) {
	_url := "integrations/disable"

	return service.client.Delete(_url)
}
