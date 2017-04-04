package woolies

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type prodcutsDetail struct {
	Description string
	Price       json.Number
	CupString   string
	//MediumImageFile string
}

type prodcuts struct {
	Products []prodcutsInfo `json:"Products"`
}

type prodcutsInfo struct {
	Products []prodcutsDetail `json:"Products"`
}

//TODO reorder according to CupString
//TODO add imageURL
func GetItem(products string) prodcutsDetail {

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
	fmt.Println(r[0])
	return r[0]
}
