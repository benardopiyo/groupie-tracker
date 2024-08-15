package main

import (
	"fmt"
	"groupie-Tracker/internal"
	"net/http"
)

func main() {
	internal.Routes()

	serve := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: nil,
	}
	fmt.Println("Listening on http://127.0.0.1:8000")

	err := serve.ListenAndServe()
	if err != nil {
		fmt.Println("Error Starting the Server error:", err)
	}

}
