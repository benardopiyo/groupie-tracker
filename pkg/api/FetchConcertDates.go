package api

import (
	"encoding/json"
	"fmt"
	"groupie-Tracker/pkg/models"
	"net/http"
)


func FetchConcertDates(id int) (models.ConcertDate, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error Failed to Fetch concert dates")
		return models.ConcertDate{}, fmt.Errorf("failed to fetch concert dates: %v", err)
	}
	defer resp.Body.Close()

	var concertDate models.ConcertDate
	if err := json.NewDecoder(resp.Body).Decode(&concertDate); err != nil {
		fmt.Println("Failed to decode concert dates")
		return models.ConcertDate{}, fmt.Errorf("failed to decode concert dates: %v", err)
	}
	return concertDate, nil
}
