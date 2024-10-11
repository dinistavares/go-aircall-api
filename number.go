package aircall

import "fmt"

// Number service
type NumbersService service

type NumbersResponse struct {
	Meta    *GenericResponseMeta `json:"meta,omitempty"`
	Numbers *[]Number            `json:"numbers,omitempty"`
}

type NumberResponse struct {
	Number *Number `json:"number,omitempty"`
}

type Number struct {
	ID                     int       `json:"id,omitempty"`
	DirectLink             string    `json:"direct_link,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Digits                 string    `json:"digits,omitempty"`
	CreatedAt              string    `json:"created_at,omitempty"`
	Country                string    `json:"country,omitempty"`
	TimeZone               string    `json:"time_zone,omitempty"`
	Open                   bool      `json:"open,omitempty"`
	AvailabilityStatus     string    `json:"availability_status,omitempty"`
	IsIVR                  bool      `json:"is_ivr,omitempty"`
	LiveRecordingActivated bool      `json:"live_recording_activated,omitempty"`
	Priority               *int      `json:"priority,omitempty"`
	Messages               *Messages `json:"messages,omitempty"`
	Users                  *User     `json:"users,omitempty"`
}

type Messages struct {
	Welcome        string `json:"welcome,omitempty"`
	Waiting        string `json:"waiting,omitempty"`
	RingingTone    string `json:"ringing_tone,omitempty"`
	UnansweredCall string `json:"unanswered_call,omitempty"`
	AfterHours     string `json:"after_hours,omitempty"`
	IVR            string `json:"ivr,omitempty"`
	Voicemail      string `json:"voicemail,omitempty"`
	Closed         string `json:"closed,omitempty"`
	CallbackLater  string `json:"callback_later,omitempty"`
}

type NumberQueries struct{}

// Query parameters for 'ListNumbers' method.
type ListNumbersQueryParams struct {
	QueryValues
}

//  ***********************************************************************************
//  LIST ALL NUMBERS
//  https://developer.aircall.io/api-references/#list-all-numbers
//  ***********************************************************************************

// Creates Query parameters for 'ListNumbers'
func (pq NumberQueries) NewListNumbers() *ListNumbersQueryParams {
	return &ListNumbersQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set from for for 'List' method.
func (p ListNumbersQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'List' method.
func (p ListNumbersQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'List' method.
func (p ListNumbersQueryParams) Order(value string) {
	p.order(value)
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListNumbersQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List calls. Reference: https://developer.aircall.io/api-references/#list-all-calls
func (service *NumbersService) List(opts *ListNumbersQueryParams) (*NumbersResponse, *Response, error) {
	_url := "numbers"

	responseBody := new(NumbersResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET NUMBER
//  https://developer.aircall.io/api-references/#retrieve-a-number
//  ***********************************************************************************

// Get a number. Reference: https://developer.aircall.io/api-references/#retrieve-a-number
func (service *NumbersService) Get(numberID int) (*NumberResponse, *Response, error) {
	_url := fmt.Sprintf("numbers/%d", numberID)

	responseBody := new(NumberResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE NUMBER
//  https://developer.aircall.io/api-references/#update-a-number
//  ***********************************************************************************

// Update a number. Reference: https://developer.aircall.io/api-references/#update-a-number
func (service *NumbersService) Update(numberID int, number *Number) (*NumberResponse, *Response, error) {
	_url := fmt.Sprintf("numbers/%d", numberID)

	responseBody := new(NumberResponse)
	response, err := service.client.Put(_url, number, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE MUSIC AND MESSAGES
//  https://developer.aircall.io/api-references/#update-music-and-messages
//  ***********************************************************************************

// Update music and messages. Reference: https://developer.aircall.io/api-references/#update-music-and-messages
func (service *MessagesService) UpdateMessages(numberID int, messages *Messages) (*NumberResponse, *Response, error) {
	_url := fmt.Sprintf("numbers/%d", numberID)

	responseBody := new(NumberResponse)
	response, err := service.client.Put(_url, messages, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
