package main

import (
	"encoding/json"
	"flag"
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
	Town     string `json:"town"`
	Category string `json:"category"`
	Created  string `json:"created"`
	Viewes   string `json:"viewes"`
	Link     string `json:"link"`
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

func clean(input string) string {
	space := regexp.MustCompile(`\s+`)
	texts := strings.TrimSpace(space.ReplaceAllString(parseDetail(input), " "))
	// Time
	texts = strings.ReplaceAll(texts, "hodinami", "hours")
	texts = strings.ReplaceAll(texts, "sekundami", "seconds")
	texts = strings.ReplaceAll(texts, "minutami", "minutes")
	texts = strings.ReplaceAll(texts, "60 minutami", "1 hour")
	texts = strings.ReplaceAll(texts, "hodinou", "1 hour")
	texts = strings.ReplaceAll(texts, "včera", "1 day")
	// Other
	texts = strings.ReplaceAll(texts, "(mapa)", "")
	texts = strings.ReplaceAll(texts, "Kraj", "")
	texts = strings.ReplaceAll(texts, "Okres", "")
	texts = strings.ReplaceAll(texts, "Město", "")
	texts = strings.ReplaceAll(texts, "Vloženo", "")
	texts = strings.ReplaceAll(texts, "Zobrazeno", "")
	texts = strings.ReplaceAll(texts, "před", "")

	return texts
}

func Scrape(url string, category string) {

	//var region, district, town, created, viewed string

	c := colly.NewCollector()

	o := Output{}
	o.Category = category // TODO Validate category, allow string and number

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
		o.Link = fmt.Sprintf("%s%s", r.Request.URL.Host, r.Request.URL.Path)
	})

	c.OnHTML("h3.product-inline-title a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnHTML("div.item-detail-content", func(e *colly.HTMLElement) {
		space := regexp.MustCompile(`\s+`)
		title := strings.TrimSpace(space.ReplaceAllString(parseDetail(e.Text), " "))
		o.Title = title
	})

	c.OnHTML(".item-table tbody", func(e *colly.HTMLElement) {

		texts := clean(e.Text)

		split := strings.Split(texts, ":")
		for i, s := range split {
			split[i] = strings.TrimSpace(strings.ReplaceAll(s, ":", ""))
		}

		split = filter(split, func(v string) bool {
			return v != ""
		})

		o.Region = split[0]
		o.District = split[1]
		o.Town = split[2]
		o.Created = split[3]
		o.Viewes = strings.ReplaceAll(split[4], "x", "")

		result, err := json.Marshal(o)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(result))
	})

	url = fmt.Sprintf("%s/%s?%s", url, category, "region=14&type=1&new=1&with_photo=true&no_reservation=true&district=77")

	c.Visit(url)

}

var (
	categoryFlag string
)

func main() {

	flag.StringVar(&categoryFlag, "c", "", "Select the category")

	flag.Parse()

	baseURL := "https://vsezaodvoz.cz/inzeraty"

	Scrape(baseURL, categoryFlag)
}
