package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type exchange struct {
	*scraper
	result result
}

func NewExchange(date time.Time, url string, dompath string) *exchange {
	y := date.Format("06")
	m := fmt.Sprintf("%02d", int(date.Month()))
	d := date.Format("02")
	qs := fmt.Sprintf("%v%v%v", y, m, d)

	return &exchange{
		scraper: &scraper{
			date:    date,
			url:     url + qs,
			dompath: dompath,
		},
	}
}

func (e *exchange) get() error {
	fmt.Println("Querying " + e.url + " ...")

	err := e.scrape()
	if err != nil {
		return err
	}

	tr := e.doc.Find("div#main table.data-table7 tbody tr").Eq(1)
	h2 := e.doc.Find("div#main h2")

	date, err := time.Parse("January 2, 2006", strings.TrimSpace(strings.Split(h2.Text(), "As of ")[1]))
	if err != nil {
		return err
	}

	e.result.date = date
	e.result.currency = strings.TrimSpace(tr.Find("td").First().Text())

	sTTS := strings.TrimSpace(tr.Find("td").Eq(3).Text())
	if sTTS == "" {
		return errors.New("TTS is empty")
	}

	TTS, err := strconv.ParseFloat(sTTS, 64)
	if err != nil {
		return err
	}

	sTTB := strings.TrimSpace(tr.Find("td").Eq(4).Text())
	if sTTB == "" {
		return errors.New("TTB is empty")
	}

	TTB, err := strconv.ParseFloat(sTTB, 64)
	if err != nil {
		return err
	}

	TTM := strconv.FormatFloat((TTS+TTB)/2, 'f', -1, 64)
	e.result.value = TTM

	fmt.Printf("%v: %v = %v\n", e.result.date.Format("2006/01/02"), "USD TTM", e.result.value)

	return nil
}
