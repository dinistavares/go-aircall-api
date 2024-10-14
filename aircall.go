package aircall

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion               = "1.1"
	defaultAuthHeaderName        = "Authorization"
	defaultOAuthPrefix           = "Bearer"
	defaultAuthPrefix            = "Basic"
	defaultRestEndpointURL       = "https://api.aircall.io/v1/"
	acceptedContentType          = "application/json"
	userAgent                    = "go-aircall-api/" + libraryVersion
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000
)

var (
	errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
	errorDoAttemptNilRequest    = errors.New("request could not be constructed")
)

type ApiType string

type ClientConfig struct {
	HttpClient      *http.Client
	RestEndpointURL string
}

type auth struct {
	Available   bool
	AccessToken string
	HeaderName  string
	Prefix      string
}

type Client struct {
	config  *ClientConfig
	client  *http.Client
	auth    *auth
	baseURL *url.URL

	Call                     *CallsService
	Contact                  *ContactsService
	ConversationIntelligence *ConversationIntelligenceService
	Company                  *CompaniesService
	DialerCampaign           *DialerCampaignsService
	Number                   *NumbersService
	Message                  *MessagesService
	User                     *UsersService
	Tag                      *TagsService
	Team                     *TeamsService
	Webhook                  *WebhookService
}

type service struct {
	client *Client
}

type ErrorResponse struct {
	Response *http.Response

	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
	Troubleshoot string `json:"troubleshoot,omitempty"`
	Success      bool   `json:"success,omitempty"`
}

type Response struct {
	*http.Response
}

func (response *ErrorResponse) Error() string {
	errorString := fmt.Sprintf("%v %v: %d %s",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode, response.ErrorMessage)

	if response.Troubleshoot != "" {
		errorString = fmt.Sprintf("%s [%s]", errorString, response.Troubleshoot)
	}

	if response.Message != "" {
		errorString = fmt.Sprintf("%s [%s]", errorString, response.Message)
	}

	return errorString
}

func NewWithConfig(config ClientConfig) *Client {
	if config.HttpClient == nil {
		config.HttpClient = http.DefaultClient
	}

	if config.RestEndpointURL == "" {
		config.RestEndpointURL = defaultRestEndpointURL
	}

	// Create client
	baseURL, _ := url.Parse(config.RestEndpointURL)

	client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.Call = &CallsService{client: client}
	client.Contact = &ContactsService{client: client}
	client.ConversationIntelligence = &ConversationIntelligenceService{client: client}
	client.Company = &CompaniesService{client: client}
	client.DialerCampaign = &DialerCampaignsService{client: client}
	client.Number = &NumbersService{client: client}
	client.Message = &MessagesService{client: client}
	client.Tag = &TagsService{client: client}
	client.Team = &TeamsService{client: client}
	client.User = &UsersService{client: client}
	client.Webhook = &WebhookService{client: client}

	return client
}

func New() *Client {
	return NewWithConfig(ClientConfig{})
}

func (client *Client) Authenticate(accessToken string) {
	client.auth.HeaderName = defaultAuthHeaderName
	client.auth.AccessToken = accessToken
	client.auth.Available = true
	client.auth.Prefix = defaultOAuthPrefix
}

func (client *Client) AuthenricateBasic(apiID string, apiToken string) {
	client.auth.HeaderName = defaultAuthHeaderName
	client.auth.Available = true
	client.auth.Prefix = defaultAuthPrefix
	client.auth.AccessToken = base64.StdEncoding.EncodeToString([]byte(apiID + ":" + apiToken))
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts interface{}, body interface{}) (*http.Request, error) {
	// Append Query Params to URL
	if opts, ok := isPointerWithQueryValues(opts); ok {
		if v, ok := opts.(QueryValues); ok {
			urlStr += v.getQueryValues().encode()
		}
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	if client.auth.Available {
		req.Header.Add(client.auth.HeaderName, fmt.Sprintf("%s %s", client.auth.Prefix, client.auth.AccessToken))
	}

	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, false, err
}

func newResponse(httpResponse *http.Response) *Response {
	response := Response{Response: httpResponse}

	return &response
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error, or response status is a server error? (HTTP 5xx)
	if err != nil || response.StatusCode >= 500 {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &ErrorResponse{Response: response}

	data, err := io.ReadAll(response.Body)

	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

func (client *Client) Get(url string, opts interface{}, v interface{}) (*Response, error) {
	req, _ := client.NewRequest("GET", url, opts, nil)

	return client.Do(req, v)
}

func (client *Client) Post(url string, body interface{}, v interface{}) (*Response, error) {
	req, _ := client.NewRequest("POST", url, nil, body)

	return client.Do(req, v)
}

func (client *Client) Put(url string, body interface{}, v interface{}) (*Response, error) {
	req, _ := client.NewRequest("PUT", url, nil, body)

	return client.Do(req, v)
}

func (client *Client) Delete(url string, v ...interface{}) (*Response, error) {
	req, _ := client.NewRequest("DELETE", url, nil, nil)

	if len(v) > 0 {
		return client.Do(req, v[0])
	}

	return client.Do(req, nil)
}
