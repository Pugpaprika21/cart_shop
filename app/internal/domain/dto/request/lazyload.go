package request

type Lazyload struct {
	PageNo      *int    `json:"page_no"`
	PageSize    *int    `json:"page_size"`
	SortBy      *string `json:"sort_by"`
	SearchField *string `json:"search_field"`
	IsActive    *string `json:"is_active"`
}
