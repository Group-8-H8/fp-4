package dto

import (
	"time"
)

type RegisterRequest struct {
	Fullname string `json:"full_name" valid:"required~full name is empty"`
	Email    string `json:"email" valid:"required~email is empty,email~email format is not correct"`
	Password string `json:"password" valid:"required~password is empty,minstringlength(6)~password has to be atleast 6 character long"`
}

type RegisterResponse struct {
	ID        int    `json:"id,omitempty"`
	Fullname  string `json:"full_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Balance   int    `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	StatusCode int `json:"-"`
}

type UserResponse struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Fullname  string `json:"full_name,omitempty"`
	Balance   int    `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type LoginRequest struct {
	Email string `json:"email" valid:"required~password is empty,email~email format is not correct"`
	Password string `json:"password" valid:"required~password is empty"`
}

type LoginResponse struct {
	Token string `json:"token"`
	StatusCode int `json:"-"`
}

type TopUpRequest struct {
	Balance int `json:"balance" valid:"required,range(0|100000000)~balance cant hold more than 100000000"`
}

type TopUpResponse struct {
	Message string `json:"message,omitempty"`
	StatusCode int `json:"-"`
}