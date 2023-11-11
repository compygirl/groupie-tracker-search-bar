package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	endpoints "groupieTrecker/internal/endpoints"
)

// artist pages

func ErrorCheck(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	styles := http.FileServer(http.Dir("internal/templates/css"))

	router := http.NewServeMux()
	router.Handle("/css/", http.StripPrefix("/css/", styles))
	router.HandleFunc("/", endpoints.GetMainPage)
	router.HandleFunc("/artist", endpoints.GetArtist)
	router.HandleFunc("/search", endpoints.SearchHandler)
	fmt.Println("listening on: http://localhost:" + port + "/")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error: several connecrtions")
		return
	}
}
