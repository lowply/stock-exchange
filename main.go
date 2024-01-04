package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func usage() {
	fmt.Println("Usage: ./stock-exchange [ticker] [YYYY-MM-DD]")
	os.Exit(1)
}

func exec() error {
	ticker := os.Args[1]
	date, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		return err
	}

	stock := NewStock(
		ticker,
		date,
		"https://finance.yahoo.co.jp/quote/",
		"body div#wrapper div#root main div div div div section div table tbody tr td",
	)

	exchange := NewExchange(
		date,
		"https://www.murc-kawasesouba.jp/fx/past/index.php?id=",
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
