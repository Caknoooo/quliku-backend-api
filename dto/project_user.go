package dto

import (
	"errors"
	"mime/multipart"
)

var (
	ErrCreateProjectUser       = errors.New("failed to create project user")
	ErrDetailCategory          = errors.New("failed to create detail category")
	ErrProofOfDamage           = errors.New("failed to create proof of damage")
	ErrTypeOfCraftsman         = errors.New("failed to create type of craftsman")
	ErrGetAllProjectUser       = errors.New("failed to get all project user")
	ErrGetProjectUser          = errors.New("failed to get project user")
	ErrChangeStatusProjectUser = errors.New("failed to change status project user")
	ErrStatusIsNotValid        = errors.New("status is not valid")
	ErrProjectUserIsVerified   = errors.New("project user is verified")
)

type (
	CreateProjectRequest struct {
		Category              string                  `json:"category"`
		DetailDescription     string                  `json:"detail_description"`
		DetailCategoryRequest []DetailCategoryRequest `json:"detail_category"`
		Alamat                string                  `json:"alamat"`

		ProofOfDamageRequest []ProofOfDamageRequest `json:"proof_of_damage"`

		DateStart string `json:"date_start"`
		DateEnd   string `json:"date_end"`

		Estimation string `json:"estimation"`

		TypeOfCraftsmanRequest []TypeOfCraftsmanRequest `json:"type_of_craftsman"`
		UserId                 string                   `json:"user_id"`

		IsAccepted bool `json:"is_accepted"`
		TotalPrice int  `json:"total_price"`
	}

	CreateProjectResponse struct {
		ID              string `json:"id"`
		Category        string `json:"category"`
		Alamat          string `json:"alamat"`
		DateStart       string `json:"date_start"`
		DateEnd         string `json:"date_end"`
		Estimation      string `json:"estimation"`
		TotalPrice      int    `json:"price"`
		IsVerifiedAdmin bool   `json:"is_verified_admin"`
	}

	ChangeStatusProjectUserRequest struct {
		ProjectId string `json:"project_id"`
		Status    string `json:"status"`
	}

	DetailCategoryResponse struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		IsActive bool   `json:"is_active"`
	}

	DetailCategoryRequest struct {
		Name     string `json:"name"`
		IsActive bool   `json:"is_active"`
	}

	ProofOfDamageRequest struct {
		Photo *multipart.FileHeader `json:"photo"`
	}

	ProofOfDamageResponse struct {
		ID       string `json:"id"`
		Photo    string `json:"photo"`
		Filename string `json:"filename"`
		Path     string `json:"path"`
	}

	TypeOfCraftsmanRequest struct {
		IsHalfDay bool   `json:"is_half_day"`
		IsFullDay bool   `json:"is_full_day"`
		Duration  string `json:"duration"`
		Price     int    `json:"price"`
		Type      string `json:"type"`
	}

	TypeOfCraftsmanResponse struct {
		ID        string `json:"id"`
		IsHalfDay bool   `json:"is_half_day"`
		IsFullDay bool   `json:"is_full_day"`
		Duration  string `json:"duration"`
		Price     string `json:"price"`
	}
)
