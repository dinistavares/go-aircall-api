package aircall

import (
	"fmt"
)

// Call service
type CallsService service

type CallsResponse struct {
	Meta  *GenericResponseMeta `json:"meta,omitempty"`
	Calls *[]Call              `json:"calls,omitempty"`
}

type CallResponse struct {
	Call *Call `json:"call,omitempty"`
}

type Call struct {
	ID                 int                `json:"id,omitempty"`
	Sid                string             `json:"sid,omitempty"`
	DirectLink         string             `json:"direct_link,omitempty"`
	Direction          string             `json:"direction,omitempty"`
	Status             string             `json:"status,omitempty"`
	MissedCallReason   string             `json:"missed_call_reason,omitempty"`
	StartedAt          int                `json:"started_at,omitempty"`
	AnsweredAt         int                `json:"answered_at,omitempty"`
	EndedAt            int                `json:"ended_at,omitempty"`
	Duration           int                `json:"duration,omitempty"`
	Voicemail          string             `json:"voicemail,omitempty"`
	Recording          string             `json:"recording,omitempty"`
	Asset              string             `json:"asset,omitempty"`
	RawDigits          string             `json:"raw_digits,omitempty"`
	Archived           bool               `json:"archived,omitempty"`
	Cost               string             `json:"cost,omitempty"`
	Contact            *Contact           `json:"contact,omitempty"`
	Number             *Number            `json:"number,omitempty"`
	User               *User              `json:"user,omitempty"`
	TransferredBy      *User              `json:"transferred_by,omitempty"`
	TransferredTo      *User              `json:"transferred_to,omitempty"`
	AssignedTo         *User              `json:"assigned_to,omitempty"`
	IvrOptionsSelected *CallIVROption     `json:"ivr_options_selected,omitempty"`
	Comments           *[]CallComment     `json:"comments,omitempty"`
	Tags               *[]CallTag         `json:"tags,omitempty"`
	Participants       *[]CallParticipant `json:"participants,omitempty"`
	Teams              *[]Team            `json:"teams,omitempty"`
}

type CallComment struct {
	ID       int           `json:"id,omitempty"`
	Content  string        `json:"content,omitempty"`
	PostedAt int           `json:"posted_at,omitempty"`
	PostedBy *CallActionBy `json:"posted_by,omitempty"`
}

type CallActionBy struct {
	ID                 int         `json:"id,omitempty"`
	DirectLink         string      `json:"direct_link,omitempty"`
	Name               string      `json:"name,omitempty"`
	Email              string      `json:"email,omitempty"`
	Available          bool        `json:"available,omitempty"`
	AvailabilityStatus string      `json:"availability_status,omitempty"`
	CreatedAt          interface{} `json:"created_at,omitempty"`
	TimeZone           string      `json:"time_zone,omitempty"`
	Language           string      `json:"language,omitempty"`
	State              string      `json:"state,omitempty"`

	// Included in Webhook event only
	Substatus string `json:"substatus,omitempty"`
}

type CallTag struct {
	ID        int           `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	CreatedAt interface{}   `json:"created_at,omitempty"`
	TaggedBy  *CallActionBy `json:"tagged_by,omitempty"`

	// Included in Webhook event only
	TaggedAt int `json:"tagged_at,omitempty"`
}

type CallInsightCard struct {
	Contents *[]CallInsightCardContents `json:"contents,omitempty"`
}

type CallInsightCardContents struct {
	Type   string `json:"type,omitempty"`
	Text   string `json:"text,omitempty"`
	Link   string `json:"link,omitempty"`
	Label  string `json:"label,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

type CallTransfer struct {
	UserID              string `json:"user_id,omitempty"`
	TeamID              string `json:"team_id,omitempty"`
	Number              string `json:"number,omitempty"`
	DispatchingStrategy string `json:"dispatching_strategy,omitempty"`
}

type CallTags struct {
	Tags []int `json:"tags,omitempty"`
}

type CallParticipant struct {
	ID          *string `json:"id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Name        *string `json:"name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

type CallIVROption struct {
	ID                  string      `json:"id,omitempty"`
	Title               string      `json:"title,omitempty"`
	Key                 string      `json:"key,omitempty"`
	Branch              string      `json:"branch,omitempty"`
	CreatedAt           interface{} `json:"created_at,omitempty"`
	TransitionStartedAt string      `json:"transition_started_at,omitempty"`
	TransitionEndedAt   string      `json:"transition_ended_at,omitempty"`
}

type CallQueries struct{}

// Query parameters for 'List' method.
type ListCallsQueryParams struct {
	QueryValues
}

// Query parameters for 'SearchCalls' method.
type SearchCallsQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *CallsService) Query() *CallQueries {
	return &CallQueries{}
}

//  ***********************************************************************************
//  LIST ALL CALLS
//  https://developer.aircall.io/api-references/#list-all-calls
//  ***********************************************************************************

// Creates Query parameters for 'List'
func (pq CallQueries) NewListCalls() *ListCallsQueryParams {
	return &ListCallsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set from for for 'List' method.
func (p ListCallsQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'List' method.
func (p ListCallsQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'List' method.
func (p ListCallsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'fetch_contact' for for 'List' method.
func (p ListCallsQueryParams) FetchContact(value bool) {
	p.set("fetch_contact", fmt.Sprintf("%t", value))
}

// Set 'fetch_short_urls' for for 'List' method.
func (p ListCallsQueryParams) FetchShortURLs(value bool) {
	p.set("fetch_short_urls", fmt.Sprintf("%t", value))
}

// Set 'fetch_call_timeline' for for 'List' method.
func (p ListCallsQueryParams) FetchCallTimeline(value bool) {
	p.set("fetch_call_timeline", fmt.Sprintf("%t", value))
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListCallsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List calls. Reference: https://developer.aircall.io/api-references/#list-all-calls
func (service *CallsService) List(opts *ListCallsQueryParams) (*CallsResponse, *Response, error) {
	_url := "calls"

	responseBody := new(CallsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  SEARCH CALLS
//  https://developer.aircall.io/api-references/#search-calls
//  ***********************************************************************************

// Creates Query parameters for 'SearchCalls'
func (pq CallQueries) NewSearchCalls() *SearchCallsQueryParams {
	return &SearchCallsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set from for for 'Search' method.
func (p SearchCallsQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'Search' method.
func (p SearchCallsQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'Search' method.
func (p SearchCallsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'direction' for for 'Search' method.
func (p SearchCallsQueryParams) Direction(value string) {
	p.set("direction", value)
}

// Set 'phone_number' for for 'Search' method.
func (p SearchCallsQueryParams) PhoneNumber(value string) {
	p.set("phone_number", value)
}

// Set 'user_id' for for 'Search' method.
func (p SearchCallsQueryParams) UserID(value int) {
	p.set("user_id", fmt.Sprintf("%d", value))
}

// Set 'fetch_contact' for for 'SearchCalls' method.
func (p SearchCallsQueryParams) FetchContact(value bool) {
	p.set("fetch_contact", fmt.Sprintf("%t", value))
}

// Set 'fetch_short_urls' for for 'SearchCalls' method.
func (p SearchCallsQueryParams) FetchShortURLs(value bool) {
	p.set("fetch_short_urls", fmt.Sprintf("%t", value))
}

// Set 'fetch_call_timeline' for for 'SearchCalls' method.
func (p SearchCallsQueryParams) FetchCallTimeline(value bool) {
	p.set("fetch_call_timeline", fmt.Sprintf("%t", value))
}

// Set 'page' and 'per_page' for for 'Search' method.
func (p SearchCallsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// Search calls. Reference: https://developer.aircall.io/api-references/#search-calls
func (service *CallsService) Search(opts *SearchCallsQueryParams) (*CallsResponse, *Response, error) {
	_url := "calls/search"

	responseBody := new(CallsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET CALL
//  https://developer.aircall.io/api-references/#retrieve-a-call
//  ***********************************************************************************

// Get a call. Reference: https://developer.aircall.io/api-references/#retrieve-a-call
func (service *CallsService) Get(callID int) (*CallResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d", callID)

	responseBody := new(CallResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  TRANSFER CALL
//  https://developer.aircall.io/api-references/#transfer-a-call
//  ***********************************************************************************

// Transfer a call. Reference: https://developer.aircall.io/api-references/#transfer-a-call
func (service *CallsService) Transfer(callID int, transferCall *CallTransfer) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/transfers", callID)

	return service.client.Post(_url, transferCall, nil)
}

//  ***********************************************************************************
//  COMMENT CALL
//  https://developer.aircall.io/api-references/#comment-a-call
//  ***********************************************************************************

// Comment a call. Reference: https://developer.aircall.io/api-references/#comment-a-call
func (service *CallsService) Comment(callID int, comment string) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/comments", callID)

	commentCall := CallComment{
		Content: comment,
	}

	return service.client.Post(_url, commentCall, nil)
}

//  ***********************************************************************************
//  TAG CALL
//  https://developer.aircall.io/api-references/#tag-a-call
//  ***********************************************************************************

// Tag a call. Reference: https://developer.aircall.io/api-references/#tag-a-call
func (service *CallsService) Tag(callID int, tags []int) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/tags", callID)

	tagsCall := CallTags{
		Tags: tags,
	}

	return service.client.Post(_url, tagsCall, nil)
}

//  ***********************************************************************************
//  ARCHIVE CALL
//  https://developer.aircall.io/api-references/#archive-a-call
//  ***********************************************************************************

// Archive a call. Reference: https://developer.aircall.io/api-references/#archive-a-call
func (service *CallsService) Archive(callID int) (*CallResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/archive", callID)

	responseBody := new(CallResponse)
	response, err := service.client.Put(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UNARCHIVE CALL
//  https://developer.aircall.io/api-references/#unarchive-a-call
//  ***********************************************************************************

// Unarchive a call. Reference: https://developer.aircall.io/api-references/#unarchive-a-call
func (service *CallsService) Unarchive(callID int) (*CallResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/unarchive", callID)

	responseBody := new(CallResponse)
	response, err := service.client.Put(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  PAUSE RECORDING CALL
//  https://developer.aircall.io/api-references/#pause-recording-on-a-call
//  ***********************************************************************************

// Pause recording on a call. Reference: https://developer.aircall.io/api-references/#pause-recording-on-a-call
func (service *CallsService) PauseRecording(callID int) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/pause_recording", callID)

	return service.client.Post(_url, nil, nil)
}

//  ***********************************************************************************
//  RESUME RECORDING CALL
//  https://developer.aircall.io/api-references/#resume-recording-on-a-call
//  ***********************************************************************************

// Resume recording on a call. Reference: https://developer.aircall.io/api-references/#resume-recording-on-a-call
func (service *CallsService) ResumeRecording(callID int) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/resume_recording", callID)

	return service.client.Post(_url, nil, nil)
}

//  ***********************************************************************************
//  DELETE RECORDING CALL
//  https://developer.aircall.io/api-references/#delete-call-recording
//  ***********************************************************************************

// Delete call recording. Reference: https://developer.aircall.io/api-references/#delete-call-recording
func (service *CallsService) DeleteRecording(callID int) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/recording", callID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  DELETE VOICEMAIL CALL
//  https://developer.aircall.io/api-references/#delete-call-voicemail
//  ***********************************************************************************

// Delete call voicemail. Reference: https://developer.aircall.io/api-references/#delete-call-voicemail
func (service *CallsService) DeleteVoicemail(callID int) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/voicemail", callID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  ADD INSIGHT CARD TO CALL
//  https://developer.aircall.io/api-references/#insight-cards
//  ***********************************************************************************

// Add insight card to call. Reference: https://developer.aircall.io/api-references/#insight-cards
func (service *CallsService) AddInsightCard(callID int, insightCard *CallInsightCard) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/insight_cards", callID)

	return service.client.Post(_url, insightCard, nil)
}
