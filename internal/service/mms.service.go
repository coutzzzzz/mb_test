package service

import (
	"fmt"
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"github.com/coutzzzzz/mb-go-test/internal/repository"
)

type MMSService struct {
	mmsRepository repository.IMMSRepository
}

func NewMMSService(MMSRepository repository.IMMSRepository) *MMSService {
	return &MMSService{
		mmsRepository: MMSRepository,
	}
}

func (m MMSService) GetMMS(from, to time.Time, pair string, rangeDays int) ([]domain.Response, error) {
	mmss, err := m.mmsRepository.Get(from.AddDate(0, 0, -1), to, pair)
	if err != nil {
		return nil, fmt.Errorf("error getting mms from repository: %v", err)
	}

	var output []domain.Response
	for _, m := range mmss {
		var value float64
		switch rangeDays {
		case 20:
			value = m.Mms20
		case 50:
			value = m.Mms50
		case 200:
			value = m.Mms200
		}

		output = append(output, domain.Response{
			Timestamp: m.Timestamp.Unix(),
			Mms:       value,
		})
	}

	return output, nil
}
