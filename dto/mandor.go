package dto

import (
	"errors"
	"mime/multipart"
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
		PengalamanKerja            string `form:"pengalaman_kerja" json:"pengalaman_kerja" binding:"required"`

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
		NamaLengkap *string `form:"nama_lengkap" json:"nama_lengkap"`
		NoTelp      *string `form:"no_telp" json:"no_telp"`
		Password    *string `form:"password" json:"password"`
		AsalKota    *string `form:"asal_kota" json:"asal_kota"`

		// Kualifikasi Diri
		Klasifikasi                *string `form:"klasifikasi" json:"klasifikasi"`
		DeskripsiDetailKlasifikasi *string `form:"deskripsi_detail_klasifikasi" json:"deskripsi_detail_klasifikasi"`
		PengalamanKerja            *string `form:"pengalaman_kerja" json:"pengalaman_kerja"`
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

	MandorLoginDTO struct {
		Email    string `json:"email" form:"email,omitempty"`
		Password string `json:"password" form:"password"`
	}
)
