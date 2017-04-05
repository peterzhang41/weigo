package main

import (
	"fmt"
	"strconv"

	"github.com/peterzhang41/weigo/spider/csv"
	"github.com/peterzhang41/weigo/spider/cw"
	"github.com/peterzhang41/weigo/spider/woolies"
)

func main() {
	const FIRST_ITEM = 1
	fmt.Println("--- running")

	//0.4
	var data = [][]string{{}}
	items_ids, items_names := csv.Read("tmp_full_bak.csv")
	fmt.Println(items_ids[FIRST_ITEM:])

	fmt.Println("--- API getting price ---")
	var count = 0
	for index, item_id := range items_ids[FIRST_ITEM:] {
		var a []string
		item_price := cw.GetPrice(item_id)
		fmt.Println(strconv.Itoa(index+1), item_id, items_names[index+1], item_price)
		a = append(a, strconv.Itoa(index+1), item_id, items_names[index+1], item_price)
		data = append(data, a)
		count = index
	}
	var wooliesA2 = "855562"
	var milk = woolies.GetItem(wooliesA2)
	var a = []string{strconv.Itoa(count + 2), wooliesA2, milk.Description, string(milk.Price)}
	data = append(data, a)
	csv.Write("priceList.csv", data)

}
