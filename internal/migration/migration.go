package migration

import (
	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.MMS{})
	if err != nil {
		panic("Failed to migrate table mms")
	}

	return nil
}
