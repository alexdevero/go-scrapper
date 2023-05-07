package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
)

type LinkInfo struct {
	Link string
	Title string
}

const maxRetries = 5

func scrape(url string, maxRetries int, selector string, filePath string) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fetchAndExtract(url, selector, filePath)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	return err
}

func fetchAndExtract(url string, selector string, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error for URL %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("error loading document for URL %s: %w", url, err)
	}

	links := []LinkInfo{}

	fmt.Printf("Links for URL %s:\n", url)
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			title := s.Text()

			linkInfo := LinkInfo{
				Title: title,
				Link:  href,
			}

			links = append(links, linkInfo)
		}
	})

	// fmt.Println(links)
	saveDataToFile(links, filePath)

	return nil
}
