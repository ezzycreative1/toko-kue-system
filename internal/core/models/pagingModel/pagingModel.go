package pagingModel

type DataPagingResponse struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalRecord int64 `json:"total_record"`
	TotalPages  int64 `json:"total_pages"`
}

type DataPagingResponseSuccess struct {
	PageNumber       int           `json:"page_number"`
	PageSize         int           `json:"page_size"`
	TotalRecordCount int64         `json:"total_record_count"`
	Records          []interface{} `json:"records"`
}

type DataPagingResponse400 struct {
	PageNumber       int           `json:"page_number"`
	PageSize         int           `json:"page_size"`
	TotalRecordCount int64         `json:"total_record_count"`
	Records          []interface{} `json:"records"`
}
