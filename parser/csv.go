package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func parseCSV(i interface{}) {
	fmt.Printf("parsing csv...\n%s\n", i)
}

func parseSampleCSV() {
	// Open the file
	recordFile, err := os.Open("../data/values-2.csv")
	if err != nil {
		fmt.Println("open error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("read error encountered ::", err)
		return
	}

	for _, r := range allRecords {
		fmt.Println(r)
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("close error encountered ::", err)
		return
	}
}
