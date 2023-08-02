package entities

import "github.com/google/uuid"

type (
	MandorRating struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		Rating   int       `gorm:"type:int" json:"rating"`
		Review   string    `gorm:"type:varchar(128)" json:"review"`
		
		MandorID uuid.UUID `gorm:"type:uuid" json:"mandor_id"`
		UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`

		Timestamp
	}
)
