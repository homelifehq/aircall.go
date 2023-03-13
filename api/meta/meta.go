package meta

type Meta struct {
	Count            uint   `json:"count"`
	Total            uint   `json:"total"`
	CurrentPage      uint   `json:"current_page"`
	PerPage          uint   `json:"per_page"`
	NextPageLink     string `json:"next_page_link"`
	PreviousPageLink string `json:"previous_page_link"`
}
