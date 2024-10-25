package common

type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request parameters"`
	Message string `json:"message,omitempty" example:"Email is already taken"`
}

type PaginationQuery struct {
	Page    int `form:"page,default=1" example:"1"`
	PerPage int `form:"per_page,default=10" example:"10"`
}
