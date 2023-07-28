package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdminByEmail(ctx context.Context, email string) (entities.Admin, error)
	GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (ar *adminRepository) GetAdminByEmail(ctx  context.Context, email string) (entities.Admin, error) {
	var admin entities.Admin
	if err := ar.db.Where("email = ?", email).Take(&admin).Error; err != nil {
		return entities.Admin{}, err
	}

	return admin, nil
}

func (ar *adminRepository) GetAdminByID(ctx context.Context, adminID uuid.UUID) (entities.Admin, error) {
	var admin entities.Admin
	if err := ar.db.Where("id = ?", adminID).Take(&admin).Error; err != nil {
		return entities.Admin{}, err
	}

	return admin, nil
}
