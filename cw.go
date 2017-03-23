package main

import (
	//0.1
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"

	//0.2
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

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

//0.2
// read template CSV file
// parameter file path
// return a list
func readCSV(path string) (items []string) {
	f, _ := os.Open(path)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.

		// fmt.Println(len(record))
		// for value := range record {
		// 	fmt.Printf("  %v\n", record[value])
		// }

		//0.2
		//read item english name
		// fmt.Println(record[1])
		items = append(items, record[1])
	}
	return
}

func main() {

	fmt.Println("--- running")
	//0.1
	//cwScrape("A2")

	//0.2

	fmt.Println(readCSV("template_10.csv")[1])
	cwScrape(readCSV("template_10.csv")[1])
}
