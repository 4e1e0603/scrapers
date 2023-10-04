package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// TODO Configure with YAML

func main() {

	fmt.Println("Start")

	c := colly.NewCollector()

	// Find and visit all links.
	c.OnHTML("p[data-test='product-overview-availability']", func(e *colly.HTMLElement) {
		// Vyprodáno
		// Dočasně není na skladě
		fmt.Println("LEGO 51515", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.lego.com/cs-cz/product/robot-inventor-51515")
}
