/*
The code demonstrates web scraping of hospital details with colly library.
Name, address, establishment date, number of beds, physician name, and reviews
are among the extracted data. The code shows how to construct HTML element callbacks,
handle errors, set up a collector for each URL, and publish the hospital information that has been gathered.
*/
package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

func main() {
	// List of websites to scrape
	websites := []string{

		"https://www.vaidam.com/hospitals/university-hospital-rechts-der-isar",
		"https://www.vaidam.com/hospitals/university-hospital-heidelberg",
		"https://www.vaidam.com/hospitals/charite-university-hospital",
		"https://www.vaidam.com/hospitals/university-hospital-dusseldorf",
		"https://www.vaidam.com/hospitals/university-hospital-frankfurt-am-main-frankfurt",
	}

	fmt.Println("\n----Hospital Details----\n")
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for _, url := range websites {
		wg.Add(1) //increments the WaitGroup counter, indicating the creation of a new goroutine.

		// Visit the URL and perform scraping in a goroutine
		go func(url string) {
			defer wg.Done() //indicates WaitGroup counter is decremented when the goroutine completes its execution

			// Create a new collector for each URL
			c := colly.NewCollector()

			var hospitalName, address, established, beds, doctorName, reviews string

			c.OnHTML("h1", func(e *colly.HTMLElement) {
				// Extract hospital name
				hospitalName = e.Text
			})

			c.OnHTML("#section-address > div > p:nth-child(3)", func(e *colly.HTMLElement) {
				// Extract address
				address = strings.TrimSpace(e.Text)
			})

			c.OnHTML("#mainSlider > div.col-md-12.col-sm-12.col-xs-12.hosp-detail-listing.all-padding-0 > div.col-sm-8.col-xs-12.col-sm-pull-4.all-padding-0 > div > ul:nth-child(2) > li:nth-child(1) > span", func(e *colly.HTMLElement) {
				// Extract hospital built info
				established = strings.TrimSpace(e.Text)
			})

			c.OnHTML("#mainSlider > div.col-md-12.col-sm-12.col-xs-12.hosp-detail-listing.all-padding-0 > div.col-sm-8.col-xs-12.col-sm-pull-4.all-padding-0 > div > ul:nth-child(1) > li:nth-child(2)", func(e *colly.HTMLElement) {
				// Extract no of beds
				beds = strings.TrimSpace(e.Text)
			})

			c.OnHTML(".video-treatment-block", func(e *colly.HTMLElement) {
				// Extract doctor name
				doctorName = strings.TrimSpace(e.Text)
			})

			// Assuming reviews are in an element with class ".review"
			c.OnHTML("body > section.rating-section.section-padding.bg-light-grey.pt-0 > div > h4", func(e *colly.HTMLElement) {
				// Extract reviews
				reviews = strings.TrimSpace(e.Text)
			})

			// Define error handling for each request
			c.OnError(func(r *colly.Response, err error) {
				log.Printf("Request URL: %s failed with response: %v\nError: %s\n", r.Request.URL, r, err)
			})

			// Visit the URL using the collector
			err := c.Visit(url)

			// Handle errors if the visit fails
			if err != nil {
				log.Printf("Error visiting URL %s: %s\n", url, err)
			}

			// Printing hospital details including reviews
			fmt.Printf("Hospital Name: %s\nAddress: %s Germany\nEstablished: %s\nNumber of Beds: %s\nDoctor Name: %s\nReviews: %s\n\n",
				hospitalName, address, established, beds, doctorName, reviews)

		}(url)
	}

	// Wait for all goroutines to finish before exiting
	wg.Wait()

	fmt.Println("All URLs scraped.")
}
