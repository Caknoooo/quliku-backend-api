package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama       string    `gorm:"type:varchar(100)" json:"nama"`
	Email      string    `gorm:"type:varchar(100)" json:"email"`
	Password   string    `gorm:"type:varchar(100)" json:"password"`
	Role       string    `gorm:"type:varchar(100)" json:"role"`
	IsVerified bool      `gorm:"type:boolean" json:"is_verified"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.Password, err = helpers.HashPassword(a.Password)
	if err != nil {
		return err
	}
	return nil
}
