package v1

import "time"

// UserCreateRequest represents the request body for creating a user
type UserCreateRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secretpassword123"`
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" binding:"email" example:"john@example.com"`
	Password string `json:"password" example:"newpassword123"`
}

// UserResponse represents the response structure for user data
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Username  string    `json:"username" example:"johndoe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-10-26T12:34:56Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-10-26T12:34:56Z"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request parameters"`
	Message string `json:"message,omitempty" example:"Email is already taken"`
}

// ListUserResponse represents the paginated response for user listing
type ListUserResponse struct {
	Users      []UserResponse `json:"users"`
	TotalCount int64          `json:"total_count" example:"100"`
	Page       int            `json:"page" example:"1"`
	PerPage    int            `json:"per_page" example:"10"`
}
