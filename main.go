package main

import (
	"fmt"
	"log"
	"net/http"
)

func hn(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")

	fetchedPages := getFirstPage()
	rssFeed, err := RssFromHNItems(fetchedPages)

	if err != nil {
		log.Printf("Error processing request: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(rssFeed))
}

func main() {
	http.HandleFunc("/", hn)

	listenPort := "80"

	http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)
}
