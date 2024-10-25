package user

type CreateRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secretpassword123"`
}

type UpdateRequest struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" binding:"email" example:"john@example.com"`
	Password string `json:"password" example:"newpassword123"`
}
