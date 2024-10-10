package aircall

type GenericResponseMeta struct {
	Count            int    `json:"count,omitempty"`
	Total            int    `json:"total,omitempty"`
	CurrentPage      int    `json:"current_page,omitempty"`
	PerPage          int    `json:"per_page,omitempty"`
	NextPageLink     string `json:"next_page_link,omitempty"`
	PreviousPageLink string `json:"previous_page_link,omitempty"`
}