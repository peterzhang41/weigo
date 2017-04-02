package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type prodcutsDetail struct {
	Description string
	Price       float64
	CupString   string
	//MediumImageFile string
}

type prodcuts struct {
	Products []prodcutsInfo `json:"Products"`
}

type prodcutsInfo struct {
	Products []prodcutsDetail `json:"Products"`
}

func fillSpace(orignal string) string {
	return (strings.Replace(orignal, " ", "%20", -1))
}

//TODO add CupString imgaeURL
//TODO binding name with brand to fit with woolworths Description
//TODO array result --> reorder according to CupString
func colesScrape(products string) {
	link := "https://shop.coles.com.au/online/SearchDisplay?storeId=10601&catalogId=10576&langId=-1&beginIndex=0&browseView=false&searchSource=Q&sType=SimpleSearch&resultCatEntryType=2&showResultsPage=true&pageView=image&supermarketRefer=yes&searchTerm=" + products

	fmt.Println("---link:", link)
	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".wrapper.clearfix").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		brand := s.Find(".brand").Text()
		//title := s.Find("i").Text()
		fmt.Print(i, " brand:", brand)

		name := s.Find(".product-url").Text()
		//title := s.Find("i").Text()
		fmt.Print(" name:", name)

		price := s.Find(".price").Text()
		//title := s.Find("i").Text()
		fmt.Println(price)
	})
}

//TODO reorder according to CupString
//TODO add imageURL
func woolworthScrape(products string) {

	var p prodcuts
	//var f interface{}
	// var body struct {
	//      // httpbin.org sends back key/value pairs, no map[string][]string
	//      Headers map[string]string `json:"headers"`
	//      Origin  string            `json:"origin"`
	//   }
	link := "https://www.woolworths.com.au/apis/ui/Search/products?IsSpecial=false&PageNumber=1&PageSize=36&SearchTerm=" + products + "&SortType=Relevance"

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err) // handle error
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	var r []prodcutsDetail
	for _, element := range p.Products {
		// index is the index where we are
		// element is the element from someSlice for where we are
		for _, element := range element.Products {
			r = append(r, element)
		}
	}
	fmt.Println(r)
}

//TODO fetch coles searchSuggestion
//TODO fetch woolworths searchSuggestion
//TODO fetch coles resultNumber
//TODO fetch woolworths resultNumber
//TODO a html template with searchbar and table
//TODO fill data into template when requesting
//TODO host a server and listen on 8080
//TODO httprequest reponse-->searchSuggestion from Coles and woolworths
//TODO httprequest reponse-->search if resultNumber less than 6

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
