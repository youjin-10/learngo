package main

import (
	"errors"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
	}

	for _, url := range urls {
		hitUrl(url)
	}
}

func hitUrl(url string) error {
	resp, err := http.Get(url)
	
	if err != nil || resp.StatusCode >= 400 {
		// handle error
		return errRequestFailed
	}
	return nil
}


