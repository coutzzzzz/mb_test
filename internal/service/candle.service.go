package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"github.com/coutzzzzz/mb-go-test/internal/repository"
)

type Response struct {
	T []int64
	C []string
}

type CandleService struct {
	mmsRepository *repository.MMSRepository
}

func NewCandleService(mmsRepository *repository.MMSRepository) *CandleService {
	return &CandleService{
		mmsRepository: mmsRepository,
	}
}

func (c CandleService) Run(leftCoin, rightCoin string) error {
	baseUrl := os.Getenv("BASE_URL")
	url := fmt.Sprintf(baseUrl+"candles?symbol=%v-%v&from=1684159716&to=%v&resolution=1d", leftCoin, rightCoin, time.Now().Unix())

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting data from request: %v", err)
	}

	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding request body: %v", err)
	}

	current20, current50, current200 := 0.0, 0.0, 0.0
	current20ToInsert, current50ToInsert, current200ToInsert := 0.0, 0.0, 0.0

	var mmsToInsert []domain.MMS
	closes := convertStrToFloat(response.C)
	for idx, close := range closes {
		current20 += close
		current50 += close
		current200 += close

		if idx >= 20 {
			current20ToInsert = current20 / 20
			current20 -= closes[idx-20]
		}

		if idx >= 50 {
			current50ToInsert = current50 / 50
			current50 -= closes[idx-50]
		}

		if idx >= 200 {
			current200ToInsert = current200 / 200
			current200 -= closes[idx-200]
		}

		mmsToInsert = append(mmsToInsert, domain.MMS{
			Pair:      fmt.Sprintf("%v%v", rightCoin, leftCoin),
			Timestamp: time.Unix(response.T[idx], 0),
			Mms20:     current20ToInsert,
			Mms50:     current50ToInsert,
			Mms200:    current200ToInsert,
		})
	}

	if err = c.mmsRepository.CreateMany(mmsToInsert); err != nil {
		return fmt.Errorf("error creating batch of mms: %v", err)
	}

	return nil
}

func convertStrToFloat(strCloses []string) []float64 {
	var closes []float64
	for _, close := range strCloses {
		f, err := strconv.ParseFloat(close, 64)
		if err != nil {
			return []float64{}
		}

		closes = append(closes, f)
	}

	return closes
}
