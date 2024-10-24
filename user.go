package aircall

import (
	"fmt"
)

// User service
type UsersService service

type UsersResponse struct {
	Meta  *GenericResponseMeta `json:"meta,omitempty"`
	Users *[]User              `json:"users,omitempty"`
}

type UserResponse struct {
	User *User `json:"user,omitempty"`
}

type User struct {
	ID                 int         `json:"id,omitempty"`
	DirectLink         string      `json:"direct_link,omitempty"`
	Name               string      `json:"name,omitempty"`
	Email              string      `json:"email,omitempty"`
	CreatedAt          interface{} `json:"created_at,omitempty"`
	Available          bool        `json:"available,omitempty"`
	AvailabilityStatus string      `json:"availability_status,omitempty"`
	Substatus          string      `json:"substatus,omitempty"`
	TimeZone           string      `json:"time_zone,omitempty"`
	Language           string      `json:"language,omitempty"`
	WrapUpTime         int         `json:"wrap_up_time,omitempty"`
	Numbers            *[]Number   `json:"numbers,omitempty"`
}

type CreateUpdateUser struct {
	Email              string   `json:"email,omitempty"`
	FirstName          string   `json:"first_name,omitempty"`
	LastName           string   `json:"last_name,omitempty"`
	AvailabilityStatus string   `json:"availability_status,omitempty"`
	RoleIDs            []string `json:"role_ids,omitempty"`
	WrapUpTime         int      `json:"wrap_up_time,omitempty"`
}

type UserAvailabilitiesResponse struct {
	Meta  *GenericResponseMeta `json:"meta,omitempty"`
	Users *[]UserAvailability  `json:"users,omitempty"`
}

type UserAvailability struct {
	ID           int    `json:"id,omitempty"`
	Availability string `json:"availability,omitempty"`
}

type NewUserCall struct {
	NumberID int    `json:"number_id,omitempty"`
	To       string `json:"to,omitempty"`
}

type UserQueries struct{}

// Query parameters for 'List' method.
type ListUsersQueryParams struct {
	QueryValues
}

// Query parameters for 'ListAvailability' method.
type ListUsersAvailabilityQueryParams struct {
	QueryValues
}

// Query parameters for 'SearchUsers' method.
type SearchUsersQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *UsersService) Query() *UserQueries {
	return &UserQueries{}
}

//  ***********************************************************************************
//  LIST ALL USERS
//  https://developer.aircall.io/api-references/#list-all-users
//  ***********************************************************************************

// Creates Query parameters for 'List'
func (pq UserQueries) NewListUsers() *ListUsersQueryParams {
	return &ListUsersQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'from' for for 'List' method.
func (p ListUsersQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'List	' method.
func (p ListUsersQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'List' method.
func (p ListUsersQueryParams) Order(value string) {
	p.order(value)
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListUsersQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List users. Reference: https://developer.aircall.io/api-references/#list-all-users
func (service *UsersService) List(opts *ListUsersQueryParams) (*UsersResponse, *Response, error) {
	_url := "users"

	responseBody := new(UsersResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET USER
//  https://developer.aircall.io/api-references/#retrieve-a-user
//  ***********************************************************************************

// Get a user. Reference: https://developer.aircall.io/api-references/#retrieve-a-user
func (service *UsersService) Get(userID int) (*UserResponse, *Response, error) {
	_url := fmt.Sprintf("users/%d", userID)

	responseBody := new(UserResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE USER
//  https://developer.aircall.io/api-references/#create-a-user
//  ***********************************************************************************

// Create a user. Reference: https://developer.aircall.io/api-references/#create-a-user
func (service *UsersService) Create(user *CreateUpdateUser) (*UserResponse, *Response, error) {
	_url := "users"

	responseBody := new(UserResponse)
	response, err := service.client.Post(_url, user, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE USER
//  https://developer.aircall.io/api-references/#update-a-user
//  ***********************************************************************************

// Update a user. Reference: https://developer.aircall.io/api-references/#update-a-user
func (service *UsersService) Update(userID int, user *CreateUpdateUser) (*UserResponse, *Response, error) {
	_url := fmt.Sprintf("users/%d", userID)

	responseBody := new(UserResponse)
	response, err := service.client.Put(_url, user, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE USER
//  https://developer.aircall.io/api-references/#delete-a-user
//  ***********************************************************************************

// Delete a user. Reference: https://developer.aircall.io/api-references/#delete-a-user
func (service *UsersService) Delete(userID int) (*Response, error) {
	_url := fmt.Sprintf("users/%d", userID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  LIST ALL USERS AVAILABLILITY
//  https://developer.aircall.io/api-references/#retrieve-list-of-users-availability
//  ***********************************************************************************

// Creates Query parameters for 'List'
func (pq UserQueries) NewListUsersAvailability() *ListUsersAvailabilityQueryParams {
	return &ListUsersAvailabilityQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'from' for for 'ListAvailability' method.
func (p ListUsersAvailabilityQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'ListAvailability	' method.
func (p ListUsersAvailabilityQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'ListAvailability' method.
func (p ListUsersAvailabilityQueryParams) Order(value string) {
	p.order(value)
}

// Set 'page' and 'per_page' for for 'ListAvailability' method.
func (p ListUsersAvailabilityQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List users availabilies. Reference: https://developer.aircall.io/api-references/#retrieve-list-of-users-availability
func (service *UsersService) ListAvailabilities(opts *ListUsersAvailabilityQueryParams) (*UserAvailabilitiesResponse, *Response, error) {
	_url := "users/availabilities"

	responseBody := new(UserAvailabilitiesResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET USER AVAILABILITY
//  https://developer.aircall.io/api-references/#check-availability-of-a-user
//  ***********************************************************************************

// Get a user availability. Reference: https://developer.aircall.io/api-references/#check-availability-of-a-user
func (service *UsersService) GetAvailability(userID int) (*UserAvailability, *Response, error) {
	_url := fmt.Sprintf("users/%d/availability", userID)

	responseBody := new(UserAvailability)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  START OUTBOUND CALL
//  https://developer.aircall.io/api-references/#start-an-outbound-call
//  ***********************************************************************************

// Start an outbound call. Reference: https://developer.aircall.io/api-references/#start-an-outbound-call
func (service *UsersService) StartOutboundCall(userID int, call *NewUserCall) (*Response, error) {
	_url := fmt.Sprintf("users/%d/calls", userID)

	return service.client.Post(_url, call, nil)
}

//  ***********************************************************************************
//  DIAL NUMBER IN PHONE
//  https://developer.aircall.io/api-references/#dial-a-phone-number-in-the-phone
//  ***********************************************************************************

// Dial a phone number in the phone. Reference: https://developer.aircall.io/api-references/#dial-a-phone-number-in-the-phone
func (service *UsersService) DialNumber(userID int, call *NewUserCall) (*Response, error) {
	_url := fmt.Sprintf("users/%d/dail", userID)

	return service.client.Post(_url, call, nil)
}
