package migrations

import (
	"errors"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListUserSeeder(db); err != nil {
		return err
	}

	if err := ListAdminSeeder(db); err != nil {
		return err
	}

	return nil
}

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
		}
	}

	return nil
}

func ListUserSeeder(db *gorm.DB) error {
	var listUser = []entities.User{
		{
			NamaLengkap: "Admin",
			Username:    "admin",
			NoTelp:      "081234567890",
			Email:       "admin@gmail.com",
			Password:    "admin123",
			Role:        helpers.ADMIN,
			IsVerified:  true,
		},
		{
			NamaLengkap: "User",
			Username:    "user",
			NoTelp:      "081234567891",
			Email:       "user@gmail.com",
			Password:    "user123",
			Role:        helpers.USER,
			IsVerified:  true,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.User{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var user entities.User
		err := db.Where(&entities.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
