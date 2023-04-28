package service

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/pkg/helpers"
	"github.com/fydhfzh/fp-4/repository/user_repository"
	"github.com/gin-gonic/gin"
)

type authService struct {
	userRepo user_repository.UserRepository
}

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(c *gin.Context){
		bearerToken := c.GetHeader("Authorization")

		userData, err := helpers.GetUserData(bearerToken)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserByEmail(userData.Email)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
		}

		c.Set("userData", userData)
		c.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userData").(*entity.User)

		role := user.Role

		if role != "admin" {
			forbiddenError := errs.NewUnauthorizedError("forbidden request")
			
			c.AbortWithStatusJSON(forbiddenError.Status(), forbiddenError)
			return
		}

		c.Next()
	}
}