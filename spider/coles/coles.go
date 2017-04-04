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