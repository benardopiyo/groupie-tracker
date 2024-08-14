package api

import (
	"encoding/json"
	"fmt"
	"groupie-Tracker/pkg/models"
	"net/http"
)

func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %v", err)
	}
	defer resp.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to decode artists: %v", err)
	}
	return artists, nil
}
