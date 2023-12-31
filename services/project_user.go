package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"

	// "github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/google/uuid"
)

const (
	PROJECT = "project"
)

type (
	ProjectUserService interface {
		CreateProjectUser(ctx context.Context, req dto.CreateProjectRequest, id uuid.UUID) (dto.CreateProjectResponse, error)
		GetAllProjectUserByAdmin(ctx context.Context, adminId uuid.UUID) ([]entities.CreateProjectUser, error)
		GetProjectUserByIdByAdmin(ctx context.Context, adminId uuid.UUID, projectId string) (entities.CreateProjectUser, error)
		ChangeStatusProjectUser(ctx context.Context, adminId uuid.UUID, req dto.ChangeStatusProjectUserRequest) (dto.CreateProjectResponse, error)
		GetAllProjectUserByUserId(ctx context.Context, userId uuid.UUID) ([]entities.CreateProjectUser, error)
		GetDetailProjectUserById(ctx context.Context, projectId string, userId uuid.UUID) (entities.CreateProjectUser, error)
	}

	projectUserService struct {
		ar repository.AdminRepository
		r   repository.ProjectUserRepository
		prf repository.ProofOfDamageRepository
		dtc repository.DetailCategoryUserRepository
		toc repository.TypeOfCraftsmanRepository
		ur repository.UserRepository
	}
)

func NewProjectUserService(ar repository.AdminRepository, r repository.ProjectUserRepository, prf repository.ProofOfDamageRepository, dtc repository.DetailCategoryUserRepository, toc repository.TypeOfCraftsmanRepository, ur repository.UserRepository) *projectUserService {
	return &projectUserService{
		ar: ar,
		r:   r,
		prf: prf,
		dtc: dtc,
		toc: toc,
		ur: ur,
	}
}

func (p *projectUserService) CreateProjectUser(ctx context.Context, req dto.CreateProjectRequest, userId uuid.UUID) (dto.CreateProjectResponse, error) {
	// var proofOfDamage []entities.ProofOfDamage
	// var detailCategory []entities.DetailCategory
	// var typeOfCraftsman []entities.TypeOfCraftsman
	var totalPrice int

	projectUserID := uuid.New()
	projectUser := entities.CreateProjectUser{
		ID:         projectUserID,
		Category:   req.Category,
		DetailDescription: req.DetailDescription,
		Alamat:     req.Alamat,
		DateStart:  req.DateStart,
		DateEnd:    req.DateEnd,
		Estimation: req.Estimation,
		UserID:     userId,
		TotalPrice: req.TotalPrice,
	}

	projectUserCreate, err := p.r.CreateProjectUser(ctx, projectUser)
	if err != nil {
		return dto.CreateProjectResponse{}, dto.ErrCreateProjectUser
	}

	for _, v := range req.DetailCategoryRequest {
		detailCategoryItem := entities.DetailCategory{
			ID:                  uuid.New(),
			Name:                v.Name,
			IsActive:            v.IsActive,
			CreateProjectUserID: projectUserID,
		}

		_, err := p.dtc.CreateDetailCategoryUser(ctx, detailCategoryItem)
		if err != nil {
			return dto.CreateProjectResponse{}, dto.ErrDetailCategory
		}
	}

	for _, v := range req.ProofOfDamageRequest {
		var imageName string

		if v.Photo != nil {
			projectPhoto, err := utils.IsBase64(*v.Photo)
			if err != nil {
				return dto.CreateProjectResponse{}, dto.ErrToBase64
			}

			imageId := uuid.New()

			projectPhotoSave := imageId.String() + utils.Getextension(v.Photo.Filename)

			_ = utils.SaveImage(projectPhoto, PATH, PROJECT, projectPhotoSave)

			imageName = utils.GenerateFileName(PATH, PROJECT, projectPhotoSave)
		}

		proofOfDamageItem := entities.ProofOfDamage{
			ID:                  uuid.New(),
			ImageUrl:            imageName,
			Filename:            v.Photo.Filename,
			CreateProjectUserID: projectUserID,
		}

		_, err := p.prf.CreateProofOfDamage(ctx, proofOfDamageItem)
		if err != nil {
			return dto.CreateProjectResponse{}, dto.ErrProofOfDamage
		}
	}

	for _, v := range req.TypeOfCraftsmanRequest {
		var typeOfCraftsmanItem entities.TypeOfCraftsman

		totalPrice += v.Price

		if v.IsHalfDay && !v.IsFullDay {
			typeOfCraftsmanItem = entities.TypeOfCraftsman{
				ID:                  uuid.New(),
				IsHalfDay:           true,
				Duration:            v.Duration,
				Price:               v.Price,
				Type:                v.Type,
				CreateProjectUserID: projectUserID,
			}
		} else {
			typeOfCraftsmanItem = entities.TypeOfCraftsman{
				ID:                  uuid.New(),
				IsFullDay:           true,
				Duration:            v.Duration,
				Price:               v.Price,
				Type:                v.Type,
				CreateProjectUserID: projectUserID,
			}
		}
		_, err := p.toc.CreateTypeOfCraftsman(ctx, typeOfCraftsmanItem)
		if err != nil {
			return dto.CreateProjectResponse{}, dto.ErrTypeOfCraftsman
		}
	}

	return dto.CreateProjectResponse{
		ID:              projectUserCreate.ID.String(),
		Category:        projectUserCreate.Category,
		Alamat:          projectUserCreate.Alamat,
		DateStart:       projectUserCreate.DateStart,
		DateEnd:         projectUserCreate.DateEnd,
		Estimation:      projectUserCreate.Estimation,
		TotalPrice:      projectUserCreate.TotalPrice,
		IsVerifiedAdmin: projectUserCreate.IsVerifiedAdmin,
	}, nil
}

func (p *projectUserService) GetAllProjectUserByUserId(ctx context.Context, userId uuid.UUID) ([]entities.CreateProjectUser, error) {
	user, err := p.ur.GetUserByID(ctx, userId)
	if err != nil {
		return []entities.CreateProjectUser{}, dto.ErrorUserNotFound
	}

	if user.Role != helpers.USER {
		return []entities.CreateProjectUser{}, dto.ErrRoleDontHaveAccess
	}

	projectUser, err := p.r.GetAllProjectByUserId(ctx, userId.String())
	if err != nil {
		return []entities.CreateProjectUser{}, dto.ErrGetAllProjectUser
	}

	return projectUser, nil
}

func (p *projectUserService) GetDetailProjectUserById(ctx context.Context, projectId string, userId uuid.UUID) (entities.CreateProjectUser, error) {
	user, err := p.ur.GetUserByID(ctx, userId)
	if err != nil {
		return entities.CreateProjectUser{}, dto.ErrorUserNotFound
	}

	if user.Role != helpers.USER {
		return entities.CreateProjectUser{}, dto.ErrRoleDontHaveAccess
	}

	projectUser, err := p.r.GetDetailProjectUserById(ctx, projectId, userId.String())
	if err != nil {
		return entities.CreateProjectUser{}, dto.ErrGetProjectUser
	}

	return projectUser, nil
}

func (p *projectUserService) GetAllProjectUserByAdmin(ctx context.Context, adminId uuid.UUID) ([]entities.CreateProjectUser, error) {
	admin, err := p.ar.GetAdminByID(ctx, adminId)
	if err != nil {
		return []entities.CreateProjectUser{}, dto.ErrorUserNotFound
	}

	if admin.Role != helpers.ADMIN {
		return []entities.CreateProjectUser{}, dto.ErrRoleDontHaveAccess
	}

	projectUser, err := p.r.GetAllProjectUser(ctx)
	if err != nil {
		return []entities.CreateProjectUser{}, dto.ErrGetAllProjectUser
	}

	return projectUser, nil
}

func (p *projectUserService) GetProjectUserByIdByAdmin(ctx context.Context, adminId uuid.UUID, projectId string) (entities.CreateProjectUser, error) {
	admin, err := p.ar.GetAdminByID(ctx, adminId)
	if err != nil {
		return entities.CreateProjectUser{}, dto.ErrAdminNotFound
	}

	if admin.Role != helpers.ADMIN {
		return entities.CreateProjectUser{}, dto.ErrRoleDontHaveAccess
	}

	projectUser, err := p.r.GetProjectUserDetail(ctx, projectId)
	if err != nil {
		return entities.CreateProjectUser{}, dto.ErrGetProjectUser
	}

	return projectUser, nil
} 

func (p *projectUserService) ChangeStatusProjectUser(ctx context.Context, adminId uuid.UUID, req dto.ChangeStatusProjectUserRequest) (dto.CreateProjectResponse, error) {
	admin, err := p.ar.GetAdminByID(ctx, adminId)
	if err != nil {
		return dto.CreateProjectResponse{}, dto.ErrAdminNotFound
	}

	if admin.Role != helpers.ADMIN {
		return dto.CreateProjectResponse{}, dto.ErrRoleDontHaveAccess
	}

	projectUser, err := p.r.GetProjectUserDetail(ctx, req.ProjectId)
	if err != nil {
		return dto.CreateProjectResponse{}, dto.ErrGetProjectUser
	}

	var isVerifiedAdmin bool

	if req.Status == ACCEPTED {
		isVerifiedAdmin = true
	} else if req.Status == REJECTED {
		isVerifiedAdmin = false
	} else {
		return dto.CreateProjectResponse{}, dto.ErrStatusIsNotValid
	}
	
	projectUserRes, err := p.r.ChangeStatusProjectUser(ctx, req.ProjectId, isVerifiedAdmin)
	if err != nil {
		return dto.CreateProjectResponse{}, dto.ErrChangeStatusProjectUser
	}

	return dto.CreateProjectResponse{
		ID:              projectUser.ID.String(),
		Category:        projectUser.Category,
		Alamat:          projectUser.Alamat,
		DateStart:       projectUser.DateStart,
		DateEnd:         projectUser.DateEnd,
		Estimation:      projectUser.Estimation,
		TotalPrice:      projectUser.TotalPrice,
		IsVerifiedAdmin: projectUserRes.IsVerifiedAdmin,
	}, nil
}