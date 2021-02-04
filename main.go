package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const hnTopStories = "https://hacker-news.firebaseio.com/v0/topstories.json"
const hnStoryURL = "https://hacker-news.firebaseio.com/v0/item/%d.json"
const hnPageSize = 30

type hnLink struct {
	title string
	time  int
	url   string
}

func main() {
	fmt.Println("Requesting top stories")

	resp, err := http.Get(hnTopStories)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var pages []int
	json.Unmarshal(body, &pages)

	pagesToFetch := pages[:hnPageSize]

	fetchedPages := retrieveLinks(pagesToFetch)

	for _, link := range fetchedPages {
		fmt.Printf("%s %s\n", link.title, link.url)
	}

	fmt.Println("Finished")
}

func retrieveLinks(pages []int) []hnLink {
	result := make([]hnLink, len(pages))

	for _, page := range pages {

		pageDetails, err := getStory(page)

		if err != nil {
			log.Printf("Error processing page %d: %s", page, err.Error())
			continue
		}

		result = append(result, *pageDetails)
	}

	return result
}

func getStory(pageID int) (*hnLink, error) {
	pageURL := fmt.Sprintf(hnStoryURL, pageID)
	fmt.Printf("Requesting %s\n", pageURL)

	resp, err := http.Get(pageURL)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// unmarshalling to map because I don't care about most of the fields in the response
	var pageMap map[string]json.RawMessage
	json.Unmarshal(body, &pageMap)

	var pageType string
	json.Unmarshal(pageMap["type"], &pageType)

	if pageType != "story" {
		return nil, errors.New("cannot process non-story post")
	}

	var title string
	json.Unmarshal(pageMap["title"], &title)

	var url string
	json.Unmarshal(pageMap["url"], &url)
	var time int
	json.Unmarshal(pageMap["time"], &time)

	result := hnLink{
		title: title,
		time:  time,
		url:   url,
	}

	return &result, nil
}
