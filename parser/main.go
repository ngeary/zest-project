package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Request represents a json request to insert data
type Request struct {
	RequestID string `json:"request_id"`
	Rows      []Row
}

// Row struct
type Row struct {
	RowID   string   `json:"row_id"`
	Sources []Source `json:"sources"`
}

// Source struct
type Source struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Format  string `json:"format"`
	// Values  Values `json:"values"`
}

// Values struct
type Values struct {
	ID        int    `json:"id"`
	MemberID  string `json:"member_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
}

func main() {
	data := Request{}

	file, err := ioutil.ReadFile("../data/dataset1.json")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v\n", err)
	}

	log.Println(data)

	for _, r := range data.Rows {
		for _, s := range r.Sources {
			fmt.Printf("Request ID: %s\tRow ID: %s\tSource Name: %s\t", data.RequestID, r.RowID, s.Name)

			switch strings.ToLower(s.Format) {
			case "json":
				parseJSON()
			case "csv":
				parseCSV()
			case "xml":
				parseXML()
			default:
				log.Printf("unrecognized data format: %s\n", s.Format)
			}
		}
	}
}

func parseJSON() {
	fmt.Println("parsing json...")
}

func parseCSV() {
	fmt.Println("parsing csv...")
}

func parseXML() {
	fmt.Println("parsing xml...")
}
