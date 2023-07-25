package entities

type (
	Bank struct {
		ID         int    `gorm:"primary_key;auto_increment" json:"id"`
		NamaBank   string `gorm:"type:varchar(100)" json:"nama_bank"`
		NoRekening string `gorm:"type:varchar(100)" json:"no_rekening,omitempty"`
		AtasNama   string `gorm:"type:varchar(100)" json:"atas_nama,omitempty"`

		MandorID int    `gorm:"type:int" json:"mandor_id,omitempty"`
		Mandor   Mandor `gorm:"foreignKey:MandorID" json:"mandor,omitempty"`
	}
)
