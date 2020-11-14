package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func parseCSV(i interface{}) {
	fmt.Printf("parsing csv...\n%s\n", i)
}

func parseSampleCSV() {
	file, err := os.Open("../data/values-2.csv")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}

	defer file.Close()

	values := []Values{}

	err = gocsv.UnmarshalFile(file, &values)
	if err != nil {
		fmt.Println("error unmarshaling file:", err)
		return
	}

	for _, v := range values {
		fmt.Println(v)
	}
}
