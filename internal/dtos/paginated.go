package dtos

type PaginatedResponse struct {
	Total int64       `json:"total_count"`
	Data  interface{} `json:"data"`
}

func NewPaginatedResponse(total int64, data interface{}) *PaginatedResponse {
	return &PaginatedResponse{
		Total: total,
		Data:  data,
	}
}

type PaginationOpts struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
