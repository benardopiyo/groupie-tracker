package models

type Relation struct {
	Id            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
