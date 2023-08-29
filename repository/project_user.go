package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	ProjectUserRepository interface {
		CreateProjectUser(ctx context.Context, projectUser entities.CreateProjectUser) (entities.CreateProjectUser, error)
		GetAllProjectUser(ctx context.Context) ([]entities.CreateProjectUser, error)
		GetProjectUserDetail(ctx context.Context, projectUserID string) (entities.CreateProjectUser, error)
	}

	projectUserRepository struct {
		db *gorm.DB
	}
)

func NewProjectUserRepository(db *gorm.DB) *projectUserRepository {
	return &projectUserRepository{db: db}
}

func (r *projectUserRepository) CreateProjectUser(ctx context.Context, projectUser entities.CreateProjectUser) (entities.CreateProjectUser, error) {
	if err := r.db.Create(&projectUser).Error; err != nil {
		return entities.CreateProjectUser{}, err
	}

	return projectUser, nil
}

func (r *projectUserRepository) GetAllProjectUser(ctx context.Context) ([]entities.CreateProjectUser, error) {
	var projectUsers []entities.CreateProjectUser

	if err := r.db.Preload("DetailCategory").Preload("ProofOfDamage").Preload("TypeOfCraftsman").Find(&projectUsers).Error; err != nil {
		return []entities.CreateProjectUser{}, err
	}

	return projectUsers, nil
}

func (r *projectUserRepository) GetProjectUserDetail(ctx context.Context, projectUserID string) (entities.CreateProjectUser, error) {
	var projectUser entities.CreateProjectUser

	if err := r.db.Preload("DetailCategory").Preload("ProofOfDamage").Preload("TypeOfCraftsman").Where("id = ?", projectUserID).Take(&projectUser).Error; err != nil {
		return entities.CreateProjectUser{}, err
	}

	return projectUser, nil
}