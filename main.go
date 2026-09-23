package main

import (
	"fmt"
	"net/http"
	"web-crawler/extractor"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client := &http.Client{}
	URL := "https://books.toscrape.com/"
	request, err := http.NewRequest("GET", URL, nil)
	check(err)
	response, err := client.Do(request)
	check(err)
	defer response.Body.Close()
	links, err := extractor.Extract_links(response.Body)
	check(err)
	fmt.Println(links[10])
}
