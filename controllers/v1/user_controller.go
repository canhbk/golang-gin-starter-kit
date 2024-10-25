package v1

import (
	"net/http"
	"strconv"

	"github.com/canhbk/golang-gin-starter-kit/config"
	"github.com/canhbk/golang-gin-starter-kit/models"
	_ "github.com/canhbk/golang-gin-starter-kit/types/v1/common"
	_ "github.com/canhbk/golang-gin-starter-kit/types/v1/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// Create godoc
// @Summary      Create user
// @Description  Create a new user
// @Tags         v1/users
// @Accept       json
// @Produce      json
// @Param        request body     user.CreateRequest true "User Information"
// @Success      201    {object}  user.Response
// @Failure      400    {object}  common.ErrorResponse
// @Failure      409    {object}  common.ErrorResponse
// @Router       /api/v1/users [post]
func (uc *UserController) Create(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Internal server error",
			Message: "Failed to process password",
		})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Failed to create user",
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// List godoc
// @Summary      List users
// @Description  Get paginated list of users
// @Tags         v1/users
// @Accept       json
// @Produce      json
// @Param        request query    common.PaginationQuery false "Pagination params"
// @Success      200    {object}  user.ListResponse
// @Failure      400    {object}  common.ErrorResponse
// @Router       /api/v1/users [get]
func (uc *UserController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	var users []models.User
	var total int64

	offset := (page - 1) * perPage

	config.DB.Model(&models.User{}).Count(&total)
	result := config.DB.Offset(offset).Limit(perPage).Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Failed to fetch users",
			Message: result.Error.Error(),
		})
		return
	}

	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, ListUserResponse{
		Users:      userResponses,
		TotalCount: total,
		Page:       page,
		PerPage:    perPage,
	})
}

// Get godoc
// @Summary      Get user
// @Description  Get user by ID
// @Tags         v1/users
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "User ID"
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (uc *UserController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a positive integer",
		})
		return
	}

	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User not found",
			Message: "No user exists with the provided ID",
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// Update godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         v1/users
// @Accept       json
// @Produce      json
// @Param        id      path    uint              true  "User ID"
// @Param        request body    UserUpdateRequest true  "User Information"
// @Success      200     {object} UserResponse
// @Failure      400     {object} ErrorResponse
// @Failure      404     {object} ErrorResponse
// @Router       /api/v1/users/{id} [put]
func (uc *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a positive integer",
		})
		return
	}

	var req UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User not found",
			Message: "No user exists with the provided ID",
		})
		return
	}

	// Update fields if provided
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Internal server error",
				Message: "Failed to process password",
			})
			return
		}
		user.Password = string(hashedPassword)
	}

	result = config.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Failed to update user",
			Message: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         v1/users
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "User ID"
// @Success      204  {object}  nil
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [delete]
func (uc *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid user ID",
			Message: "User ID must be a positive integer",
		})
		return
	}

	result := config.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "User not found",
			Message: "No user exists with the provided ID",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
