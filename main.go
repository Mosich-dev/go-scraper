package main

import (
	"fmt"
	"github.com/Mosich-dev/go-scraper/types"
	"github.com/gocolly/colly"
	"log"
)

const URL = "https://www.tgju.org/currency"

func main() {
	var currencies []types.Currency

	c := colly.NewCollector()

	c.OnHTML("tbody tr[data-title].pointer", func(e *colly.HTMLElement) {
		var currency types.Currency
		currency.Name = e.ChildText("th")
		currency.Price = types.ProcessCurrencyPriceData(e.ChildText("td.nf"))

		currencies = append(currencies, currency)
	})
	err := c.Visit(URL)
	if err != nil {
		fmt.Println("Error in visit:", err)
	}

	if err = types.CurrenciesToCSV(currencies, "Prices.csv"); err != nil {
		log.Fatal(err)
	}
}
