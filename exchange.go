package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type exchange struct {
	url     string
	date    string
	dompath string
	result  Result
}

func NewExchange() *exchange {
	E := &exchange{
		date:    os.Args[2],
		url:     "http://www.murc-kawasesouba.jp/fx/past/index.php?id=",
		dompath: "div#main table.data-table7 tbody tr",
	}
	return E
}

func (req *exchange) qs() (string, error) {
	t, err := time.Parse("2006-01-02", req.date)
	if err != nil {
		return "", err
	}
	y := t.Format("06")
	m := fmt.Sprintf("%02d", int(t.Month()))
	d := t.Format("02")
	return fmt.Sprintf("%v%v%v", y, m, d), nil
}

func (req *exchange) query() error {
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

	tr := doc.Find("div#main table.data-table7 tbody tr").Eq(1)
	h2 := doc.Find("div#main h2")

	req.result.date = strings.TrimSpace(strings.Split(h2.Text(), "As of ")[1])
	req.result.currency = strings.TrimSpace(tr.Find("td").First().Text())

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
	req.result.value = TTM

	return nil
}
