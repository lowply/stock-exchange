package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type scraper struct {
	date    time.Time
	url     string
	dompath string
	doc     *goquery.Document
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
