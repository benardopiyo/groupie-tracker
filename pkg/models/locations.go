package models

type LocationWrapper struct {
	Index []Location `json:"index"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"date"`
}
