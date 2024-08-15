package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-Tracker/pkg/models"
)

func FetchConcertDates() (models.ConcertDatesResponse, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error Failed to Fetch concert dates")
		return models.ConcertDatesResponse{}, fmt.Errorf("failed to fetch concert dates: %v", err)
	}
	defer resp.Body.Close()

	var concertDates models.ConcertDatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&concertDates); err != nil {
		fmt.Println("Failed to decode concert dates")
		return models.ConcertDatesResponse{}, fmt.Errorf("failed to decode concert dates: %v", err)
	}
	return concertDates, nil
}
