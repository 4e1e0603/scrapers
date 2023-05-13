package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	c.OnHTML("h3.product-inline-title a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnHTML("div.item-detail-content", func(e *colly.HTMLElement) {
		space := regexp.MustCompile(`\s+`)
		texts := space.ReplaceAllString(strings.TrimSpace(e.Text), " ")
		fmt.Println("---")
		fmt.Println(strings.TrimSpace(texts))
	})

	c.OnHTML(".item-table tr", func(e *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(e.Text))
	})

	c.Visit("https://vsezaodvoz.cz/inzeraty/dum-a-zahrada?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77")
}
