package models

type PaginationOpts struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginatedResult struct {
	Total int64         `json:"total"`
	Data  []interface{} `json:"data"`
}
