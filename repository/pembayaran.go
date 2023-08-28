package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	PembayaranRepository interface {
		Create(ctx context.Context, pembayaran entities.Pembayaran) (entities.Pembayaran, error)
		GetPembayaranById(ctx context.Context, pembayaranId string) (entities.Pembayaran, error)
		GetAllPembayaranByUserId(ctx context.Context, userId string) ([]entities.Pembayaran, error)
		GetAllPembayaran(ctx context.Context) ([]entities.Pembayaran, error)
	}

	pembayaranRepository struct {
		db *gorm.DB
	}
)

func NewPembayaranRepository(db *gorm.DB) PembayaranRepository {
	return &pembayaranRepository{
		db: db,
	}
}

func (pr *pembayaranRepository) Create(ctx context.Context, pembayaran entities.Pembayaran) (entities.Pembayaran, error) {
	if err := pr.db.Create(&pembayaran).Error; err != nil {
		return entities.Pembayaran{}, err
	}

	return pembayaran, nil
}

func (pr *pembayaranRepository) GetPembayaranById(ctx context.Context, pembayaranId string) (entities.Pembayaran, error) {
	var pembayaran entities.Pembayaran

	if err := pr.db.Where("id = ?", pembayaranId).Take(&pembayaran).Error; err != nil {
		return entities.Pembayaran{}, err
	}

	return pembayaran, nil
}

func (pr *pembayaranRepository) GetAllPembayaranByUserId(ctx context.Context, userId string) ([]entities.Pembayaran, error) {
	var pembayaran []entities.Pembayaran

	if err := pr.db.Where("user_id = ?", userId).Find(&pembayaran).Error; err != nil {
		return []entities.Pembayaran{}, err
	}

	return pembayaran, nil
}

func (pr *pembayaranRepository) GetAllPembayaran(ctx context.Context) ([]entities.Pembayaran, error) {
	var pembayaran []entities.Pembayaran

	if err := pr.db.Find(&pembayaran).Error; err != nil {
		return []entities.Pembayaran{}, err
	}

	return pembayaran, nil
}