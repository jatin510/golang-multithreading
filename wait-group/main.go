package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func getStatusCode(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Println("Oops, something went wrong")
	}

	log.Println("Status code for", endpoint, "is", res.StatusCode)

}

func main() {
	websiteList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.medium.com",
		"https://fb.com",
		"https://linkedin.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.medium.com",
		"https://fb.com",
		"https://linkedin.com",
	}

	start := time.Now()
	var wg sync.WaitGroup

	for _, web := range websiteList {
		go getStatusCode(web, &wg)
		wg.Add(1)
	}
	wg.Wait()

	log.Println("took", time.Since(start), "seconds")
}
