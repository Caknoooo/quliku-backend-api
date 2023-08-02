package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	MandorRatingRepository interface {
		CreateRatingMandor(ctx context.Context, mandorRating entities.MandorRating) (entities.MandorRating, error)
		GetAllMandorRatingByMandorID(ctx context.Context, mandorID uuid.UUID) ([]entities.MandorRating, error)
	}

	mandorRatingRepository struct {
		db *gorm.DB
	}
)

func NewMandorRatingRepository(db *gorm.DB) MandorRatingRepository {
	return &mandorRatingRepository{
		db: db,
	}
}

func (mr *mandorRatingRepository) CreateRatingMandor(ctx context.Context, mandorRating entities.MandorRating) (entities.MandorRating, error) {
	if err := mr.db.Create(&mandorRating).Error; err != nil {
		return entities.MandorRating{}, err
	}

	return mandorRating, nil
}

func (mr *mandorRatingRepository) GetAllMandorRatingByMandorID(ctx context.Context, mandorID uuid.UUID) ([]entities.MandorRating, error) {
	var mandorRatings []entities.MandorRating
	if err := mr.db.Where("mandor_id = ?", mandorID).Find(&mandorRatings).Error; err != nil {
		return nil, err
	}

	return mandorRatings, nil
}