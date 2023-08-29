package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	TypeOfCraftsmanRepository interface {
		CreateTypeOfCraftsman(ctx context.Context, typeOfCraftsman entities.TypeOfCraftsman) (entities.TypeOfCraftsman, error)
		GetAllTypeOfCraftsman(ctx context.Context) ([]entities.TypeOfCraftsman, error)
		GetDetailTypeOfCraftsman(ctx context.Context, id string) (entities.TypeOfCraftsman, error)
	}

	typeOfCraftsmanRepository struct {
		db *gorm.DB
	}
)

func NewTypeOfCraftsmanRepository(db *gorm.DB) *typeOfCraftsmanRepository {
	return &typeOfCraftsmanRepository{
		db: db,
	}
}

func (r *typeOfCraftsmanRepository) CreateTypeOfCraftsman(ctx context.Context, typeOfCraftsman entities.TypeOfCraftsman) (entities.TypeOfCraftsman, error) {
	if err := r.db.Create(&typeOfCraftsman).Error; err != nil {
		return entities.TypeOfCraftsman{}, err
	}

	return typeOfCraftsman, nil
}

func (r *typeOfCraftsmanRepository) GetAllTypeOfCraftsman(ctx context.Context) ([]entities.TypeOfCraftsman, error) {
	var typeOfCraftsmans []entities.TypeOfCraftsman

	if err := r.db.Find(&typeOfCraftsmans).Error; err != nil {
		return []entities.TypeOfCraftsman{}, err
	}

	return typeOfCraftsmans, nil
}

func (r *typeOfCraftsmanRepository) GetDetailTypeOfCraftsman(ctx context.Context, id string) (entities.TypeOfCraftsman, error) {
	var typeOfCraftsman entities.TypeOfCraftsman

	if err := r.db.Where("id = ?", id).Take(&typeOfCraftsman).Error; err != nil {
		return entities.TypeOfCraftsman{}, err
	}

	return typeOfCraftsman, nil
}