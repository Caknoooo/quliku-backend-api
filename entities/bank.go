package entities

import "github.com/google/uuid"

type (
	Bank struct {
		ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NamaBank   string    `gorm:"type:varchar(100)" json:"nama_bank"`
		NoRekening string    `gorm:"type:varchar(100)" json:"no_rekening,omitempty"`
		AtasNama   string    `gorm:"type:varchar(100)" json:"atas_nama,omitempty"`

		MandorID uuid.UUID `gorm:"type:uuid" json:"mandor_id,omitempty"`
		Mandor   Mandor    `gorm:"foreignKey:MandorID" json:"-"`
	}
)
