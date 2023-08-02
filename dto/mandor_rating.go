package dto

import "github.com/google/uuid"

type (
	MandorRatingCreateDTO struct {
		ID       uuid.UUID `json:"id"`
		Rating   int       `json:"rating"`
		Review   string    `json:"review"`
		MandorID uuid.UUID `json:"mandor_id"`
		UserID   uuid.UUID `json:"user_id"`
	}

	MandorRatingDeleteDTO struct {
		ID uuid.UUID `json:"id"`
	}
)
