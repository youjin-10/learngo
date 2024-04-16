package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

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

	// an empty map
	var results = make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	fmt.Println(results)
}

func hitUrl(url string) error {
	resp, err := http.Get(url)
	
	if err != nil || resp.StatusCode >= 400 {
		// handle error
		return errRequestFailed
	}
	return nil
}


