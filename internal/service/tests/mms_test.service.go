package tests

import (
	"fmt"
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"github.com/coutzzzzz/mb-go-test/internal/repository/mocks"
	"github.com/coutzzzzz/mb-go-test/internal/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MMSService", Label("MMS"), Ordered, func() {

	var repo = &mocks.MockMMSRepository{}
	var s = service.NewMMSService(repo)

	Context("Get MMS", func() {
		It("return error when get from database return error", func() {
			repo.GetFunc = func(from, to time.Time, pair string) ([]domain.MMS, error) {
				return nil, fmt.Errorf("error")
			}

			_, err := s.GetMMS(time.Now(), time.Now(), "", 0)
			Expect(err).ToNot(BeNil())
		})

		It("returns m20 when rangedays equals 20", func() {
			var value float64 = 10
			var rangeDay int = 20
			repo.GetFunc = func(from, to time.Time, pair string) ([]domain.MMS, error) {
				return []domain.MMS{
					{
						Mms20: value,
					},
				}, nil
			}

			resp, err := s.GetMMS(time.Now(), time.Now(), "", rangeDay)
			Expect(err).To(BeNil())
			Expect(resp[0].Mms).To(Equal(value))
		})

		It("returns m50 when rangedays equals 50", func() {
			var value float64 = 100
			var rangeDay int = 50
			repo.GetFunc = func(from, to time.Time, pair string) ([]domain.MMS, error) {
				return []domain.MMS{
					{
						Mms50: value,
					},
				}, nil
			}

			resp, err := s.GetMMS(time.Now(), time.Now(), "", rangeDay)
			Expect(err).To(BeNil())
			Expect(resp[0].Mms).To(Equal(value))
		})

		It("returns m200 when rangedays equals 200", func() {
			var value float64 = 100
			var rangeDay int = 200
			repo.GetFunc = func(from, to time.Time, pair string) ([]domain.MMS, error) {
				return []domain.MMS{
					{
						Mms200: value,
					},
				}, nil
			}

			resp, err := s.GetMMS(time.Now(), time.Now(), "", rangeDay)
			Expect(err).To(BeNil())
			Expect(resp[0].Mms).To(Equal(value))
		})

		It("returns unix timestamp from timedate of database", func() {
			var value float64 = 100
			var rangeDay int = 200
			var timeNow = time.Now()
			repo.GetFunc = func(from, to time.Time, pair string) ([]domain.MMS, error) {
				return []domain.MMS{
					{
						Timestamp: timeNow,
						Mms200:    value,
					},
				}, nil
			}

			resp, err := s.GetMMS(time.Now(), time.Now(), "", rangeDay)
			Expect(err).To(BeNil())
			Expect(resp[0].Timestamp).To(Equal(timeNow.Unix()))
		})
	})

})
