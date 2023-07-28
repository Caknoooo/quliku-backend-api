package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type ListBankRepository interface {
	GetAllListBank(ctx context.Context) ([]entities.ListBank, error)
	GetBankByID(ctx context.Context, bankID int) (entities.ListBank, error)
}

type listBankRepository struct {
	db *gorm.DB
}

func NewListBankRepository(db *gorm.DB) *listBankRepository {
	return &listBankRepository{db: db}
}

func (l *listBankRepository) GetAllListBank(ctx context.Context) ([]entities.ListBank, error) {
	var listBank []entities.ListBank
	if err := l.db.Find(&listBank).Error; err != nil {
		return nil, err
	}

	return listBank, nil
}

func (l *listBankRepository) GetBankByID(ctx context.Context, bankID int) (entities.ListBank, error) {
	var bank entities.ListBank
	if err := l.db.Where("id = ?", bankID).Take(&bank).Error; err != nil {
		return entities.ListBank{}, err
	}

	return bank, nil
}