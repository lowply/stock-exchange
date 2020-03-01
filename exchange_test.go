package main

import (
	"testing"
	"time"
)

type testCaseExchange struct {
	date   string
	should string
}

func TestExchange(t *testing.T) {
	testCases := []testCaseExchange{
		{date: "2020-02-14", should: "109.89"},
		{date: "2020-01-30", should: "109.01"},
	}

	for _, v := range testCases {
		t.Logf("Testing JPY/USD TTM on %v...", v.date)

		date, err := time.Parse("2006-01-02", v.date)
		if err != nil {
			t.Fatal(err)
		}

		e := NewExchange(
			date,
			"http://www.murc-kawasesouba.jp/fx/past/index.php?id=",
			"div#main table.data-table7 tbody tr",
		)

		e.get()

		if e.result != v.should {
			t.Errorf("Actual: %v, Should: %v\n", e.result, v.should)
		}
	}
}
