package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	DetailCategoryUserRepository interface {
		CreateDetailCategoryUser(ctx context.Context, detailCategoryUser entities.DetailCategory) (entities.DetailCategory, error)
		GetAllDetailCategoryUser(ctx context.Context) ([]entities.DetailCategory, error)
		GetDetailCategoryUserDetail(ctx context.Context, detailCategoryUserID string) (entities.DetailCategory, error)
	}

	detailCategoryUserRepository struct {
		db *gorm.DB
	}
)

func NewDetailCategoryUserRepository(db *gorm.DB) *detailCategoryUserRepository {
	return &detailCategoryUserRepository{db: db}
}

func (r *detailCategoryUserRepository) CreateDetailCategoryUser(ctx context.Context, detailCategoryUser entities.DetailCategory) (entities.DetailCategory, error) {
	if err := r.db.Create(&detailCategoryUser).Error; err != nil {
		return entities.DetailCategory{}, err
	}

	return detailCategoryUser, nil
}

func (r *detailCategoryUserRepository) GetAllDetailCategoryUser(ctx context.Context) ([]entities.DetailCategory, error) {
	var detailCategoryUsers []entities.DetailCategory

	if err := r.db.Find(&detailCategoryUsers).Error; err != nil {
		return []entities.DetailCategory{}, err
	}

	return detailCategoryUsers, nil
}

func (r *detailCategoryUserRepository) GetDetailCategoryUserDetail(ctx context.Context, detailCategoryUserID string) (entities.DetailCategory, error) {
	var detailCategoryUser entities.DetailCategory

	if err := r.db.Where("id = ?", detailCategoryUserID).Take(&detailCategoryUser).Error; err != nil {
		return entities.DetailCategory{}, err
	}

	return detailCategoryUser, nil
}
