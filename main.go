package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

// Příklad URL a parametrů:
// https://vsezaodvoz.cz/inzeraty/elektro?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77"

// Filtry:
// /elektro? ... => kateorie
// - type=1 => pouze nabídky (bez potávek)
// - new=1 => pouze nové (24h)
// - with_photo=true => pouze s fotografií
// - no_reservation=true => pouze bez rezervace
// - region=14 => Kraj Praha (další viz kód)
// - district=77 => Okres Praha (další viz kód)

func GetRegionCode(name string) int {
	var region = map[string]int{
		"Praha": 14,
	}
	return region[name]
}

func Scrape(url string) {

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

	c.Visit(url)
}

func main() {

	url := "https://vsezaodvoz.cz/inzeraty/dum-a-zahrada?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77"

	Scrape(url)
}
