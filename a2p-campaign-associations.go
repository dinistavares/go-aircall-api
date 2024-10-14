package aircall

import "fmt"

// A2P Campaign Associations service
type A2PCampaignAssociationsService service

type A2PCampaignAssociationsResponse struct {
	Meta     *GenericResponseMeta      `json:"meta,omitempty"`
	Contacts *[]A2PCampaignAssociation `json:"a2p_campaign_associations,omitempty"`
}

type A2PCampaignAssociationResponse struct {
	Contact *A2PCampaignAssociation `json:"a2p_campaign_associations,omitempty"`
}

type A2PCampaignAssociation struct {
	ID            int    `json:"id,omitempty"`
	CompanyID     int    `json:"company_id,omitempty"`
	ExternalID    string `json:"external_id,omitempty"`
	UpdateStatus  string `json:"update_status,omitempty"`
	UpdateMessage string `json:"update_message,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	Numbers       *[]int `json:"numbers,omitempty"`
}

type CreateUpdateA2PCampaignAssociation struct {
	ExternalCampaignID string `json:"external_campaign_id,omitempty"`
	Numbers            *[]int `json:"numbers,omitempty"`
}

type A2PCampaignAssociationsQueries struct{}

// Query parameters for 'ListA2PCampaignAssociations' method.
type ListA2PCampaignAssociationsQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *A2PCampaignAssociationsService) Query() *A2PCampaignAssociationsQueries {
	return &A2PCampaignAssociationsQueries{}
}

//  ***********************************************************************************
//  LIST ALL A2P CAMPAIGN ASSOCIATIONS
//  https://developer.aircall.io/api-references/#list-a2p-campaign-associations
//  ***********************************************************************************

// Creates Query parameters for 'ListContacts'
func (pq A2PCampaignAssociationsQueries) NewListA2PCampaignAssociations() *ListA2PCampaignAssociationsQueryParams {
	return &ListA2PCampaignAssociationsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListA2PCampaignAssociationsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// Get all A2P campaign associations. Reference: https://developer.aircall.io/api-references/#list-a2p-campaign-associations
func (service *A2PCampaignAssociationsService) List(queryParams *ListA2PCampaignAssociationsQueryParams) (*A2PCampaignAssociationsResponse, *Response, error) {
	_url := "a2p_campaign_associations"

	responseBody := new(A2PCampaignAssociationsResponse)
	response, err := service.client.Get(_url, queryParams, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE A2P CAMPAIGN ASSOCIATION
//  https://developer.aircall.io/api-references/#create-an-a2p-campaign-association
//  ***********************************************************************************

// Create A2P campaign association. Reference: https://developer.aircall.io/api-references/#create-an-a2p-campaign-association
func (service *A2PCampaignAssociationsService) Create(a2pCampaignAssociation *CreateUpdateA2PCampaignAssociation) (*A2PCampaignAssociationResponse, *Response, error) {
	_url := "a2p_campaign_associations"

	responseBody := new(A2PCampaignAssociationResponse)
	response, err := service.client.Post(_url, a2pCampaignAssociation, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE A2P CAMPAIGN ASSOCIATION
//  https://developer.aircall.io/api-references/#update-an-a2p-campaign-association
//  ***********************************************************************************

// Update A2P campaign association. Reference: https://developer.aircall.io/api-references/#update-an-a2p-campaign-association
func (service *A2PCampaignAssociationsService) Update(a2pCampaignAssociation *CreateUpdateA2PCampaignAssociation) (*A2PCampaignAssociationResponse, *Response, error) {
	_url := "a2p_campaign_associations"

	responseBody := new(A2PCampaignAssociationResponse)
	response, err := service.client.Put(_url, a2pCampaignAssociation, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE A2P CAMPAIGN ASSOCIATION
//  https://developer.aircall.io/api-references/#delete-an-a2p-campaign-association
//  ***********************************************************************************

// Delete a contact. Reference: https://developer.aircall.io/api-references/#delete-an-a2p-campaign-association
func (service *A2PCampaignAssociationsService) Delete(a2pCampaignAssociationID int) (*Response, error) {
	_url := fmt.Sprintf("a2p_campaign_associations/%d", a2pCampaignAssociationID)

	return service.client.Delete(_url)
}
