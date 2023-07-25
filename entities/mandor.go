package entities

import "github.com/google/uuid"

type (
	Mandor struct {
		ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NamaLengkap string    `gorm:"type:varchar(100)" json:"nama_lengkap"`
		Username    string    `gorm:"type:varchar(100)" json:"username"`
		NoTelp      string    `gorm:"type:varchar(30)" json:"no_telp"`
		Email       string    `gorm:"type:varchar(100)" json:"email"`
		Password    string    `gorm:"type:varchar(100)" json:"password"`
		AsalKota    string    `gorm:"type:varchar(100)" json:"asal_kota"`

		// Kualifikasi Diri
		Klasifikasi                string `gorm:"type:varchar(100)" json:"klasifikasi"`
		DeskripsiDetailKlasifikasi string `gorm:"type:varchar(100)" json:"deskripsi_detail_klasifikasi"`
		PengalamanKerja            string `gorm:"type:varchar(100)" json:"pengalaman_kerja"`
		HargaMandor                int    `gorm:"type:int" json:"harga_mandor,omitempty"`

		// Range Kuli
		RangeKuliAwal  int `gorm:"type:int" json:"range_kuli_awal"`
		RangeKuliAkhir int `gorm:"type:int" json:"range_kuli_akhir"`

		// Unggah Dokumen
		FotoProfil     string `gorm:"type:varchar(100)" json:"foto_profil"`
		FotoKTP        string `gorm:"type:varchar(100)" json:"foto_ktp"`
		FotoSertifikat string `gorm:"type:varchar(100)" json:"foto_sertifikat"`
		FotoPortofolio string `gorm:"type:varchar(100)" json:"foto_portofolio"`

		// Data Bank
		Banks []Bank `gorm:"foreignKey:MandorID" json:"banks,omitempty"`

		Role       string `gorm:"type:varchar(100)" json:"role"`
		IsVerified bool   `gorm:"type:boolean;default:false" json:"is_verified"`

		Timestamp
	}
)
