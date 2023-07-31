package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MandorRepository interface {
	CreateMandor(ctx context.Context, mandor entities.Mandor) (entities.Mandor, error)
	GetMandorByMandorID(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error)
	GetMandorByUsername(ctx context.Context, username string) (entities.Mandor, error)
	GetMandorByEmail(ctx context.Context, email string) (entities.Mandor, error)
	GetAllMandor(ctx context.Context) ([]entities.Mandor, error)
	GetDetailMandor(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error)
	ChangeStatus(ctx context.Context, mandorDTO dto.ChangeStatusMandorRequest) (error)
	UpdateMandor(ctx context.Context, mandor entities.Mandor) (entities.Mandor, error)
}

type mandorRepository struct {
	db *gorm.DB
}

func NewMandorRepository(db *gorm.DB) *mandorRepository {
	return &mandorRepository{
		db: db,
	}
}

func (mr *mandorRepository) CreateMandor(ctx context.Context, mandor entities.Mandor) (entities.Mandor, error) {
	if err := mr.db.Create(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
}

func (mr *mandorRepository) GetAllMandor(ctx context.Context) ([]entities.Mandor, error) {
	var mandors []entities.Mandor

	if err := mr.db.Find(&mandors).Error; err != nil {
		return []entities.Mandor{}, err
	}

	return mandors, nil
}

func (mr *mandorRepository) GetDetailMandor(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error) {
	var mandor entities.Mandor

	if err := mr.db.Where("id = ?", mandorID).Take(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
}

func (mr *mandorRepository) ChangeStatus(ctx context.Context, mandorDTO dto.ChangeStatusMandorRequest) (error) {
	var mandor entities.Mandor
	if err := mr.db.Where("id = ?", mandorDTO.MandorID).Take(&mandor).Error; err != nil {
		return err
	}

	mandor.Status = mandorDTO.Status
	mandor.IsVerifiedAdmin = true
	if err := mr.db.Save(&mandor).Error; err != nil {
		return err
	}

	return nil
}

func (mr *mandorRepository) GetMandorByMandorID(ctx context.Context, mandorID uuid.UUID) (entities.Mandor, error) {
	var mandor entities.Mandor

	if err := mr.db.Where("id = ?", mandorID).Take(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
}

func (mr *mandorRepository) GetMandorByUsername(ctx context.Context, username string) (entities.Mandor, error) {
	var mandor entities.Mandor
	if err := mr.db.Where("username = ?", username).Take(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}
	return mandor, nil
}

func (mr *mandorRepository) GetMandorByEmail(ctx context.Context, email string) (entities.Mandor, error) {
	var mandor entities.Mandor
	if err := mr.db.Where("email = ?", email).Take(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
}

func (mr *mandorRepository) UpdateMandor(ctx context.Context, mandor entities.Mandor) (entities.Mandor, error) {
	if err := mr.db.Updates(&mandor).Error; err != nil {
		return entities.Mandor{}, err
	}

	return mandor, nil
} 