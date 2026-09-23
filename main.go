package main

import (
	"fmt"
	"io"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://books.toscrape.com/", nil)
	check(err)
	response, err := client.Do(request)
	check(err)
	response_body, err := io.ReadAll(response.Body)
	check(err)
	fmt.Println(string(response_body))
}
