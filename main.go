package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func usage() {
	fmt.Println("Usage: go run . [ticker] [YYYY-MM-DD]")
	os.Exit(1)
}

func exec() error {
	date, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		return err
	}

	ticker := os.Args[1]

	stock := NewStock(
		ticker,
		date,
		"https://stocks.finance.yahoo.co.jp/us/history/",
		"div#main .padT12 table tbody tr td",
	)

	exchange := NewExchange(
		date,
		"http://www.murc-kawasesouba.jp/fx/past/index.php?id=",
		"div#main table.data-table7 tbody tr",
	)

	err = stock.get()
	if err != nil {
		return err
	}

	err = exchange.get()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}
