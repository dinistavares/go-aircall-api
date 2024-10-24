package aircall

import "fmt"

// Dialer Campaign service
type DialerCampaignsService service

type DialerCampaign struct {
	ID        int         `json:"id,omitempty"`
	NumberID  string      `json:"number_id,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
}

type CreateUpdateDialerCampaign struct {
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
}

type DialerCampaignPhoneNumberResponse struct {
	PhoneNumbers *[]DialerCampaignPhoneNumber `json:"numbers,omitempty"`
}

type DialerCampaignPhoneNumber struct {
	ID        int         `json:"id,omitempty"`
	Number    string      `json:"number,omitempty"`
	Called    bool        `json:"called,omitempty"`
	CreatedAt interface{} `json:"created_at,omitempty"`
}

//  ***********************************************************************************
//  GET DIALER CAMPAIGN
//  https://developer.aircall.io/api-references/#retrieve-a-dialer-campaign
//  ***********************************************************************************

// Get a dialer campaign. Reference: https://developer.aircall.io/api-references/#retrieve-a-dialer-campaign
func (service *DialerCampaignsService) Get(userID int) (*DialerCampaign, *Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign", userID)

	responseBody := new(DialerCampaign)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE DIALER CAMPAIGN
//  https://developer.aircall.io/api-references/#create-a-dialer-campaign
//  ***********************************************************************************

// Create a dialer campaign. Reference: https://developer.aircall.io/api-references/#create-a-dialer-campaign
func (service *DialerCampaignsService) Create(userID int, dialerCampaign *CreateUpdateDialerCampaign) (*Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign", userID)

	return service.client.Post(_url, dialerCampaign, nil)
}

//  ***********************************************************************************
//  DELETE DIALER CAMPAIGN
//  https://developer.aircall.io/api-references/#delete-a-dialer-campaign
//  ***********************************************************************************

// Delete a dialer campaign. Reference: https://developer.aircall.io/api-references/#delete-a-dialer-campaign
func (service *DialerCampaignsService) Delete(userId int) (*Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign", userId)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  LIST DIALER CAMPAIGN PHONE NUMBERS
//  https://developer.aircall.io/api-references/#retrieve-phone-numbers
//  ***********************************************************************************

// List dialer campaign phone numbers. Reference: https://developer.aircall.io/api-references/#retrieve-phone-numbers
func (service *DialerCampaignsService) ListNumbers(userId int) (*DialerCampaignPhoneNumberResponse, *Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign/phone_numbers", userId)

	responseBody := new(DialerCampaignPhoneNumberResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  ADD DIALER CAMPAIGN PHONE NUMBERS
//  https://developer.aircall.io/api-references/#add-phone-numbers
//  ***********************************************************************************

// Add dialer campaign phone numbers. Reference: https://developer.aircall.io/api-references/#add-phone-numbers
func (service *DialerCampaignsService) AddNumbers(userId int, phoneNumbers *CreateUpdateDialerCampaign) (*Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign/phone_numbers", userId)

	return service.client.Post(_url, phoneNumbers, nil)
}

//  ***********************************************************************************
//  DELETE DIALER CAMPAIGN PHONE NUMBER
//  https://developer.aircall.io/api-references/#delete-a-phone-number
//  ***********************************************************************************

// Delete dialer campaign phone number. Reference: https://developer.aircall.io/api-references/#delete-a-phone-number
func (service *DialerCampaignsService) DeleteNumber(userId int, phoneNumberId int) (*Response, error) {
	_url := fmt.Sprintf("users/%d/dialer_campaign/phone_numbers/%d", userId, phoneNumberId)

	return service.client.Delete(_url)
}
