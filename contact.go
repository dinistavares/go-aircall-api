package aircall

import "fmt"

// Contact service
type ContactsService service

type ContactsResponse struct {
	Meta     *GenericResponseMeta `json:"meta,omitempty"`
	Contacts *[]Contact           `json:"contacts,omitempty"`
}

type ContactResponse struct {
	Contact *Contact           `json:"contact,omitempty"`
}

type ContactInfoResponse struct {
	PhoneDetail *ContactInfo           `json:"phone_detail,omitempty"`
	EmailDetail *ContactInfo           `json:"email_detail,omitempty"`
}

type CreateUpdateContact struct {
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Information  string         `json:"information,omitempty"`
	CompanyName  string         `json:"company_name,omitempty"`
	PhoneNumbers []ContactInfo `json:"phone_numbers,omitempty"`
	Emails       []ContactInfo `json:"emails,omitempty"`
}

type Contact struct {
	ID           int            `json:"id,omitempty"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	CompanyName  string         `json:"company_name,omitempty"`
	Information  string         `json:"information,omitempty"`
	IsShared     bool           `json:"is_shared,omitempty"`
	DirectLink   string         `json:"direct_link,omitempty"`
	CreatedAt    int            `json:"created_at,omitempty"`
	UpdatedAt    int            `json:"updated_at,omitempty"`
	PhoneNumbers *[]ContactInfo `json:"phone_numbers,omitempty"`
	Emails       *[]ContactInfo `json:"emails,omitempty"`
}

type ContactInfo struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContactQueries struct{}

// Query parameters for 'ListContacts' method.
type ListContactsQueryParams struct {
	QueryValues
}

// Query parameters for 'SearchContacts' method.
type SearchContactsQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *ContactsService) Query() *ContactQueries {
	return &ContactQueries{}
}

//  ***********************************************************************************
//  LIST ALL CONTACTS
//  https://developer.aircall.io/api-references/#list-all-contacts
//  ***********************************************************************************

// Creates Query parameters for 'ListContacts'
func (pq ContactQueries) NewListContacts() *ListContactsQueryParams {
	return &ListContactsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set 'from' for for 'List' method.
func (p ListContactsQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'List' method.
func (p ListContactsQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'List' method.
func (p ListContactsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'order_by' by for for 'List' method.
func (p ListContactsQueryParams) OrderBy(value string) {
	p.orderBy(value)
}

// Set 'page' and 'per_page' for for 'List' method.
func (p ListContactsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// List contacts. Reference: https://developer.aircall.io/api-references/#list-all-contacts
func (service *ContactsService) List(opts *ListContactsQueryParams) (*ContactsResponse, *Response, error) {
	_url := "contacts"

	responseBody := new(ContactsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  SEARCH CONTACTS
//  https://developer.aircall.io/api-references/#search-contacts
//  ***********************************************************************************

// Creates Query parameters for 'ListContacts'
func (pq ContactQueries) NewSearchContacts() *SearchContactsQueryParams {
	return &SearchContactsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set from for for 'Search' method.
func (p SearchContactsQueryParams) From(value string) {
	p.from(value)
}

// Set 'to' for for 'Search' method.
func (p SearchContactsQueryParams) To(value string) {
	p.to(value)
}

// Set 'order' for for 'Search' method.
func (p SearchContactsQueryParams) Order(value string) {
	p.order(value)
}

// Set 'order_by' by for for 'Search' method.
func (p SearchContactsQueryParams) OrderBy(value string) {
	p.orderBy(value)
}

// Set 'phone_number' for for 'Search' method.
func (p SearchContactsQueryParams) PhoneNumber(value string) {
	p.set("phone_number", value)
}

// Set 'email' for for 'Search' method.
func (p SearchContactsQueryParams) Email(value string) {
	p.set("email", value)
}

// Set 'page' and 'per_page' for for 'Search' method.
func (p SearchContactsQueryParams) Paginate(page int, perPage int) {
	p.page(page, perPage)
}

// Search contacts. Reference: https://developer.aircall.io/api-references/#search-contacts
func (service *ContactsService) Search(opts *SearchContactsQueryParams) (*ContactsResponse, *Response, error) {
	_url := "contacts/search"

	responseBody := new(ContactsResponse)
	response, err := service.client.Get(_url, opts, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET CONTACT
//  https://developer.aircall.io/api-references/#retrieve-a-contact
//  ***********************************************************************************

// Get a contact. Reference: https://developer.aircall.io/api-references/#retrieve-a-contact
func (service *ContactsService) Get(contactID int) (*ContactResponse, *Response, error) {
	_url := fmt.Sprintf("contacts/%d", contactID)

	responseBody := new(ContactResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  CREATE CONTACT
//  https://developer.aircall.io/api-references/#create-a-contact
//  ***********************************************************************************

// Create a contact. Reference: https://developer.aircall.io/api-references/#create-a-contact
func (service *ContactsService) Create(contact *CreateUpdateContact) (*ContactResponse, *Response, error) {
	_url := "contacts"

	responseBody := new(ContactResponse)
	response, err := service.client.Post(_url, contact, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE CONTACT
//  https://developer.aircall.io/api-references/#update-a-contact
//  ***********************************************************************************

// Update a contact. Reference: https://developer.aircall.io/api-references/#update-a-contact
func (service *ContactsService) Update(contactID int, contact *CreateUpdateContact) (*CreateUpdateContact, *Response, error) {
	_url := fmt.Sprintf("contacts/%d", contactID)
	

	responseBody := new(CreateUpdateContact)
	response, err := service.client.Post(_url, contact, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE CONTACT
//  https://developer.aircall.io/api-references/#delete-a-contact
//  ***********************************************************************************

// Delete a contact. Reference: https://developer.aircall.io/api-references/#delete-a-contact
func (service *ContactsService) Delete(contactID int) (*Response, error) {
	_url := fmt.Sprintf("contacts/%d", contactID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  ADD PHONE NUMBER TO CONTACT
//  https://developer.aircall.io/api-references/#add-phone-number-to-a-contact
//  ***********************************************************************************

// Add a phone number to a contact. Reference: https://developer.aircall.io/api-references/#add-phone-number-to-a-contact
func (service *ContactsService) AddNumber(contactID int, number *ContactInfo) (*ContactInfoResponse, *Response, error) {
	_url := fmt.Sprintf("contacts/%d/phone_details", contactID)

	responseBody := new(ContactInfoResponse)
	response, err := service.client.Post(_url, number, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE PHONE NUMBER TO CONTACT
//  https://developer.aircall.io/api-references/#update-phone-number-to-a-contact
//  ***********************************************************************************

// Update a phone number to a contact. Reference: https://developer.aircall.io/api-references/#update-phone-number-to-a-contact
func (service *ContactsService) UpdateNumber(contactID int, numberID int, number *ContactInfo) (*ContactInfoResponse, *Response, error) {
	_url := fmt.Sprintf("contacts/%d/phone_details/%d", contactID, numberID)

	responseBody := new(ContactInfoResponse)
	response, err := service.client.Put(_url, number, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE PHONE NUMBER FROM CONTACT
//  https://developer.aircall.io/api-references/#delete-phone-number-from-a-contact
//  ***********************************************************************************

// Delete a phone number from a contact. Reference: https://developer.aircall.io/api-references/#delete-phone-number-from-a-contact
func (service *ContactsService) DeleteNumber(contactID int, numberID int) (*Response, error) {
	_url := fmt.Sprintf("contacts/%d/phone_details/%d", contactID, numberID)

	return service.client.Delete(_url)
}

//  ***********************************************************************************
//  ADD EMAIL TO CONTACT
//  https://developer.aircall.io/api-references/#add-email-to-a-contact
//  ***********************************************************************************

// Add an email to a contact. Reference: https://developer.aircall.io/api-references/#add-email-to-a-contact
func (service *ContactsService) AddEmail(contactID int, email *ContactInfo) (*ContactInfoResponse, *Response, error) {
	_url := fmt.Sprintf("contacts/%d/email_details", contactID)

	responseBody := new(ContactInfoResponse)
	response, err := service.client.Post(_url, email, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  UPDATE EMAIL TO CONTACT
//  https://developer.aircall.io/api-references/#update-email-to-a-contact
//  ***********************************************************************************

// Update a phone number to a contact. Reference: https://developer.aircall.io/api-references/#update-phone-number-to-a-contact
func (service *ContactsService) UpdateEmail(contactID int, emailID int, email *ContactInfo) (*ContactInfoResponse, *Response, error) {
	_url := fmt.Sprintf("contacts/%d/email_details/%d", contactID, emailID)

	responseBody := new(ContactInfoResponse)
	response, err := service.client.Put(_url, email, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  DELETE EMAIL FROM CONTACT
//  https://developer.aircall.io/api-references/#delete-email-from-a-contact
//  ***********************************************************************************

// Delete a phone number from a contact. Reference: https://developer.aircall.io/api-references/#delete-phone-number-from-a-contact
func (service *ContactsService) DeleteEmail(contactID int, emailID int) (*Response, error) {
	_url := fmt.Sprintf("contacts/%d/email_details/%d", contactID, emailID)

	return service.client.Delete(_url)
}
