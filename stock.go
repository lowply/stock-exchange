package main

import (
	"fmt"
	"time"
)

type stock struct {
	*scraper
}

func NewStock(ticker string, date time.Time, url string, dompath string) *stock {
	ymd := date.Format("20060102")
	qs := fmt.Sprintf("from=%v&to=%v&timeFrame=d&page=1", ymd, ymd)

	return &stock{
		scraper: &scraper{
			name:    ticker,
			date:    date,
			url:     url + ticker + "/history?" + qs,
			dompath: dompath,
		},
	}
}

func (s *stock) get() error {
	fmt.Println("Querying " + s.url + " ...")
	err := s.scrape()
	if err != nil {
		return err
	}

	td := s.doc.Find(s.dompath)

	if td.Text() == "" {
		fmt.Printf("%v: No stock data\n", s.date.Format("2006/01/02"))
		return nil
	}

	s.result = td.Last().Text()
	s.output()

	return nil
}
