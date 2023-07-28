package dto

import "errors"

var (
	ErrEmailNotFound    = errors.New("email not found")
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
