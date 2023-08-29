package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	ProofOfDamageRepository interface {
		CreateProofOfDamage(ctx context.Context, proofOfDamage entities.ProofOfDamage) (entities.ProofOfDamage, error)
	}

	proofOfDamageRepository struct {
		db *gorm.DB
	}
)

func NewProofOfDamageRepository(db *gorm.DB) *proofOfDamageRepository {
	return &proofOfDamageRepository{
		db: db,
	}
}

func (r *proofOfDamageRepository) CreateProofOfDamage(ctx context.Context, proofOfDamage entities.ProofOfDamage) (entities.ProofOfDamage, error) {
	if err := r.db.Create(&proofOfDamage).Error; err != nil {
		return entities.ProofOfDamage{}, err
	}

	return proofOfDamage, nil
}