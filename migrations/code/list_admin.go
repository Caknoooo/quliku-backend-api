package code

import (
	"errors"
	"fmt"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ListAdminSeeder(db *gorm.DB) error {
	var listAdmin = []entities.Admin{
		{
			ID:         uuid.New(),
			Nama:       "Quliku Indonesia",
			Email:      "Qulikuindonesia@gmail.com",
			Password:   "Quliku5822",
			Role:       helpers.ADMIN,
			IsVerified: true,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.Admin{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.Admin{}); err != nil {
			return err
		}
	}

	for _, data := range listAdmin {
		var admin entities.Admin
		err := db.Where(&entities.Admin{Email: data.Email}).First(&admin).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&admin, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
			fmt.Println("Admin Created")
		}
	}

	return nil
}