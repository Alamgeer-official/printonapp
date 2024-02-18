package models

type Pagination struct {
	TotalCount int64       `json:"total_count"`
	TotalPage  int         `json:"total_page"`
	Data       interface{} `json:"data"`
}
