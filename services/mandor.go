package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	LOCALHOST  = "http://localhost:8888/api/"
	IMAGE      = "image/get/"
	PRODUCTION = "https://quliku-backend-api-production.up.railway.app/api/"
)

type MandorService interface {
	RegisterMandorStart(ctx context.Context, mandorDTO dto.MandorCreateDTO) (dto.MandorCreateDTO, error)
	RegisterMandorEnd(ctx context.Context, mandorDTO dto.MandorNextDTO) (entities.Mandor, error)
}

type mandorService struct {
	mandorRepository repository.MandorRepository
}

func NewMandorService(mr repository.MandorRepository) MandorService {
	return &mandorService{
		mandorRepository: mr,
	}
}

func (ms *mandorService) RegisterMandorStart(ctx context.Context, mandorDTO dto.MandorCreateDTO) (dto.MandorCreateDTO, error) {
	email, _ := ms.GetMandorByEmail(ctx, mandorDTO.Email)
	if email {
		return dto.MandorCreateDTO{}, dto.ErrorEmailAlreadyExist
	}

	return dto.MandorCreateDTO{
		Email:    mandorDTO.Email,
		Password: mandorDTO.Password,
	}, nil
}

func (ms *mandorService) RegisterMandorEnd(ctx context.Context, mandorDTO dto.MandorNextDTO) (entities.Mandor, error) {
	var mandor entities.Mandor

	email, _ := ms.GetMandorByEmail(ctx, mandorDTO.Email)
	if email {
		return entities.Mandor{}, dto.ErrorEmailAlreadyExist
	}
	mandor.Email = mandorDTO.Email

	mandor.Password = mandorDTO.Password
	mandor.NoTelp = mandorDTO.NoTelp
	mandor.NamaLengkap = mandorDTO.NamaLengkap
	mandor.AsalKota = mandorDTO.AsalKota

	// Kualifikasi diri
	mandor.Klasifikasi = mandorDTO.Klasifikasi
	mandor.DeskripsiDetailKlasifikasi = mandorDTO.DeskripsiDetailKlasifikasi
	mandor.PengalamanKerja = mandorDTO.PengalamanKerja

	// Range Kuli
	mandor.RangeKuliAwal = mandorDTO.RangeKuliAwal
	mandor.RangeKuliAkhir = mandorDTO.RangeKuliAkhir
	mandor.Role = "mandor"

	path := "storage/images"

	// Unggah Dokumen
	// Opsional
	if mandorDTO.FotoProfil != nil {
		fotoProfil, err := utils.IsBase64(*mandorDTO.FotoProfil)
		if err != nil {
			return entities.Mandor{}, dto.ErrToBase64
		}

		_ = utils.SaveImage(fotoProfil, path, mandorDTO.FotoProfil.Filename)

		mandor.FotoProfil = GenerateFileName(path, mandorDTO.FotoProfil.Filename)
	}

	if mandorDTO.FotoSertifikat != nil {
		fotoSertifikat, err := utils.IsBase64(*mandorDTO.FotoSertifikat)
		if err != nil {
			return entities.Mandor{}, dto.ErrToBase64
		}

		_ = utils.SaveImage(fotoSertifikat, path, mandorDTO.FotoSertifikat.Filename)

		mandor.FotoSertifikat = GenerateFileName(path, mandorDTO.FotoSertifikat.Filename)
	}

	fotoKTP, err := utils.IsBase64(*mandorDTO.FotoKTP)
	if err != nil {
		return entities.Mandor{}, dto.ErrToBase64
	}

	fotoPortofolio, err := utils.IsBase64(*mandorDTO.FotoPortofolio)
	if err != nil {
		return entities.Mandor{}, dto.ErrToBase64
	}

	_ = utils.SaveImage(fotoKTP, path, mandorDTO.FotoKTP.Filename)

	_ = utils.SaveImage(fotoPortofolio, path, mandorDTO.FotoPortofolio.Filename)

	mandor.FotoKTP = GenerateFileName(path, mandorDTO.FotoKTP.Filename)

	mandor.FotoPortofolio = GenerateFileName(path, mandorDTO.FotoPortofolio.Filename)

	// Data Bank
	mandor.Banks = []entities.Bank{
		{
			ID:         uuid.New(),
			NamaBank:   mandorDTO.NamaBank,
			NoRekening: mandorDTO.NoRekening,
			AtasNama:   mandorDTO.AtasNama,
		},
	}

	return ms.mandorRepository.CreateMandor(ctx, mandor)
}

func GenerateFileName(path string, filename string) string {
	return LOCALHOST + IMAGE + path + "/" + filename
}

func (ms *mandorService) GetMandorByEmail(ctx context.Context, emailz string) (bool, error) {
	email, err := ms.mandorRepository.GetMandorByEmail(ctx, emailz)
	if err != nil {
		return false, err
	}

	if email.Email == "" {
		return false, nil
	}

	return true, nil
}
