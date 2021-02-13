package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	listenPort := os.Getenv("HN_LISTEN_PORT")

	if listenPort == "" {
		log.Fatal("Environment variable HN_LISTEN_PORT must be set")
	}

	http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)
}
