package cw

import (
	//0.1
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type product struct {
	Id           string
	Label        string
	Value        string
	Category     string
	Image        string
	Price        string
	Savings      string
	IsScript     string
	Ams_category string
	Ams_schedule string
}

//0.4 CW API
func GetPrice(id string) (result string) {
	result = "	"
	var products []product
	link := "http://www.chemistwarehouse.com.au/cmsglobalfiles/handlers/predictive_search.ashx?term=" + id
	//fmt.Println(link)

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err) // handle error
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range products {
		result = element.Price
		resp.Body.Close()
		return
	}
	return
}

//0.1
func cwScrape(item string) {

	cw_link := "http://www.chemistwarehouse.com.au/search?searchtext=" + strings.Replace(item, " ", "%20", -1) + "&searchmode=allwords"

	fmt.Println(cw_link)

	doc, err := goquery.NewDocument(cw_link)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".Product").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		name := s.Find(".product-name").Text()
		price := strings.TrimSpace(s.Find(".Price").Text())
		fmt.Printf("--- %d: %s price: %s\n", i, name, price)
	})
}

func fillSpace(orignal string) string {
	return (strings.Replace(orignal, " ", "%20", -1))
}
