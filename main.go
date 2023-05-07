package main

import (
	"log"
	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"https://www.firmy.cz/?q=%C3%BA%C4%8Detn%C3%AD+firmy",
	}

	var g errgroup.Group

	// Make concurrent requests
	for _, url := range urls {
		url := url // Create a new variable to avoid a data race
		g.Go(func() error {
			return scrape(url)
		})
	}

	// Wait for all the goroutines to finish
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
