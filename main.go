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

type Output struct {
	Title    string `json:"title"`
	Region   string `json:"region"`
	District string `json:"district"`
	Category string `json:"category"`
}

func GetRegionCode(name string) int {
	var region = map[string]int{
		"Praha": 14,
	}
	return region[name]
}

// Returns (kraj, okres, město, vloženo, zobrazeno)
func parseDetail(input string) string {
	space := regexp.MustCompile(`\s+`)
	// detail := strings.Split(space.ReplaceAllString(strings.TrimSpace(input), " "), "\r\n")
	detail := strings.TrimSpace(space.ReplaceAllString(input, " "))
	// return detail[0], detail[1], detail[2], detail[3], detail[4]
	return detail
}

func filter(slice []string, f func(string) bool) []string {
	var n []string
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func Scrape(url string, category string) {

	//var region, district, town, created, viewed string

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
		title := strings.TrimSpace(space.ReplaceAllString(parseDetail(e.Text), " "))
		fmt.Println("---")
		fmt.Println(category)
		fmt.Println(title)
	})

	c.OnHTML(".item-table tbody", func(e *colly.HTMLElement) {
		// region, district, town, created, viewed = parseDetail(e.Text)
		space := regexp.MustCompile(`\s+`)
		texts := strings.TrimSpace(space.ReplaceAllString(parseDetail(e.Text), " "))

		texts = strings.ReplaceAll(texts, "(mapa)", "")
		texts = strings.ReplaceAll(texts, "hodinami", "hodin")
		texts = strings.ReplaceAll(texts, "minutami", "minut")
		texts = strings.ReplaceAll(texts, "Kraj", "")
		texts = strings.ReplaceAll(texts, "Okres", "")
		texts = strings.ReplaceAll(texts, "Město", "")
		texts = strings.ReplaceAll(texts, "Vloženo", "")
		texts = strings.ReplaceAll(texts, "Zobrazeno", "")
		texts = strings.ReplaceAll(texts, "před", "")

		split := strings.Split(texts, ":")
		for i, s := range split {
			split[i] = strings.TrimSpace(strings.ReplaceAll(s, ":", ""))
		}

		split = filter(split, func(v string) bool {
			return v != ""
		})

		fmt.Println(split[0], split[1], split[2], split[3], split[4])

	})

	c.Visit(url)
}

func main() {

	url := "https://vsezaodvoz.cz/inzeraty/dum-a-zahrada?region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77"

	Scrape(url, "dum-a-zahrada")
}
