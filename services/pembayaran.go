package services

import (
	"context"
	// "fmt"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	PAYMENT = "payment"
)

type (
	PembayaranService interface {
		Create(ctx context.Context, req dto.PembayaranRequest, projectId string, userId uuid.UUID) (dto.PembayaranResponse, error)
		GetPembayaranById(ctx context.Context, adminId uuid.UUID, pembayaranId string) (dto.PembayaranResponse, error)
		GetAllPembayaranByUserId(ctx context.Context, userId string) ([]dto.PembayaranResponse, error)
		GetAllPembayaran(ctx context.Context, adminId uuid.UUID) ([]entities.Pembayaran, error)
	}

	pembayaranService struct {
		pr repository.PembayaranRepository
		ur repository.UserRepository
		pu repository.ProjectUserRepository
		ar repository.AdminRepository
	}
)

func NewPembayaranService(pr repository.PembayaranRepository, ur repository.UserRepository, pu repository.ProjectUserRepository, ar repository.AdminRepository) PembayaranService {
	return &pembayaranService{
		pr: pr,
		ur: ur,
		pu: pu,
		ar: ar,
	}
}

func (ps *pembayaranService) Create(ctx context.Context, req dto.PembayaranRequest, projectId string, userId uuid.UUID) (dto.PembayaranResponse, error) {
	project, err := ps.pu.GetProjectUserDetail(ctx, projectId)
	if err != nil {
		return dto.PembayaranResponse{}, err
	}	
	// fmt.Println(projectId)
	// fmt.Println(userId)
	user, err := ps.ur.GetUserByID(ctx, userId)
	if err != nil {
		return dto.PembayaranResponse{}, dto.ErrorUserNotFound
	}

	var imageName string
	if req.PaymentPhoto != nil {
		paymentPhoto, err := utils.IsBase64(*req.PaymentPhoto)
		if err != nil {
			return dto.PembayaranResponse{}, dto.ErrToBase64
		}

		imageId := uuid.New()

		paymentPhotoSave := imageId.String() + utils.Getextension(req.PaymentPhoto.Filename)

		_ = utils.SaveImage(paymentPhoto, PATH, PAYMENT, paymentPhotoSave)

		imageName = utils.GenerateFileName(PATH, PAYMENT, paymentPhotoSave)
	} else {
		imageName = ""
	}

	pembayaran := entities.Pembayaran{
		Name:                req.Name,
		AccountNumber:       req.AccountNumber,
		TotalPrice:          req.TotalPrice,
		BankName:            req.BankName,
		PaymentUrl:          imageName,
		CreateProjectUserId: project.ID,
		UserId:              user.ID,
	}

	
	pembayaranCreate, err := ps.pr.Create(ctx, pembayaran)
	if err != nil {
		return dto.PembayaranResponse{}, dto.ErrCreatePembayaran
	}

	return dto.PembayaranResponse{
		ID:            pembayaranCreate.ID.String(),
		Name:          pembayaranCreate.Name,
		AccountNumber: pembayaranCreate.AccountNumber,
		BankName:      pembayaranCreate.BankName,
		TotalPrice:    pembayaranCreate.TotalPrice,
		PaymentUrl:    pembayaranCreate.PaymentUrl,
	}, nil
}

func (ps *pembayaranService) GetPembayaranById(ctx context.Context, adminId uuid.UUID, pembayaranId string) (dto.PembayaranResponse, error) {
	admin, err := ps.ar.GetAdminByID(ctx, adminId)
	if err != nil {
		return dto.PembayaranResponse{}, dto.ErrAdminNotFound
	}

	if admin.Role != helpers.ADMIN {
		return dto.PembayaranResponse{}, dto.ErrNotAdminID
	}

	pembayaran, err := ps.pr.GetPembayaranById(ctx, pembayaranId)
	if err != nil {
		return dto.PembayaranResponse{}, dto.ErrPembayaranNotFound
	}

	return dto.PembayaranResponse{
		ID:            pembayaran.ID.String(),
		Name:          pembayaran.Name,
		AccountNumber: pembayaran.AccountNumber,
		TotalPrice:    pembayaran.TotalPrice,
		PaymentUrl:    pembayaran.PaymentUrl,
	}, nil
}

func (ps *pembayaranService) GetAllPembayaranByUserId(ctx context.Context, userId string) ([]dto.PembayaranResponse, error) {
	pembayaran, err := ps.pr.GetAllPembayaranByUserId(ctx, userId)
	if err != nil {
		return []dto.PembayaranResponse{}, dto.ErrPembayaranNotFound
	}

	var pembayaranResponse []dto.PembayaranResponse
	for _, v := range pembayaran {
		pembayaranResponse = append(pembayaranResponse, dto.PembayaranResponse{
			ID:            v.ID.String(),
			Name:          v.Name,
			AccountNumber: v.AccountNumber,
			TotalPrice:    v.TotalPrice,
			PaymentUrl:    v.PaymentUrl,
		})
	}

	return pembayaranResponse, nil
}

func (ps *pembayaranService) GetAllPembayaran(ctx context.Context, adminId uuid.UUID) ([]entities.Pembayaran, error) {
	admin, err := ps.ar.GetAdminByID(ctx, adminId)
	if err != nil {
		return []entities.Pembayaran{}, dto.ErrAdminNotFound
	}

	if admin.Role != helpers.ADMIN {
		return []entities.Pembayaran{}, dto.ErrNotAdminID
	}

	pembayaran, err := ps.pr.GetAllPembayaran(ctx)
	if err != nil {
		return []entities.Pembayaran{}, dto.ErrPembayaranNotFound
	}

	return pembayaran, nil
}
