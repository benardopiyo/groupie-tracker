package api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"groupie-Tracker/pkg/models"
)

type RelationsResponse struct {
    Index []models.Relation `json:"index"`
}

func FetchRelations() ([]models.Relation, error) {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil {
        return nil, fmt.Errorf("failed to fetch relations: %v", err)
    }
    defer resp.Body.Close()

    var response RelationsResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return nil, fmt.Errorf("failed to decode relations: %v", err)
    }

    return response.Index, nil
}
