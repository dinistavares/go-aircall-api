package aircall

import (
	"fmt"
)

// Webhook service
type WebhookService service

type WebhooksResponse struct {
	Meta     *GenericResponseMeta `json:"meta,omitempty"`
	Webhooks *[]Webhook           `json:"webhooks,omitempty"`
}

type WebhookResponse struct {
	Webhook *Webhook `json:"webhook,omitempty"`
}

type CreateUpdateWebhook struct {
	CustomName string   `json:"custom_name,omitempty"`
	URL        string   `json:"url,omitempty"`
	Events     []string `json:"events,omitempty"`
}

type Webhook struct {
	WebhookID  string      `json:"webhook_id,omitempty"`
	DirectLink string      `json:"direct_link,omitempty"`
	CreatedAt  interface{} `json:"created_at,omitempty"`
	URL        string      `json:"url,omitempty"`
	Active     bool        `json:"active,omitempty"`
	Token      string      `json:"token,omitempty"`
	Events     *[]string   `json:"events,omitempty"`
}

type WebhookQueries struct{}

// Query parameters for 'ListWebhooks' method.
type ListWebhooksQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *WebhookService) Query() *WebhookQueries {
	return &WebhookQueries{}
}

//  ***********************************************************************************
//  LIST ALL WEBHOOKS
//  https://developer.aircall.io/api-references/#list-all-webhooks
//  ***********************************************************************************

// Creates Query parameters for 'ListContacts'
func (pq WebhookQueries) NewListWebhooks() *ListWebhooksQueryParams {
	return &ListWebhooksQueryParams{
		QueryValues: QueryValues{},
	}
}

func (p ListWebhooksQueryParams) Order(value string) {
	p.order(value)
}

func (p ListWebhooksQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List webhooks. Reference: https://developer.aircall.io/api-references/#list-all-webhooks
func (service *WebhookService) List(opts *ListWebhooksQueryParams) (*WebhooksResponse, *Response, error) {
	_url := "webhooks"

	responseBody := new(WebhooksResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET WEBHOOK
//  https://developer.aircall.io/api-references/#retrieve-a-webhook
//  ***********************************************************************************

// Get a webhook. Reference: https://developer.aircall.io/api-references/#retrieve-a-webhook
func (service *WebhookService) Get(webhookID string) (*WebhookResponse, *Response, error) {
	_url := fmt.Sprintf("webhooks/%s", webhookID)

	responseBody := new(WebhookResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE WEBHOOK
//  https://developer.aircall.io/api-references/#create-a-webhook
//  ***********************************************************************************

// Create a webhook. Reference: https://developer.aircall.io/api-references/#create-a-webhook
func (service *WebhookService) Create(webhook *CreateUpdateWebhook) (*WebhookResponse, *Response, error) {
	_url := "webhooks"

	responseBody := new(WebhookResponse)
	response, err := service.client.Post(_url, webhook, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE WEBHOOK
//  https://developer.aircall.io/api-references/#update-a-webhook
//  ***********************************************************************************

// Update a webhook. Reference: https://developer.aircall.io/api-references/#update-a-webhook
func (service *WebhookService) Update(webhookID string, webhook *CreateUpdateWebhook) (*WebhookResponse, *Response, error) {
	_url := fmt.Sprintf("webhooks/%s", webhookID)

	responseBody := new(WebhookResponse)
	response, err := service.client.Put(_url, webhook, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE WEBHOOK
//  https://developer.aircall.io/api-references/#delete-a-webhook
//  ***********************************************************************************

// Delete a webhook. Reference: https://developer.aircall.io/api-references/#delete-a-webhook
func (service *WebhookService) Delete(webhookID string) (*Response, error) {
	_url := fmt.Sprintf("webhooks/%s", webhookID)

	response, err := service.client.Delete(_url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
