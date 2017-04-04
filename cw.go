package main

import (
	"fmt"
	"github.com/peterzhang41/weigo/spider"
	"strconv"
)

func main() {
	const FIRST_ITEM = 1
	fmt.Println("--- running")

	//0.4
	var data = [][]string{{}}
	items_ids, items_names := spider.Read_csv_id("tmp_full_bak.csv")
	fmt.Println(items_ids[FIRST_ITEM:])

	fmt.Println("--- API getting price ---")
	for index, item_id := range items_ids[FIRST_ITEM:] {
		var a []string
		item_price := spider.CW_api_price(item_id)
		fmt.Println(strconv.Itoa(index+1), item_id, items_names[index+1], item_price)
		a = append(a, strconv.Itoa(index+1), item_id, items_names[index+1], item_price)
		data = append(data, a)
	}
	spider.WriteCSV("write_price.csv", data)

}
