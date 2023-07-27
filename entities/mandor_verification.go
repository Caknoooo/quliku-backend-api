package entities

import (
	"time"

	"github.com/google/uuid"
)

type (
	MandorVerification struct {
		ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		ReceiveCode string    `gorm:"type:varchar(7)" json:"receive_code"`
		SendCode    string    `gorm:"type:varchar(7)" json:"send_code,omitempty"`
		ExpiredAt   time.Time `gorm:"timestamp with time zone" json:"expired_at"`
		IsActive    bool      `gorm:"type:boolean;default:false" json:"is_active"`

		MandorID uuid.UUID `gorm:"type:uuid;not null" json:"mandor_id"`
		Mandor   Mandor    `gorm:"foreignKey:MandorID" json:"mandor"`
	}
)