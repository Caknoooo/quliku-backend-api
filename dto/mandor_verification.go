package dto

import "github.com/google/uuid"

type (
	MandorVerificationDTO struct {
		MandorID uuid.UUID `gorm:"type:uuid;not null" json:"mandor_id" binding:"required"`
		SendCode string    `gorm:"type:varchar(7)" json:"send_code" binding:"required"`
	}

	ResendMandorVerificationCodeDTO struct {
		MandorID uuid.UUID `gorm:"type:uuid;not null" json:"mandor_id" binding:"required"`
	}

	FailedMandorVerificationLoginDTO struct {
		Email string `json:"email"`
	}
)
