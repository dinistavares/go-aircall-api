package aircall

import (
	"fmt"
)

// Tag service
type TagsService service

type TagsResponse struct {
	Meta *GenericResponseMeta `json:"meta,omitempty"`
	Tags *[]Tag               `json:"tags,omitempty"`
}

type TagResponse struct {
	Tag *Tag `json:"tag,omitempty"`
}

type Tag struct {
	ID          int    `json:"id,omitempty"`
	DirectLink  string `json:"direct_link,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateUpdateTag struct {
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type TagQueries struct{}

// Query parameters for 'List' method.
type ListTagsQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *TagsService) Query() *TagQueries {
	return &TagQueries{}
}

//  ***********************************************************************************
//  LIST ALL TAGS
//  https://developer.aircall.io/api-references/#list-all-tags
//  ***********************************************************************************

// Creates Query parameters for 'List'
func (pq TagQueries) NewListTags() *ListTagsQueryParams {
	return &ListTagsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListTagsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List tags. Reference: https://developer.aircall.io/api-references/#list-all-tags
func (service *TagsService) List(opts *ListTagsQueryParams) (*TagsResponse, *Response, error) {
	_url := "tags"

	responseBody := new(TagsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET TAG
//  https://developer.aircall.io/api-references/#retrieve-a-tag
//  ***********************************************************************************

// Get a tag. Reference: https://developer.aircall.io/api-references/#retrieve-a-tag
func (service *TagsService) Get(tagID int) (*TagResponse, *Response, error) {
	_url := fmt.Sprintf("tags/%d", tagID)

	responseBody := new(TagResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE TAG
//  https://developer.aircall.io/api-references/#create-a-tag
//  ***********************************************************************************

// Create a tag. Reference: https://developer.aircall.io/api-references/#create-a-tag
func (service *TagsService) Create(tag *CreateUpdateTag) (*TagResponse, *Response, error) {
	_url := "tags"

	responseBody := new(TagResponse)
	response, err := service.client.Post(_url, tag, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE TAG
//  https://developer.aircall.io/api-references/#update-a-tag
//  ***********************************************************************************

// Update a tag. Reference: https://developer.aircall.io/api-references/#update-a-tag
func (service *TagsService) Update(tagID int, tag *CreateUpdateTag) (*TagResponse, *Response, error) {
	_url := fmt.Sprintf("tags/%d", tagID)
	

	responseBody := new(TagResponse)
	response, err := service.client.Put(_url, tag, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE TAG
//  https://developer.aircall.io/api-references/#delete-a-tag
//  ***********************************************************************************

// Delete a tag. Reference: https://developer.aircall.io/api-references/#delete-a-tag
func (service *TagsService) Delete(tagID int) (*Response, error) {
	_url := fmt.Sprintf("tags/%d", tagID)

	return service.client.Delete(_url)
}
