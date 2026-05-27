package users

import (
	"exampleWithGin/internal/users/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", c.GetAllUsers)
		users.GET("/:id", c.GetUser)
		users.POST("", c.CreateUser)
		users.PUT("/:id", c.UpdateUser)
		users.DELETE("/:id", c.DeleteUser)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	createdUser, err := c.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)

}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	user, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	updatedUser, err := c.userService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
