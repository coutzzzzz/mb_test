package mocks

import (
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
)

type MockMMSRepository struct {
	GetFunc        func(from, to time.Time, pair string) ([]domain.MMS, error)
	CreateManyFunc func(mms []domain.MMS) error
}

func (m MockMMSRepository) Get(from, to time.Time, pair string) ([]domain.MMS, error) {
	if m.GetFunc != nil {
		return m.GetFunc(from, to, pair)
	}
	return nil, nil
}

func (m MockMMSRepository) CreateMany(mms []domain.MMS) error {
	if m.CreateManyFunc(mms) != nil {
		return m.CreateManyFunc(mms)
	}

	return nil
}
