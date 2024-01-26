package types

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Currency struct {
	Name  string
	Price string
}

func ProcessCurrencyPriceData(price string) string {
	price = price[:strings.Index(price, "(")]
	price = strings.ReplaceAll(price, ",", "")

	return price
}

func CurrenciesToCSV(currencies []Currency, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"name",
		"price",
	}
	writer.Write(headers)

	for _, currency := range currencies {
		record := []string{
			currency.Name,
			currency.Price,
		}

		writer.Write(record)
	}
	defer writer.Flush()
	defer fmt.Printf("currencies saved in %s successfully.", fileName)
	return nil
}
