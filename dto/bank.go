package dto

type (
	BankCreateDTO struct {
		NamaBank   string `json:"nama_bank"`
		NoRekening string `json:"no_rekening"`
		AtasNama   string `json:"atas_nama"`	
	}
)