package aircall

// Contact service
type ContactsService service

type ListContactsResponse struct {
	Meta     *GenericResponseMeta `json:"meta,omitempty"`
	Contacts *[]Contact    `json:"contacts,omitempty"`
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

// Set sort for for 'GetEvents' method.
func (p ListContactsQueryParams) From(value string) {
	p.from(value)
}

func (p ListContactsQueryParams) To(value string) {
	p.to(value)
}

func (p ListContactsQueryParams) Order(value string) {
	p.order(value)
}

func (p ListContactsQueryParams) OrderBy(value string) {
	p.orderBy(value)
}

// List contacts. Reference: https://developer.aircall.io/api-references/#list-all-contacts
func (service *ContactsService) ListContacts(opts *ListContactsQueryParams) (*ListContactsResponse, *Response, error) {
	_url := "contacts"

	accounts := new(ListContactsResponse)
	response, err := service.client.Get(_url, opts, accounts)

	if err != nil {
		return nil, response, err
	}

	return accounts, response, nil
}
