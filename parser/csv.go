package main

import (
	"fmt"
	"log"

	"github.com/gocarina/gocsv"
)

func parseCSV(rawValues interface{}) {
	fmt.Println("parsing csv...")

	s, ok := rawValues.(string)
	if !ok {
		fmt.Println("could not convert input to string")
		return
	}

	vals := []Values{}
	err := gocsv.UnmarshalString(s, &vals)
	if err != nil {
		log.Println("error unmarshaling csv values:", err)
		return
	}

	fmt.Println("Values:", vals)
}
