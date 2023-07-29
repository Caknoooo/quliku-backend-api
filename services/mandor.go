package services

import (
	"context"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	PATH = "storage/images"
	PROFILE = "profile"
	KTP = "ktp"
	SERTIFIKAT = "sertifikat"
	PORTOFOLIO = "portofolio"

	// Status Mandor
	REJECTED = "rejected"
	WAITING = "waiting"
	ACCEPTED = "accepted"
)

type MandorService interface {
	RegisterMandorStart(ctx context.Context, mandorDTO dto.MandorCreateDTO) (dto.MandorCreateDTO, error)
	RegisterMandorEnd(ctx context.Context, mandorDTO dto.MandorNextDTO) (entities.Mandor, error)
	VerifyLogin(ctx context.Context, mandorDTO dto.MandorLoginDTO) (bool, error)
	CheckMandorByEmail(ctx context.Context, emailz string) (entities.Mandor, error)
	GetMandorByMandorID(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error)
	GetMandorByEmail(ctx context.Context, emailz string) (bool, error)
	VerifyEmail(ctx context.Context, mandorVerificationDTO dto.MandorVerificationDTO) (bool, error)
	ResendVerificationCode(ctx context.Context, mandorVerificationDTO dto.ResendMandorVerificationCodeDTO) (bool, error)
	ResendFailedLoginNotVerified(ctx context.Context, email string) (bool, error)
}

type mandorService struct {
	mandorRepository             repository.MandorRepository
	mandorVerificationRepository repository.MandorVerificationRepository
}

func NewMandorService(mr repository.MandorRepository, mvr repository.MandorVerificationRepository) MandorService {
	return &mandorService{
		mandorRepository:             mr,
		mandorVerificationRepository: mvr,
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
	
	mandor.ID = uuid.New()
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
	mandor.Role = helpers.MANDOR
	
	// Unggah Dokumen
	// Opsional
	if mandorDTO.FotoProfil != nil {
		fotoProfil, err := utils.IsBase64(*mandorDTO.FotoProfil)
		if err != nil {
			return entities.Mandor{}, dto.ErrToBase64
		}

		fotoProfilSave := mandor.ID.String() + utils.Getextension(mandorDTO.FotoProfil.Filename)

		_ = utils.SaveImage(fotoProfil, PATH, PROFILE, fotoProfilSave)

		mandor.FotoProfil = utils.GenerateFileName(PATH, PROFILE, fotoProfilSave)
	}

	// Foto Sertifikat
	if mandorDTO.FotoSertifikat != nil {
		fotoSertifikat, err := utils.IsBase64(*mandorDTO.FotoSertifikat)
		if err != nil {
			return entities.Mandor{}, dto.ErrToBase64
		}

		fotoSertifikatSave := mandor.ID.String() + utils.Getextension(mandorDTO.FotoSertifikat.Filename)

		_ = utils.SaveImage(fotoSertifikat, PATH, SERTIFIKAT, fotoSertifikatSave)

		mandor.FotoSertifikat = utils.GenerateFileName(PATH, SERTIFIKAT, fotoSertifikatSave)
	}

	fotoKTP, err := utils.IsBase64(*mandorDTO.FotoKTP)
	if err != nil {
		return entities.Mandor{}, dto.ErrToBase64
	}

	fotoPortofolio, err := utils.IsBase64(*mandorDTO.FotoPortofolio)
	if err != nil {
		return entities.Mandor{}, dto.ErrToBase64
	}

	// KTP
	fotoKTPSave := mandor.ID.String() + utils.Getextension(mandorDTO.FotoKTP.Filename)

	_ = utils.SaveImage(fotoKTP, PATH, KTP, fotoKTPSave)

	mandor.FotoKTP = utils.GenerateFileName(PATH, KTP, fotoKTPSave)

	// Portfolio
	fotoPortofolioSave := mandor.ID.String() + utils.Getextension(mandorDTO.FotoPortofolio.Filename)

	_ = utils.SaveImage(fotoPortofolio, PATH, PORTOFOLIO, fotoPortofolioSave)

	mandor.FotoPortofolio = utils.GenerateFileName(PATH, PORTOFOLIO, fotoPortofolioSave)

	// Data Bank
	mandor.Banks = []entities.Bank{
		{
			ID:         uuid.New(),
			NamaBank:   mandorDTO.NamaBank,
			NoRekening: mandorDTO.NoRekening,
			AtasNama:   mandorDTO.AtasNama,
		},
	}

	mandorData, err := ms.mandorRepository.CreateMandor(ctx, mandor)
	if err != nil {
		return entities.Mandor{}, err
	}

	draftEmail, err := utils.MakeVerificationEmail(mandor.Email)
	if err != nil {
		return entities.Mandor{}, err
	}

	_ = ms.mandorVerificationRepository.Create(mandorData.ID, draftEmail["code"], time.Now().Add(time.Minute*3))

	err = utils.SendMail(mandor.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return entities.Mandor{}, err
	}

	return mandorData, nil
}

func (ms *mandorService) VerifyLogin(ctx context.Context, mandorDTO dto.MandorLoginDTO) (bool, error) {
	email, err := ms.mandorRepository.GetMandorByEmail(ctx, mandorDTO.Email)
	if err != nil {
		return false, err
	}

	if email.Email == "" {
		return false, nil
	}

	if checkPassword, err := helpers.CheckPassword(email.Password, []byte(mandorDTO.Password)); !checkPassword {
		return false, err
	}

	return true, nil
}

func (ms *mandorService) GetMandorByMandorID(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error) {
	mandor, err := ms.mandorRepository.GetMandorByMandorID(ctx, mandorID)
	if err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
}

func (ms *mandorService) CheckMandorByEmail(ctx context.Context, emailz string) (entities.Mandor, error) {
	email, err := ms.mandorRepository.GetMandorByEmail(ctx, emailz)
	if err != nil {
		return entities.Mandor{}, err
	}

	return email, nil
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

func (ms *mandorService) VerifyEmail(ctx context.Context, mandorVerificationDTO dto.MandorVerificationDTO) (bool, error) {
	mandorVerification, err := ms.mandorVerificationRepository.Check(mandorVerificationDTO.MandorID)
	if err != nil {
		return false, err
	}

	if mandorVerification.ReceiveCode != mandorVerificationDTO.SendCode {
		return false, dto.ErrorVerificationCodeNotMatch
	}

	if mandorVerification.ExpiredAt.Before(time.Now()) {
		return false, dto.ErrorExpiredVerificationCode
	}

	if mandorVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	if err := ms.mandorVerificationRepository.SendCode(mandorVerificationDTO.MandorID, mandorVerificationDTO.SendCode); err != nil {
		return false, err
	}

	return true, nil
}

func (ms *mandorService) ResendVerificationCode(ctx context.Context, mandorVerificationDTO dto.ResendMandorVerificationCodeDTO) (bool, error) {
	mandorVerification, err := ms.mandorVerificationRepository.Check(mandorVerificationDTO.MandorID)
	if err != nil {
		return false, err
	}

	if mandorVerification.ExpiredAt.After(time.Now()) {
		return false, dto.ErrorNotExpiredVerificationCode
	}

	if mandorVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	mandor, err := ms.mandorRepository.GetMandorByMandorID(ctx, mandorVerificationDTO.MandorID)
	if err != nil {
		return false, err
	}

	if mandor.IsVerified {
		return false, dto.ErrorUserAlreadyActive
	}

	draftEmail, err := utils.MakeVerificationEmail(mandor.Email)
	if err != nil {
		return false, err
	}

	_ = ms.mandorVerificationRepository.Delete(mandorVerification.MandorID)

	_ = ms.mandorVerificationRepository.Create(mandorVerification.MandorID, draftEmail["code"], time.Now().Add(time.Minute*3))

	err = utils.SendMail(mandor.Email, draftEmail["subject"], draftEmail["body"])

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ms *mandorService) ResendFailedLoginNotVerified(ctx context.Context, email string) (bool, error) {
	mail, err := ms.mandorRepository.GetMandorByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if mail.IsVerified {
		return false, dto.ErrorUserAlreadyActive
	}

	mandorVerification, err := ms.mandorVerificationRepository.Check(mail.ID)
	if err != nil {
		return false, err
	}

	if mandorVerification.ExpiredAt.After(time.Now()) {
		return false, dto.ErrorNotExpiredVerificationCode
	}

	if mandorVerification.IsActive {
		return false, dto.ErrorUserVerificationCodeAlreadyUsed
	}

	draftEmail, err := utils.MakeVerificationEmail(mail.Email)
	if err != nil {
		return false, err
	}

	_ = ms.mandorVerificationRepository.Delete(mail.ID)

	_ = ms.mandorVerificationRepository.Create(mail.ID, draftEmail["code"], time.Now().Add(time.Minute*3))

	err = utils.SendMail(mail.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return false, err
	}

	return true, nil
}
