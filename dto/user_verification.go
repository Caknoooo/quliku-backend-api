package dto

import "github.com/google/uuid"

type (
	UserVerificationDTO struct {
		UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id" binding:"required"`
		SendCode string    `gorm:"type:varchar(7)" json:"send_code" binding:"required"`
	}

	ResendVerificationCode struct {
		UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id" binding:"required"`
	}

	FailedVerificationLoginDTO struct {
		Email string `json:"email"`
	}
)
