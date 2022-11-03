package main

import (
	"testing"
	"time"
)

type testCaseStock struct {
	ticker string
	date   string
	should string
}

func TestStock(t *testing.T) {
	testCases := []testCaseStock{
		{ticker: "AMZN", date: "2018-11-28", should: "83.88"},
		{ticker: "MSFT", date: "2020-02-14", should: "185.35"},
		{ticker: "MSFT", date: "2020-02-15", should: ""},
	}

	for _, v := range testCases {
		t.Logf("Testing %v on %v...", v.ticker, v.date)

		date, err := time.Parse("2006-01-02", v.date)
		if err != nil {
			t.Fatal(err)
		}

		s := NewStock(
			v.ticker,
			date,
			"https://finance.yahoo.co.jp/quote/",
			"body div#wrapper div#root main div div div div section div table tbody tr td",
		)

		s.get()

		if s.result != v.should {
			t.Errorf("Actual: %v, Should: %v\n", s.result, v.should)
		}
	}
}
