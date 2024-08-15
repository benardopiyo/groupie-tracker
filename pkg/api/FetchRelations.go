package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-Tracker/pkg/models"
)

func FetchRelations() ([]models.Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relations: %v", err)
	}
	defer resp.Body.Close()

	var response models.RelationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode relations: %v", err)
	}
	return response.Index, nil
}
