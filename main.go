package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(resp)

	defer resp.Body.Close() // preventing memory leaks

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	doc.Find(".pagination").Each()

	return 0;
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln(errors.New("Request failed with status: " + fmt.Sprint(res.StatusCode)))
	}
}


