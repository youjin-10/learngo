package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct{
	url string
	status string
}

var errRequestFailed = errors.New("request failed")

func main() {
	//  create a channel
	c := make(chan requestResult)

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
	}

	// an empty map
	results := make(map[string]string)

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c // receiving from channel is blocking operation!!

		results[result.url] = result.status
	}

	fmt.Println(results)
}

// send-only channel
func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- requestResult{url: url, status: status}
}


