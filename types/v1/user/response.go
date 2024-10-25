package user

import "time"

type Response struct {
	ID        uint      `json:"id" example:"1"`
	Username  string    `json:"username" example:"johndoe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-10-26T12:34:56Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-10-26T12:34:56Z"`
}

type ListResponse struct {
	Users      []Response `json:"users"`
	TotalCount int64      `json:"total_count" example:"100"`
	Page       int        `json:"page" example:"1"`
	PerPage    int        `json:"per_page" example:"10"`
}
