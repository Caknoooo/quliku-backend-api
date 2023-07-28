package migrations

import (
	"os"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/gocarina/gocsv"
	"gorm.io/gorm"
)

const (
	PATH = "migrations/csv/"
)

type (
	ListBank struct {
		ID   int    `csv:"id"`
		Nama string `csv:"nama"`
	}
)

func Listbank(db *gorm.DB) error {
	file, err := os.Open(PATH + "list_bank.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	var listBanks []ListBank
	if err := gocsv.Unmarshal(file, &listBanks); err != nil {
		return err
	}

	for _, data := range listBanks {
		listBank := entities.ListBank{
			ID:   data.ID,
			Nama: data.Nama,
		}

		var bank entities.ListBank
		isData := db.Find(&bank, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&listBank).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
