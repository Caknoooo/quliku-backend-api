package dto

import (
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

var (
	ErrorEmailAlreadyExist    = errors.New("email sudah terdaftar")
	ErrorUsernameAlreadyExist = errors.New("username sudah terdaftar")
	ErrFotoProfile            = errors.New("foto profil tidak boleh kosong")
	ErrFotoSertifikat         = errors.New("foto sertifikat tidak boleh kosong")
)

type (
	MandorCreateDTO struct {
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
	}

	MandorNextDTO struct {
		Email       string `form:"email" json:"email" binding:"required"`
		Password    string `form:"password" json:"password" binding:"required"`
		NoTelp      string `form:"no_telp" json:"no_telp" binding:"required"`
		NamaLengkap string `form:"nama_lengkap" json:"nama_lengkap" binding:"required"`
		AsalKota    string `form:"asal_kota" json:"asal_kota" binding:"required"`

		// Kualifikasi Diri
		Klasifikasi                string `form:"klasifikasi" json:"klasifikasi" binding:"required"`
		DeskripsiDetailKlasifikasi string `form:"deskripsi_detail_klasifikasi" json:"deskripsi_detail_klasifikasi" binding:"required"`
		PengalamanKerja            int    `form:"pengalaman_kerja" json:"pengalaman_kerja" binding:"required"`

		// Range Kuli
		RangeKuliAwal  int `form:"range_kuli_awal" json:"range_kuli_awal"`
		RangeKuliAkhir int `form:"range_kuli_akhir" json:"range_kuli_akhir"`

		// Dokumen
		FotoProfil     *multipart.FileHeader `form:"foto_profil" json:"foto_profil"`
		FotoKTP        *multipart.FileHeader `form:"foto_ktp" json:"foto_ktp"`
		FotoSertifikat *multipart.FileHeader `form:"foto_sertifikat" json:"foto_sertifikat"`
		FotoPortofolio *multipart.FileHeader `form:"foto_portofolio" json:"foto_portofolio"`

		// Data Bank
		NamaBank   string `form:"nama_bank" json:"nama_bank"`
		NoRekening string `form:"no_rekening" json:"no_rekening"`
		AtasNama   string `form:"atas_nama" json:"atas_nama"`
	}

	MandorUpdateDTO struct {
		ID          uuid.UUID `form:"id" json:"id" `
		NamaLengkap *string   `form:"nama_lengkap" json:"nama_lengkap"`
		NoTelp      *string   `form:"no_telp" json:"no_telp"`
		AsalKota    *string   `form:"asal_kota" json:"asal_kota"`

		// Kualifikasi Diri
		Klasifikasi                *string `form:"klasifikasi" json:"klasifikasi"`
		DeskripsiDetailKlasifikasi *string `form:"deskripsi_detail_klasifikasi" json:"deskripsi_detail_klasifikasi"`
		PengalamanKerja            *int    `form:"pengalaman_kerja" json:"pengalaman_kerja"`
		HargaMandor                *int    `form:"harga_mandor" json:"harga_mandor"`

		// Range Kuli
		RangeKuliAwal  *int `form:"range_kuli_awal" json:"range_kuli_awal"`
		RangeKuliAkhir *int `form:"range_kuli_akhir" json:"range_kuli_akhir"`

		// Unggah Dokumen
		FotoProfil     *multipart.FileHeader `form:"foto_profil" json:"foto_profil"`
		FotoPortofolio *multipart.FileHeader `form:"foto_portofolio" json:"foto_portofolio"`

		// Data Bank
		NamaBank   *string `form:"nama_bank" json:"nama_bank"`
		NoRekening *string `form:"no_rekening" json:"no_rekening"`
		AtasNama   *string `form:"atas_nama" json:"atas_nama"`
	}

	MandorUpdateDTOResponse struct {
		ID          uuid.UUID `json:"id"`
		NamaLengkap string    `json:"nama_lengkap"`
		NoTelp      string    `json:"no_telp"`
		AsalKota    string    `json:"asal_kota"`

		// Kualifikasi Diri
		Klasifikasi                string `json:"klasifikasi"`
		DeskripsiDetailKlasifikasi string `json:"deskripsi_detail_klasifikasi"`
		PengalamanKerja            int    `json:"pengalaman_kerja"`
		HargaMandor                int    `json:"harga_mandor"`

		// Range Kuli
		RangeKuliAwal  int `json:"range_kuli_awal"`
		RangeKuliAkhir int `json:"range_kuli_akhir"`

		// Unggah Dokumen
		FotoProfil     string `json:"foto_profil"`
		FotoPortofolio string `json:"foto_portofolio"`

		// Data Bank
		NamaBank   string `json:"nama_bank"`
		NoRekening string `json:"no_rekening"`
		AtasNama   string `json:"atas_nama"`
	}

	MandorLoginDTO struct {
		Email    string `json:"email" form:"email" binding:"email"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	GetAllMandorResponse struct {
		ID          uuid.UUID `json:"id"`
		NamaLengkap string    `json:"nama_lengkap"`
		NoTelp      string    `json:"no_telp"`

		Klasifikasi string `json:"klasifikasi"`

		Status string `json:"status"`
	}

	ChangeStatusMandorRequest struct {
		MandorID uuid.UUID `json:"mandor_id"`
		Status   string    `json:"status"`
	}
)
