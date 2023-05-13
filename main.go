package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("div.product-inline-item", func(e *colly.HTMLElement) {
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(e.Text, " ")
		fmt.Println("---")
		fmt.Println(s)
	})

	c.Visit("https://vsezaodvoz.cz/inzeraty/dum-a-zahrada?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77")
}
