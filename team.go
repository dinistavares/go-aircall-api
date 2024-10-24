package aircall

import "fmt"

// Team service
type TeamsService service

type TeamsResponse struct {
	Meta  *GenericResponseMeta `json:"meta,omitempty"`
	Teams *[]Team              `json:"teams,omitempty"`
}

type TeamResponse struct {
	Team *Team `json:"team,omitempty"`
}

type Team struct {
	ID         int         `json:"id,omitempty"`
	DirectLink string      `json:"direct_link,omitempty"`
	Name       string      `json:"name,omitempty"`
	CreatedAt  interface{} `json:"created_at,omitempty"`
	Users      *[]User     `json:"users,omitempty"`
}

type CreateTeam struct {
	Name string `json:"name,omitempty"`
}

type TeamQueries struct{}

// Query parameters for 'ListTeams' method.
type ListTeamsQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *TeamsService) Query() *TeamQueries {
	return &TeamQueries{}
}

//  ***********************************************************************************
//  LIST ALL TEAMS
//  https://developer.aircall.io/api-references/#list-all-teams
//  ***********************************************************************************

// Creates Query parameters for 'ListContacts'
func (pq TeamQueries) NewListTeams() *ListTeamsQueryParams {
	return &ListTeamsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'order' for for 'List' method.
func (p ListTeamsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListTeamsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List teams. Reference: https://developer.aircall.io/api-references/#list-all-teams
func (service *TeamsService) List(opts *ListTeamsQueryParams) (*TeamsResponse, *Response, error) {
	_url := "teams"

	responseBody := new(TeamsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET TEAM
//  https://developer.aircall.io/api-references/#retrieve-a-team
//  ***********************************************************************************

// Get a team. Reference: https://developer.aircall.io/api-references/#retrieve-a-team
func (service *TeamsService) Get(teamID int) (*TeamResponse, *Response, error) {
	_url := fmt.Sprintf("teams/%d", teamID)

	responseBody := new(TeamResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE TEAM
//  https://developer.aircall.io/api-references/#create-a-team
//  ***********************************************************************************

// Create a team. Reference: https://developer.aircall.io/api-references/#create-a-team
func (service *TeamsService) Create(team *CreateTeam) (*TeamResponse, *Response, error) {
	_url := "teams"

	responseBody := new(TeamResponse)
	response, err := service.client.Post(_url, team, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE TEAM
//  https://developer.aircall.io/api-references/#delete-a-team
//  ***********************************************************************************

// Delete a team. Reference: https://developer.aircall.io/api-references/#delete-a-team
func (service *TeamsService) Delete(teamID int) (*Response, error) {
	_url := fmt.Sprintf("teams/%d", teamID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  ADD USER TO TEAM
//  https://developer.aircall.io/api-references/#add-a-user-to-a-team
//  ***********************************************************************************

// Add a user to a team. Reference: https://developer.aircall.io/api-references/#add-a-user-to-a-team
func (service *TeamsService) AddUser(teamID int, userID int) (*TeamResponse, *Response, error) {
	_url := fmt.Sprintf("teams/%d/users/%d", teamID, userID)

	responseBody := new(TeamResponse)
	response, err := service.client.Post(_url, userID, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  REMOVE USER FROM TEAM
//  https://developer.aircall.io/api-references/#remove-a-user-from-a-team
//  ***********************************************************************************

// Remove a user from a team. Reference: https://developer.aircall.io/api-references/#remove-a-user-from-a-team
func (service *TeamsService) RemoveUser(teamID int, userID int) (*TeamResponse, *Response, error) {
	_url := fmt.Sprintf("teams/%d/users/%d", teamID, userID)

	responseBody := new(TeamResponse)
	response, err := service.client.Delete(_url, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
