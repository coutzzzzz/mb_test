package repository

import (
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"gorm.io/gorm"
)

type IMMSRepository interface {
	CreateMany([]domain.MMS) error
	Get(time.Time, time.Time, string) ([]domain.MMS, error)
}

type MMSRepository struct {
	db *gorm.DB
}

func NewMMSRepository(db *gorm.DB) *MMSRepository {
	return &MMSRepository{
		db: db,
	}
}

func (r MMSRepository) CreateMany(mms []domain.MMS) error {
	result := r.db.CreateInBatches(mms, len(mms))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r MMSRepository) Get(from, to time.Time, pair string) ([]domain.MMS, error) {
	var mmss []domain.MMS

	result := r.db.Where("pair = ? AND timestamp BETWEEN ? AND ?", pair, from, to).Find(&mmss)

	return mmss, result.Error
}
