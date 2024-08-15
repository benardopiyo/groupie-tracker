package models

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Member       []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relation     string   `json:"relations"`
}
