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
	ID               int            `json:"id,omitempty"`
	Sid              string         `json:"sid,omitempty"`
	DirectLink       string         `json:"direct_link,omitempty"`
	Direction        string         `json:"direction,omitempty"`
	Status           string         `json:"status,omitempty"`
	MissedCallReason string         `json:"missed_call_reason,omitempty"`
	StartedAt        int            `json:"started_at,omitempty"`
	AnsweredAt       int            `json:"answered_at,omitempty"`
	EndedAt          int            `json:"ended_at,omitempty"`
	Duration         int            `json:"duration,omitempty"`
	Voicemail        string         `json:"voicemail,omitempty"`
	Recording        string         `json:"recording,omitempty"`
	Asset            string         `json:"asset,omitempty"`
	RawDigits        string         `json:"raw_digits,omitempty"`
	Contact          *Contact       `json:"contact,omitempty"`
	Archived         bool           `json:"archived,omitempty"`
	Cost             string         `json:"cost,omitempty"`
	Comments         *[]CallComment `json:"comments,omitempty"`
	Tags             *[]CallTag     `json:"tags,omitempty"`

	// TODO: add types when services are implemented
	TransferredBy      interface{} `json:"transferred_by,omitempty"`
	TransferredTo      interface{} `json:"transferred_to,omitempty"`
	AssignedTo         interface{} `json:"assigned_to,omitempty"`
	User               interface{} `json:"user,omitempty"`
	Number             interface{} `json:"number,omitempty"`
	Teams              interface{} `json:"teams,omitempty"`
	Participants       interface{} `json:"participants,omitempty"`
	IvrOptionsSelected interface{} `json:"ivr_options_selected,omitempty"`
}

type CallComment struct {
	ID       int              `json:"id,omitempty"`
	Content  string           `json:"content,omitempty"`
	PostedAt int              `json:"posted_at,omitempty"`
	PostedBy *CommentPostedBy `json:"posted_by,omitempty"`
}

type CommentPostedBy struct {
	ID                 int    `json:"id,omitempty"`
	DirectLink         string `json:"direct_link,omitempty"`
	Name               string `json:"name,omitempty"`
	Email              string `json:"email,omitempty"`
	Available          bool   `json:"available,omitempty"`
	AvailabilityStatus string `json:"availability_status,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
}

type CallTag struct {
	ID        int           `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	CreatedAt int           `json:"created_at,omitempty"`
	TaggedBy  *CallTaggedBy `json:"tagged_by,omitempty"`
}

type CallTaggedBy struct {
	ID                 int    `json:"id,omitempty"`
	DirectLink         string `json:"direct_link,omitempty"`
	Name               string `json:"name,omitempty"`
	Email              string `json:"email,omitempty"`
	Available          bool   `json:"available,omitempty"`
	AvailabilityStatus string `json:"availability_status,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
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

type CallQueries struct{}

// Query parameters for 'ListCalls' method.
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

// Creates Query parameters for 'ListCalls'
func (pq CallQueries) NewListCalls() *ListCallsQueryParams {
	return &ListCallsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set from for for 'ListCalls' method.
func (p ListCallsQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'ListCalls' method.
func (p ListCallsQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'ListCalls' method.
func (p ListCallsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'fetch_contact' for for 'ListCalls' method.
func (p ListCallsQueryParams) FetchContact(value bool) {
	p.set("fetch_contact", fmt.Sprintf("%t", value))
}

// Set 'fetch_short_urls' for for 'ListCalls' method.
func (p ListCallsQueryParams) FetchShortURLs(value bool) {
	p.set("fetch_short_urls", fmt.Sprintf("%t", value))
}

// Set 'fetch_call_timeline' for for 'ListCalls' method.
func (p ListCallsQueryParams) FetchCallTimeline(value bool) {
	p.set("fetch_call_timeline", fmt.Sprintf("%t", value))
}

// Set 'page' and 'per_page' for for 'ListCalls' method.
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
//  ADD INSIGHT CARD TO CALL
//  https://developer.aircall.io/api-references/#insight-cards
//  ***********************************************************************************

// Add insight card to call. Reference: https://developer.aircall.io/api-references/#insight-cards
func (service *CallsService) AddInsightCard(callID int, insightCard *CallInsightCard) (*Response, error) {
	_url := fmt.Sprintf("calls/%d/insight_cards", callID)

	return service.client.Post(_url, insightCard, nil)
}
