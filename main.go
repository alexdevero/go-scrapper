package main

import (
	"log"
	"os"
	"strconv"
	"golang.org/x/sync/errgroup"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	maxRetries, maxRetriesErr := strconv.Atoi(os.Getenv("MAX_RETRIES"))
	if maxRetriesErr != nil {
		log.Fatal("Error converting MAX_RETRIES to int")
	}
	itemSelector := os.Getenv("LISTING_ITEM_LINK_SELECTOR")
	filePath := os.Getenv("FILE_PATH")

	urls := []string{
		os.Getenv("LISTING_URL"),
	}

	var g errgroup.Group

	// Make concurrent requests
	for _, url := range urls {
		url := url // Create a new variable to avoid a data race
		g.Go(func() error {
			return scrape(url, maxRetries, itemSelector, filePath)
		})
	}

	// Wait for all the goroutines to finish
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
