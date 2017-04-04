package spider

import (
	//0.1
	"log"

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
}

//0.3
func WriteCSV(path string, data [][]string) {
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

//0.4
func CW_api_price(id string) (result string) {
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

//0.4
// read template CSV file
// parameter file path
// return a list
func Read_csv_id(path string) (items_ids []string, items_names []string) {
	f, _ := os.Open(path)
	const NAME_COLUMN = 0
	const ID_COLUMN = 2

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		items_names = append(items_names, record[NAME_COLUMN])
		items_ids = append(items_ids, record[ID_COLUMN])
	}
	return
}
