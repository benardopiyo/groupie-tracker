package internal

import (
	"fmt"
	"groupie-Tracker/pkg/api"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var tmplt, err = template.ParseGlob("./templates/*.html")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if r.URL.Path != "/" && r.URL.Path != "/concert_dates" && r.URL.Path != "/locations" && r.URL.Path != "/relations" {
	// 	ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }

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
	err = tmplt.ExecuteTemplate(w, "locations.html", locations)
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

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/concert_dates/"))
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	concertDates, err := api.FetchConcertDates(id)
	fmt.Println(api.FetchConcertDates(id))
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(concertDates)
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "concert_dates.html", concertDates)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
