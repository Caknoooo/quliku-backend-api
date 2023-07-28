package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
)

type AdminService interface {
	RegisterAdmin()
	VerifyLogin(ctx context.Context, adminDTO dto.AdminLoginDTO) (bool, error)
	CheckAdminByEmail(ctx context.Context, email string) (entities.Admin, error)
	GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.Admin, error)
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(ar repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository: ar,
	}
}

func (as *adminService) RegisterAdmin() {
	
}

func (as *adminService) CheckAdminByEmail(ctx context.Context, email string) (entities.Admin, error) {
	admin, err := as.adminRepository.GetAdminByEmail(ctx, email)
	if err != nil {
		return entities.Admin{}, err
	}

	return admin, nil
}

func (as *adminService) GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.Admin, error) {
	admin, err := as.adminRepository.GetAdminByID(ctx, adminID)
	if err != nil {
		return entities.Admin{}, err
	}

	return admin, nil
}

func (as *adminService) VerifyLogin(ctx context.Context, adminDTO dto.AdminLoginDTO) (bool, error) {
	admin, err := as.adminRepository.GetAdminByEmail(ctx, adminDTO.Email)
	if err != nil {
		return false, dto.ErrEmailNotFound
	}

	if admin.Email == "" {
		return false, dto.ErrEmailNotFound
	}

	if checkPassword, _ := helpers.CheckPassword(admin.Password, []byte(adminDTO.Password)); !checkPassword {
		return false, dto.ErrPasswordNotMatch
	}

	return true, nil
}