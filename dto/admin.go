package dto

import "errors"

var (
	ErrEmailNotFound    = errors.New("email not found")
	ErrNotAdminID			 = errors.New("your role is not admin")
	ErrAdminNotFound    = errors.New("admin not found")
)

type (
	AdminCreateDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AdminLoginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
