package models

type Location struct {
	Id       int         `json:"id"`
	Location []string    `json:"locations"`
	Date     string `json:"date"`
}
