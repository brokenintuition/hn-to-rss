package main

import (
	"fmt"
)

func main() {
	fmt.Println("Requesting top stories")

	fetchedPages := getFirstPage()

	for _, link := range fetchedPages {
		fmt.Printf("%s %s\n", link.title, link.url)
	}

	fmt.Println("Finished")
}
