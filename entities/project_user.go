package entities

import "github.com/google/uuid"

type (
	CreateProjectUser struct {
		ID                uuid.UUID        `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Category          string           `json:"category"`
		DetailDescription string           `json:"detail_description"`
		DetailCategory    []DetailCategory `json:"detail_category,omitempty"`
		Alamat            string           `json:"alamat"`

		ProofOfDamage []ProofOfDamage `json:"proof_of_damage,omitempty"`

		DateStart string `json:"date_start"`
		DateEnd   string `json:"date_end"`

		Estimation string `json:"estimation"`
		TotalPrice int    `json:"price"`

		TypeOfCraftsman []TypeOfCraftsman `json:"type_of_craftsman,omitempty"`

		IsVerifiedAdmin bool `gorm:"type:boolean;default:false" json:"is_verified_admin"`

		UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
		User   User      `gorm:"foreignKey:UserID" json:"-"`
	}

	ProofOfDamage struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		ImageUrl string    `json:"photo_url"`
		Filename string    `json:"filename"`

		CreateProjectUserID uuid.UUID         `gorm:"type:uuid" json:"create_project_user_id"`
		CreateProjectUser   CreateProjectUser `gorm:"foreignKey:CreateProjectUserID" json:"-"`
	}

	DetailCategory struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		Name     string    `json:"name"`
		IsActive bool      `json:"is_active"`

		CreateProjectUserID uuid.UUID         `gorm:"type:uuid" json:"create_project_user_id"`
		CreateProjectUser   CreateProjectUser `gorm:"foreignKey:CreateProjectUserID" json:"-"`
	}

	TypeOfCraftsman struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		IsHalfDay bool      `json:"is_half_day"`
		IsFullDay bool      `json:"is_full_day"`
		Duration  string    `json:"duration"`
		Price     int       `json:"price"`
		Type      string    `json:"type"`

		CreateProjectUserID uuid.UUID         `gorm:"type:uuid" json:"create_project_user_id"`
		CreateProjectUser   CreateProjectUser `gorm:"foreignKey:CreateProjectUserID" json:"-"`
	}
)
