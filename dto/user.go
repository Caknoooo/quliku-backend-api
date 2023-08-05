package dto

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPasswordNotMatch                  = errors.New("password not match")
	ErrPasswordMatch                     = errors.New("the password is the same as before")
	ErrorUserNotFound                    = errors.New("user not found")
	ErrorFailedGenerateVerificationCode  = errors.New("failed generate verification code")
	ErrorVerificationCodeNotMatch        = errors.New("verification code not match")
	ErrorExpiredVerificationCode         = errors.New("expired verification code")
	ErrorNotExpiredVerificationCode      = errors.New("not expired verification code")
	ErrorUserVerificationCodeNotActive   = errors.New("user verification code not active")
	ErrorUserVerificationCodeAlreadyUsed = errors.New("user verification code already used")
	ErrorUserAlreadyActive               = errors.New("user already active")
)

type (
	UserCreateDTO struct {
		ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NamaLengkap     string    `gorm:"type:varchar(100)" form:"nama_lengkap" json:"nama_lengkap" binding:"required"`
		Username        string    `gorm:"type:varchar(100)" form:"username" json:"username" binding:"required"`
		NoTelp          string    `gorm:"type:varchar(20)" form:"no_telp" json:"no_telp" binding:"required"`
		Email           string    `gorm:"type:varchar(100)" form:"email" json:"email" binding:"required"`
		Password        string    `gorm:"type:varchar(100)" form:"password" json:"password" binding:"required"`
		ConfirmPassword string    `gorm:"-" form:"confirm_password" json:"confirm_password" binding:"required"`
	}

	UserUpdateDTO struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		Nama     *string   `gorm:"type:varchar(100)" form:"nama" json:"nama"`
		NoTelp   *string   `gorm:"type:varchar(20)" form:"no_telp" json:"no_telp"`
		Email    *string   `gorm:"type:varchar(100)" form:"email" json:"email"`
		Password *string   `gorm:"type:varchar(100)" form:"password" json:"password"`
	}

	UserLoginDTO struct {
		Email    string `json:"email" form:"email,omitempty"`
		Username string `json:"username" form:"username,omitempty"`
		Password string `json:"password" form:"password"`
	}

	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	MakeVerificationForgotPasswordRequest struct {
		Email string `json:"email" form:"email"`
	}

	MakeVerificationForgotPasswordResponse struct {
		ID    string `json:"id" form:"id"`
		Email string `json:"email" form:"email"`
	}

	KodeOTPForgotPasswordRequest struct {
		Email    string `json:"email" form:"email"`
		SendCode string `json:"send_code" form:"send_code"`
	}

	KodeOTPForgotPasswordResponse struct {
		ID string `json:"id" form:"id"`
	}

	ForgotPasswordRequest struct {
		ID          string `json:"id" form:"id"`
		NewPassword string `json:"new_password" form:"new_password"`
	}
)
