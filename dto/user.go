package dto

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPasswordNotMatch = errors.New("password not match")
	ErrorUserNotFound   = errors.New("user not found")
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
)
