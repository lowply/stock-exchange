package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type scraper struct {
	name    string
	date    time.Time
	url     string
	dompath string
	doc     *goquery.Document
	result  string
}

func (s *scraper) scrape() error {
	res, err := http.Get(s.url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	s.doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *scraper) output() {
	fmt.Printf("%v: %v = %v\n", s.date.Format("2006/01/02"), s.name, s.result)
}
