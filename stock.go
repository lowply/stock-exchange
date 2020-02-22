package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type stock struct {
	url     string
	date    string
	dompath string
	result  Result
}

func NewStock() *stock {
	ticker := os.Args[1]
	S := &stock{
		date:    os.Args[2],
		url:     "https://stocks.finance.yahoo.co.jp/us/history/" + ticker + "?",
		dompath: "div#main .padT12 table tbody tr td",
	}
	return S
}

func (req *stock) qs() (string, error) {
	t, err := time.Parse("2006-01-02", req.date)
	if err != nil {
		return "", err
	}
	y := t.Year()
	m := int(t.Month())
	d := t.Day()
	return fmt.Sprintf("sy=%v&sm=%v&sd=%v&ey=%v&em=%v&ed=%v&tm=d", y, m, d, y, m, d), nil
}

func (req *stock) query() error {
	qs, err := req.qs()
	if err != nil {
		return err
	}

	req.url = req.url + qs
	fmt.Println("Querying " + req.url + " ...")

	res, err := http.Get(req.url)
	if err != nil {
		return err

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	td := doc.Find(req.dompath)
	req.result.date = td.First().Text()
	req.result.value = td.Last().Text()

	return nil
}
