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

// Register godoc
// @Summary Create new account
// @Description Parse request body and add new user data in the database
// @Tags user
// @Accept json
// @Produce json
// @Param RequestBody body dto.RegisterRequest true "Request body json"
// @Success 201 {object} dto.RegisterResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /users/register [post]
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

// TopUp godoc
// @Summary Top up balance on logged account
// @Description Parse request body and update user balance data in the database
// @Tags user
// @Accept json
// @Produce json
// @Param RequestBody body dto.TopUpRequest true "Request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 201 {object} dto.TopUpResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /users/topup [patch]
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

// Login godoc
// @Summary Login using created account
// @Description Parse request body and check if user exist before create json web token for auth
// @Tags user
// @Accept json
// @Produce json
// @Param RequestBody body dto.LoginRequest true "Request body json"
// @Success 201 {object} dto.LoginResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /users/login [post]
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