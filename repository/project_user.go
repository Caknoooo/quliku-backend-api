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
		ChangeStatusProjectUser(ctx context.Context, projectUserId string, status bool) (entities.CreateProjectUser, error)
		GetAllProjectByUserId(ctx context.Context, userId string) ([]entities.CreateProjectUser, error)
		GetDetailProjectUserById(ctx context.Context, projectId string, userId string) (entities.CreateProjectUser, error)
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

func (r *projectUserRepository) GetAllProjectByUserId(ctx context.Context, userId string) ([]entities.CreateProjectUser, error) {
	var projectUsers []entities.CreateProjectUser

	if err := r.db.Where("user_id = ?", userId).Find(&projectUsers).Error; err != nil {
		return []entities.CreateProjectUser{}, err
	}

	return projectUsers, nil
}

func (r *projectUserRepository) GetDetailProjectUserById(ctx context.Context, projectId string, userId string) (entities.CreateProjectUser, error) {
	var projectUser entities.CreateProjectUser

	if err := r.db.Where("id = ? AND user_id = ?", projectId, userId).Take(&projectUser).Error; err != nil {
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

func (r *projectUserRepository) ChangeStatusProjectUser(ctx context.Context, projectUserId string, status bool) (entities.CreateProjectUser, error) {
	var projectUser entities.CreateProjectUser

	if err := r.db.Model(&projectUser).Where("id = ?", projectUserId).Update("is_verified_admin", status).Error; err != nil {
		return entities.CreateProjectUser{}, err
	}

	return projectUser, nil
}