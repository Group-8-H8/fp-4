package service

import (
	"fmt"
	"net/http"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/pkg/helpers"
	"github.com/fydhfzh/fp-4/repository/user_repository"
)

// User service struct
type userService struct {
	userRepo user_repository.UserRepository
}

// User service interface
type UserService interface {
	Register(userPayload dto.RegisterRequest) (*dto.RegisterResponse, errs.Errs)
	Login(userPayload dto.LoginRequest) (*dto.LoginResponse, errs.Errs)
	TopUp(userPayload dto.TopUpRequest, id int) (*dto.TopUpResponse, errs.Errs)
}

// User service generator function
func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register service
func (u *userService) Register(userPayload dto.RegisterRequest) (*dto.RegisterResponse, errs.Errs) {
	// Validate user payload struct
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	// Create user entity
	userData := entity.User{
		FullName: userPayload.Fullname,
		Email: userPayload.Email,
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(userPayload.Password)

	if err != nil {
		return nil, err
	}

	userData.Password = hashedPassword

	// Register
	userRegistered, err := u.userRepo.Register(userData)

	if err != nil {
		return nil, err
	}

	// Create register response
	response := dto.RegisterResponse{
		ID: int(userRegistered.ID),
		Fullname: userRegistered.FullName,
		Email: userRegistered.Email,
		Password: userRegistered.Password,
		Balance: userRegistered.Balance,
		CreatedAt: userRegistered.CreatedAt,
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

// Top up service
func (u *userService) TopUp(userPayload dto.TopUpRequest, id int) (*dto.TopUpResponse, errs.Errs) {	
	// Validate user payload struct
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}
	
	// Top up
	balance, err := u.userRepo.TopUp(userPayload.Balance, id)

	if err != nil {
		return nil, err
	}

	// Create top up response
	response := dto.TopUpResponse{
		Message: "Your balance has been successfully updated to Rp " + fmt.Sprint(balance),
		StatusCode: http.StatusOK,
	}

	return &response, nil
}

// Login service
func (u *userService) Login(userPayload dto.LoginRequest) (*dto.LoginResponse, errs.Errs) {
	// Validate user payload struct
	err := helpers.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	// Get user by email
	user, err := u.userRepo.GetUserByEmail(userPayload.Email)

	if err != nil {
		return nil, err
	}

	// Compare password
	match := helpers.ComparePassword(user.Password, userPayload.Password)

	if !match {
		return nil, errs.NewBadRequestError("password/email is not correct")
	}

	// Generate token if match
	signedToken, tokenErr := helpers.GenerateToken(int(user.ID), user.Email, user.Role)

	if tokenErr != nil {
		return nil, tokenErr
	}

	// Create login response
	response := dto.LoginResponse{
		Token: signedToken,
		StatusCode: http.StatusOK,
	}

	return &response, nil
}