package entities

import "github.com/google/uuid"

type (
	Pembayaran struct {
		ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		Name          string    `json:"name"`
		AccountNumber string    `json:"account_number"`
		BankName			string    `json:"bank_name"`
		TotalPrice    int       `json:"total_price"`
		PaymentUrl		string    `json:"payment_url"`

		CreateProjectUserId uuid.UUID         `gorm:"type:uuid;not null" json:"create_project_user_id"`
		CreateProjectUser   CreateProjectUser `gorm:"foreignkey:CreateProjectUserId" json:"create_project_user"`

		UserId uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
		User   User      `gorm:"foreignkey:UserId" json:"-"`
	}
)
