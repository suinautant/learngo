package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type requestResult struct {
	url    string
	status string
	code   int
}

func main() {
	results := make(map[string]string)
	codes := make(map[string]int)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
		codes[result.url] = result.code
	}
	for url, status := range results {
		fmt.Println(strconv.Itoa(codes[url]) + " : " + status + " - " + url)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	code := resp.StatusCode
	if err != nil || code >= 400 {
		status = "FAILED"
		c <- requestResult{url: url, status: status, code: code}
	}
	c <- requestResult{url: url, status: status, code: code}
}
