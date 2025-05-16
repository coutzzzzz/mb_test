package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/coutzzzzz/mb-go-test/internal/domain"
	"github.com/coutzzzzz/mb-go-test/internal/service"
	"github.com/gorilla/mux"
)

type MMSController struct {
	mmsService *service.MMSService
}

func NewMMSController(mmsService *service.MMSService) *MMSController {
	return &MMSController{
		mmsService: mmsService,
	}
}

func (m *MMSController) GetMMS(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Must have a request body", http.StatusBadRequest)
		return
	}

	pair := mux.Vars(r)["pair"]
	if !isValidPairs(pair) {
		http.Error(w, "Request only accept BRLBTC or BRLETH", http.StatusBadRequest)
		return
	}

	var request domain.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error in request body", http.StatusBadRequest)
		return
	}

	if !isFromDateLessThanToDate(request.From, request.To) {
		http.Error(w, "From date must be less than To date", http.StatusBadRequest)
		return
	}

	if !isUntilToday(request.To) {
		http.Error(w, "The maximum day must be until today", http.StatusBadRequest)
		return
	}

	if !isWithinOneYear(request.From) {
		http.Error(w, "The minimum day must be 365 before today", http.StatusBadRequest)
		return
	}

	output, err := m.mmsService.GetMMS(request.From, request.To, pair, request.Range)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func isFromDateLessThanToDate(fromDate, toDate time.Time) bool {
	return fromDate.Before(toDate)
}

func isUntilToday(toDate time.Time) bool {
	today := time.Now().Truncate(24 * time.Hour)
	toDate = toDate.Truncate(24 * time.Hour)
	return toDate.Before(today) || toDate.Equal(today)
}

func isWithinOneYear(fromDate time.Time) bool {
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	return fromDate.After(oneYearAgo)
}

func isValidPairs(pair string) bool {
	return pair == "BRLBTC" || pair == "BRLETH"
}
