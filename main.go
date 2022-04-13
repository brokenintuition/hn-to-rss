package main

import (
	"fmt"
	"log"
	"net/http"
)

func hn(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")

	fetchedPages, err := getFirstPage()
	if err != nil {
		log.Printf("Error processing request: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rssFeed, err := RssFromHNItems(fetchedPages)
	if err != nil {
		log.Printf("Error processing request: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(rssFeed))
}

func main() {
	http.HandleFunc("/", hn)
	http.ListenAndServe(":80", nil)
}
