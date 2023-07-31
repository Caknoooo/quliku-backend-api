package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Mandor struct {
		ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NamaLengkap string    `gorm:"type:varchar(64)" json:"nama_lengkap"`
		NoTelp      string    `gorm:"type:varchar(32)" json:"no_telp"`
		Email       string    `gorm:"type:varchar(32)" json:"email"`
		Password    string    `gorm:"type:varchar(128)" json:"password"`
		AsalKota    string    `gorm:"type:varchar(32)" json:"asal_kota"`
		Status      string    `gorm:"type:varchar(32);default:waiting" json:"status"`

		// Kualifikasi Diri
		Klasifikasi                string `gorm:"type:varchar(32)" json:"klasifikasi"`
		DeskripsiDetailKlasifikasi string `gorm:"type:varchar(128)" json:"deskripsi_detail_klasifikasi"`
		PengalamanKerja            int    `gorm:"type:int" json:"pengalaman_kerja"`
		HargaMandor                int    `gorm:"type:int" json:"harga_mandor,omitempty"`

		// Range Kuli
		RangeKuliAwal  int `gorm:"type:int" json:"range_kuli_awal"`
		RangeKuliAkhir int `gorm:"type:int" json:"range_kuli_akhir"`

		// Unggah Dokumen
		FotoProfil     string `gorm:"type:varchar(128)" json:"foto_profil"`
		FotoKTP        string `gorm:"type:varchar(128)" json:"foto_ktp"`
		FotoSertifikat string `gorm:"type:varchar(128)" json:"foto_sertifikat"`
		FotoPortofolio string `gorm:"type:varchar(128)" json:"foto_portofolio"`

		// Data Bank
		Banks []Bank `gorm:"foreignKey:MandorID" json:"banks,omitempty"`

		Role            string `gorm:"type:varchar(128)" json:"role"`
		IsVerified      bool   `gorm:"type:boolean;default:false" json:"is_verified"`
		IsVerifiedAdmin bool   `gorm:"type:boolean;default:false" json:"is_verified_admin"`

		Timestamp
	}
)

func (m *Mandor) BeforeCreate(tx *gorm.DB) error {
	var err error
	m.Password, err = helpers.HashPassword(m.Password)
	if err != nil {
		return err
	}
	return nil
}
