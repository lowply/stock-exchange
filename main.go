package main

import (
	"fmt"
	"log"
	"os"
)

type Result struct {
	date     string
	value    string
	currency string
}

func usage() {
	fmt.Println("Usage: go run . [ticker] [YYYY-MM-DD]")
	os.Exit(1)
}

func exec() error {
	S := NewStock()
	err := S.query()
	if err != nil {
		return err
	}
	fmt.Printf("%v, %v\n", S.result.date, S.result.value)

	E := NewExchange()
	err = E.query()
	if err != nil {
		return err
	}
	fmt.Printf("%v, %v, %v\n", E.result.date, E.result.currency, E.result.value)

	return nil
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}
