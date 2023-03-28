package models

type Paginated struct {
	TotalPages int `json:"total_pages"`
	Count      int `json:"count"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
}
