package handler

import (
	"net/http"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	TopUp(c *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) Register(c *gin.Context) {
	var userPayload dto.RegisterRequest

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := u.userService.Register(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (u *userHandler) TopUp(c *gin.Context) {
	user := c.MustGet("userData").(*entity.User)

	userID := int(user.ID)

	var userPayload dto.TopUpRequest

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := u.userService.TopUp(userPayload, userID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (u *userHandler) Login(c *gin.Context) {
	var userPayload dto.LoginRequest

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := u.userService.Login(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}