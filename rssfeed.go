package main

import "encoding/xml"

const feedTitle = "HackerNews Front Page"
const feedURL = "https://news.ycombinator.com"
const feedDescription = "Links pulled from the front page of Hacker News"

type rss struct {
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"'`
	Link  string `xml:"link"`
}

func RssFromHNItems(pages []hnLink) ([]byte, error) {
	var items []Item

	for _, link := range pages {
		item := Item{Title: link.title, Link: link.url}
		items = append(items, item)
	}

	channel := Channel{
		Title:       feedTitle,
		Link:        feedURL,
		Description: feedDescription,
		Items:       items,
	}

	rss := rss{Channel: channel, Version: "2.0"}

	return xml.MarshalIndent(rss, "", "   ")
	//return xml.Marshal(rss)
}
