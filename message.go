package aircall

import "fmt"

// Message service
type MessagesService service

type NumberConfiguration struct {
	Token       string `json:"token,omitempty"`
	CallbackURL string `json:"callbackUrl,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Message struct {
	ID             string         `json:"id,omitempty"`
	DirectLink     string         `json:"direct_link,omitempty"`
	Direction      string         `json:"direction,omitempty"`
	ExternalNumber string         `json:"external_number,omitempty"`
	Body           string         `json:"body,omitempty"`
	Status         string         `json:"status,omitempty"`
	RawDigits      string         `json:"raw_digits,omitempty"`
	CreatedAt      interface{}    `json:"created_at,omitempty"`
	UpdatedAt      interface{}    `json:"updated_at,omitempty"`
	SentAt         interface{}    `json:"sent_at,omitempty"`
	MediaURL       *[]string      `json:"media_url,omitempty"`
	User           *User          `json:"user,omitempty"`
	Number         *Number        `json:"number,omitempty"`
	Contact        *Contact       `json:"contact,omitempty"`
	MediaDetails   *[]MediaDetail `json:"media_details,omitempty"`
}

type NewMessage struct {
	To       string   `json:"to,omitempty"`
	Body     string   `json:"body,omitempty"`
	MediaURL []string `json:"media_url,omitempty"`
}

type MediaDetail struct {
	FileName     string `json:"file_name,omitempty"`
	FileType     string `json:"file_type,omitempty"`
	PresignedURL string `json:"presigned_url,omitempty"`
}

//  ***********************************************************************************
//  CREATE NUMBER CONFIGURATION
//  https://developer.aircall.io/api-references/#create-number-configuration
//  ***********************************************************************************

// Create number configuration. Reference: https://developer.aircall.io/api-references/#create-number-configuration
func (service *MessagesService) CreateNumberConfiguration(numberID int, numberConfiguration *NumberConfiguration) (*NumberConfiguration, *Response, error) {
	_url := fmt.Sprintf("numbers/%d/messages/configuration", numberID)

	responseBody := new(NumberConfiguration)
	response, err := service.client.Post(_url, numberConfiguration, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET NUMBER CONFIGURATION
//  https://developer.aircall.io/api-references/#fetch-number-configuration
//  ***********************************************************************************

// Get number configuration. Reference: https://developer.aircall.io/api-references/#fetch-number-configuration
func (service *MessagesService) GetNumberConfiguration(numberID int) (*NumberConfiguration, *Response, error) {
	_url := fmt.Sprintf("numbers/%d/messages/configuration", numberID)

	responseBody := new(NumberConfiguration)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE NUMBER CONFIGURATION
//  https://developer.aircall.io/api-references/#fetch-number-configuration
//  ***********************************************************************************

// Delete number configuration. Reference: https://developer.aircall.io/api-references/#fetch-number-configuration
func (service *MessagesService) DeleteNumberConfiguration(numberID int) (*Response, error) {
	_url := fmt.Sprintf("numbers/%d/messages/configuration", numberID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  SEND MESSAGE
//  https://developer.aircall.io/api-references/#send-message
//  ***********************************************************************************

// Send message. Reference: https://developer.aircall.io/api-references/#send-message
func (service *MessagesService) Send(numberID int, message *NewMessage) (*Message, *Response, error) {
	_url := fmt.Sprintf("numbers/%d/messages/send", numberID)

	responseBody := new(Message)
	response, err := service.client.Post(_url, message, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
