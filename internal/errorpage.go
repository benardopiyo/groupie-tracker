package internal

import (
	"html/template"
	"net/http"
	"log"
)

// ErrorPage renders the error.html template with the provided error message
func ErrorPage(w http.ResponseWriter, message string, code int) {
	tmpl, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		log.Printf("failed to parse error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code) // Set the appropriate HTTP status code
	err = tmpl.Execute(w, map[string]string{
		"ErrorMessage": message,
	})
	if err != nil {
		log.Printf("failed to execute error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
