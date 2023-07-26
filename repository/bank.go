package repository

import (
	"context"

	"gorm.io/gorm"
)

type BankRepository interface {
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{
		db: db,
	}
}

func (br *bankRepository) CreateBankMandor(ctx context.Context, ) {

}