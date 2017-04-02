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

	//0.3
	"encoding/json"
	"net/http"
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

	//MediumImageFile string
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

//0.2
// read template CSV file
// parameter file path
// return a list
func read_csv_name(path string) (items []string) {
	f, _ := os.Open(path)
	const ENGLISH_COLUMN = 1

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
		//read item english name list
		//fmt.Println(record[1])
		items = append(items, record[ENGLISH_COLUMN])
	}
	return
}

//0.3
func writeCSV(path string, data [][]string) {
	file, err := os.Create(path)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

//0.3
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//0.3
func cw_api_id(item string) (result string) {
	result = "	"
	var products []product
	link := "http://www.chemistwarehouse.com.au/cmsglobalfiles/handlers/predictive_search.ashx?term=" + strings.Replace(item, " ", "%20", -1)
	fmt.Println(link)
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
		if element.Label == item {
			//fmt.Printf(element.Id)
			result = element.Id
			return
		}

	}
	return
}

func main() {
	const FIRST_ITEM = 1
	fmt.Println("--- running")
	//0.1
	//cwScrape("A2")

	//0.3
	var data = [][]string{{}}
	item_names := read_csv_name("tmp_full.csv")[FIRST_ITEM:]
	fmt.Println(item_names)

	for _, item_name := range item_names {
		fmt.Println("--- API ing ---")
		item_id := cw_api_id(item_name)

		fmt.Println(item_name, item_id)

		var a []string
		a = append(a, item_name)
		a = append(a, item_id)
		data = append(data, a)
	}
	writeCSV("write_id.csv", data)

}
