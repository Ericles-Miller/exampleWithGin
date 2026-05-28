package users

import (
	"errors"
	"exampleWithGin/internal/users/models"
	"exampleWithGin/pkg/appErrors"
	"exampleWithGin/pkg/httputil"
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
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("Invalid request body"))
		return
	}

	createdUser, err := c.userService.CreateUser(&user)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrBadRequest):
			ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("User already exists"))
		default:
			ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.User]("Failed to create user"))
		}
		return
	}

	ctx.JSON(http.StatusCreated, httputil.Success(createdUser))
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("Invalid UUID"))
		return
	}

	user, err := c.userService.GetUser(id)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrNotFound):
			ctx.JSON(http.StatusNotFound, httputil.Fail[*models.User]("User not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.User]("Failed to retrieve user"))
		}
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(user))
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httputil.Fail[[]*models.User]("Failed to retrieve users"))
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(users))
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("Invalid UUID"))
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("Invalid request body"))
		return
	}

	updatedUser, err := c.userService.UpdateUser(id, &user)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrNotFound):
			ctx.JSON(http.StatusNotFound, httputil.Fail[*models.User]("User not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.User]("Failed to update user"))
		}
		return
	}

	ctx.JSON(http.StatusOK, httputil.Success(updatedUser))
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Fail[*models.User]("Invalid UUID"))
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrNotFound):
			ctx.JSON(http.StatusNotFound, httputil.Fail[*models.User]("User not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, httputil.Fail[*models.User]("Failed to delete user"))
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
