package internal

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
)

var tmplt, err = template.ParseGlob("./templates/*.html")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" && r.URL.Path != "/concert_dates" && r.URL.Path != "/locations" && r.URL.Path != "/relations" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := api.FetchLocations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "location.html", locations)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	relations, err := api.FetchRelations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "relations.html", relations)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConcertDatesHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	concertDates, err := api.FetchConcertDates()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredDates := []models.ConcertDate{}
	for _, date := range concertDates.Index {
		if date.Id == artistId {
			filteredDates = append(filteredDates, date)
		}
	}

	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "concert_dates.html", filteredDates)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
