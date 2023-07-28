package migrations

import (
	"github.com/Caknoooo/golang-clean_template/migrations/code"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := SeederCode(db); err != nil {
		return err
	}

	if err := SeederCSV(db); err != nil {
		return err
	}
	
	return nil
}

func SeederCode(db *gorm.DB) error {
	if err := code.ListUserSeeder(db); err != nil {
		return err
	}

	if err := code.ListAdminSeeder(db); err != nil {
		return err
	}

	return nil
}

func SeederCSV(db *gorm.DB) error {
	if err := Listbank(db); err != nil {
		return err
	}

	return nil
}

func SeederJSON(db *gorm.DB) error {
	return nil
}