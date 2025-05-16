package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(dsn string) *gorm.DB {
	driver := postgres.Open(dsn)

	client, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		log.Panicf("failed to initialize database: %v", err)
	}

	return client
}
