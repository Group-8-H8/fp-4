package user_repository

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
)

// User repository interface
type UserRepository interface {
	Register(user entity.User) (*entity.User, errs.Errs)
	TopUp(balance int, userId int) (int, errs.Errs)
	GetUserById(userId int) (*entity.User, errs.Errs)
	GetUserByEmail(email string) (*entity.User, errs.Errs)
	PayOrder(totalPrice int, userID int) errs.Errs
}