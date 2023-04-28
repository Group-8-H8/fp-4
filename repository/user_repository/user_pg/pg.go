package user_pg

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/repository/user_repository"
	"gorm.io/gorm"
)

// User repository struct
type userRepository struct {
	db *gorm.DB
}

// User repository generator function
func NewPGUserRepository(db *gorm.DB) user_repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Register repository
func (u *userRepository) Register(user entity.User) (*entity.User, errs.Errs) {
	// Create user data
	result := u.db.Create(&user)

	if err := result.Error; err != nil {
		return nil, errs.NewBadRequestError("email already registered")
	}

	return &user, nil
}

// Top up repository
func (u *userRepository) TopUp(balance int, userId int) (int, errs.Errs) {
	var user entity.User

	// Get user data
	result := u.db.First(&user, userId)

	if err := result.Error; err != nil {
		return 0, errs.NewBadRequestError(err.Error())
	}

	// Check if current balanca plus additional balance isnt more than 10000000
	if balance + user.Balance >= 100000000 {
		return 0, errs.NewBadRequestError("balance cant hold more than 100 millions")
	}

	// Update user balance data
	result = u.db.Model(&user).Update("balance", balance + user.Balance)

	if err := result.Error; err != nil {
		return 0, errs.NewBadRequestError(err.Error())
	}

	return user.Balance, nil
}

func (u *userRepository) GetUserByEmail(email string) (*entity.User, errs.Errs) {
	var user entity.User

	result := u.db.Where("email = ?", email).First(&user)

	if err := result.Error; err != nil {
		return nil, errs.NewNotFoundError("user not found")
	}

	return &user, nil
}

// Get user repository
func (u *userRepository) GetUserById(userId int) (*entity.User, errs.Errs) {
	var user entity.User

	// Get user data
	result := u.db.First(&user, userId)

	if err := result.Error; err != nil {
		return nil, errs.NewNotFoundError("user not found")
	}

	return &user, nil
}

// Pay order repository
func (u *userRepository) PayOrder(totalPrice int, userID int) errs.Errs {
	// Get user data
	user, err := u.GetUserById(userID)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	user.Balance -= totalPrice

	// Update user data
	result := u.db.Save(&user)

	if err := result.Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
