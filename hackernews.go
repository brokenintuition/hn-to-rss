package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	hnApiBase  = "https://hacker-news.firebaseio.com/v0/"
	hnPageURL  = "https://news.ycombinator.com/item?id=%d"
	hnPageSize = 30
)

var (
	emptyLink = hnLink{}
)

type hnLink struct {
	title string
	time  int
	url   string
	guid  string
}

// storyResponse is the pieces of the story json that need to be exposed for rss
type storyResponse struct {
	Type  string
	Title string
	Url   string
	Time  int
}

func getFirstPage() ([]hnLink, error) {
	resp, err := http.Get(hnApiBase + "topstories.json")
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
		storyLink, err := getStory(page)
		if err != nil {
			return nil, fmt.Errorf("processing page %d: %w", page, err)
		}

		if storyLink != emptyLink {
			result = append(result, storyLink)
		}
	}

	return result, nil
}

func getStory(pageID int) (hnLink, error) {
	var result hnLink
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", hnApiBase, pageID))
	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var story storyResponse
	err = json.Unmarshal(body, &story)
	if err != nil {
		return result, err
	}

	// skip non-story posts and posts with no URL (usually Launch HN)
	if story.Type != "story" || len(story.Url) == 0 {
		return result, nil
	}

	result = hnLink{
		title: story.Title,
		time:  story.Time,
		url:   story.Url,
		guid:  fmt.Sprintf(hnPageURL, pageID),
	}

	return result, nil
}
