package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type (
	ListBankService interface {
		GetAllListBank(ctx context.Context) ([]entities.ListBank, error)
		GetBankByID(ctx context.Context, bankID int) (entities.ListBank, error)
	}

	listBankService struct {
		listBankRepository repository.ListBankRepository
	}
)

func NewListBankService(listBankRepository repository.ListBankRepository) ListBankService {
	return &listBankService{
		listBankRepository: listBankRepository,
	}
}

func (l *listBankService) GetAllListBank(ctx context.Context) ([]entities.ListBank, error) {
	listBank, err := l.listBankRepository.GetAllListBank(ctx)
	if err != nil {
		return nil, err
	}

	return listBank, nil
}

func (l *listBankService) GetBankByID(ctx context.Context, bankID int) (entities.ListBank, error) {
	bank, err := l.listBankRepository.GetBankByID(ctx, bankID)
	if err != nil {
		return entities.ListBank{}, err
	}

	return bank, nil
}