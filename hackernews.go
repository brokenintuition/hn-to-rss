package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const hnTopStories = "https://hacker-news.firebaseio.com/v0/topstories.json"
const hnStoryURL = "https://hacker-news.firebaseio.com/v0/item/%d.json"
const hnPageURL = "https://news.ycombinator.com/item?id=%d"
const hnPageSize = 30

type hnLink struct {
	title string
	time  int
	url   string
	guid  string
}

func getFirstPage() ([]hnLink, error) {
	resp, err := http.Get(hnTopStories)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pages []int
	err = json.Unmarshal(body, &pages)
	if err != nil {
		return nil, err
	}

	var result []hnLink
	for _, page := range pages[:hnPageSize] {
		pageDetails, err := getStory(page)
		if err != nil {
			return nil, fmt.Errorf("processing page %d: %w", page, err)
		}

		result = append(result, pageDetails)
	}

	return result, nil
}

func getStory(pageID int) (hnLink, error) {
	var result hnLink
	resp, err := http.Get(fmt.Sprintf(hnStoryURL, pageID))
	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// unmarshalling to map because I don't care about most of the fields in the response
	// todo: repplace this because I don't need to specify the whole response
	var pageMap map[string]json.RawMessage
	json.Unmarshal(body, &pageMap)

	var pageType string
	json.Unmarshal(pageMap["type"], &pageType)

	if pageType != "story" {
		return result, errors.New("cannot process non-story post")
	}

	var title string
	json.Unmarshal(pageMap["title"], &title)

	var url string
	json.Unmarshal(pageMap["url"], &url)
	if len(url) == 0 {
		return result, errors.New("no link in story post. Could be Launch HN")
	}

	var time int
	json.Unmarshal(pageMap["time"], &time)

	result = hnLink{
		title: title,
		time:  time,
		url:   url,
		guid:  fmt.Sprintf(hnPageURL, pageID),
	}

	return result, nil
}
