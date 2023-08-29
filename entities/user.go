package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaLengkap string    `gorm:"type:varchar(100)" json:"nama_lengkap"`
	Username    string    `gorm:"type:varchar(100)" json:"username"`
	NoTelp      string    `gorm:"type:varchar(30)" json:"no_telp"`
	Email       string    `gorm:"type:varchar(100)" json:"email"`
	Password    string    `gorm:"type:varchar(100)" json:"password"`
	Role        string    `gorm:"type:varchar(100)" json:"role"`
	IsVerified  bool      `gorm:"type:boolean;default:false" json:"is_verified"`

	CreateProjectUser []CreateProjectUser `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"create_project_user,omitempty"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
