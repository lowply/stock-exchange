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
	result string
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

	if tr.Text() == "" {
		fmt.Printf("%v: No exchange data\n", e.date.Format("2006/01/02"))
		return nil
	}

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
	e.result = TTM

	fmt.Printf("%v: %v = %v\n", e.date.Format("2006/01/02"), "JPY/USD TTM", e.result)

	return nil
}
