//TODO add CupString imgaeURL
//TODO binding name with brand to fit with woolworths Description
//TODO array result --> reorder according to CupString
package coles

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetItem(item string) {
	link := "https://shop.coles.com.au/online/SearchDisplay?storeId=10601&catalogId=10576&langId=-1&beginIndex=0&browseView=false&searchSource=Q&sType=SimpleSearch&resultCatEntryType=2&showResultsPage=true&pageView=image&supermarketRefer=yes&searchTerm=" + strings.Replace(item, " ", "%20", -1)

	fmt.Println("---link:", link)
	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items

	// For each item found, get the band and title
	brand := doc.Find("span.product-brand").Text()
	//title := s.Find("i").Text()
	fmt.Print(" brand:", brand)

	name := doc.Find("span.product-name").Text()
	//title := s.Find("i").Text()
	fmt.Print(" name:", name)

	price := doc.Find("strong.product-price").Text()
	//title := s.Find("i").Text()
	fmt.Println(price)

}
