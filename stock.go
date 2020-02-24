package main

import (
	"fmt"
	"time"
)

type stock struct {
	ticker string
	*scraper
	result result
}

func NewStock(ticker string, date time.Time, url string, dompath string) *stock {
	y := date.Year()
	m := int(date.Month())
	d := date.Day()
	qs := fmt.Sprintf("sy=%v&sm=%v&sd=%v&ey=%v&em=%v&ed=%v&tm=d", y, m, d, y, m, d)

	return &stock{
		ticker: ticker,
		scraper: &scraper{
			date:    date,
			url:     url + ticker + "?" + qs,
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

	date, err := time.Parse("2006年1月2日", td.First().Text())
	if err != nil {
		return err
	}

	s.result.date = date
	s.result.value = td.Last().Text()

	fmt.Printf("%v: %v = %v\n", s.result.date.Format("2006/01/02"), s.ticker, s.result.value)

	return nil
}
