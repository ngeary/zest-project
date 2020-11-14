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
	Name    string      `json:"name"`
	Version int         `json:"version"`
	Format  string      `json:"format"`
	Values  interface{} `json:"values"`
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
	req := Request{}

	file, err := ioutil.ReadFile("../data/dataset1.json")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}

	err = json.Unmarshal(file, &req)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v\n", err)
	}

	for _, r := range req.Rows {
		for _, s := range r.Sources {
			fmt.Printf("\nRequest ID: %s\tRow ID: %s\tSource Name: %s\n", req.RequestID, r.RowID, s.Name)

			switch strings.ToLower(s.Format) {
			case "json":
				parseJSON(s.Values)
			case "csv":
				parseCSV(s.Values)
			case "xml":
				parseXML(s.Values)
			default:
				log.Printf("unrecognized data format: %s\n", s.Format)
			}
		}
	}
}

func parseJSON(rawValues interface{}) {
	fmt.Printf("parsing json...\n%s\n", rawValues)

	bytes, err := json.Marshal(rawValues)
	if err != nil {
		log.Println("error converting raw data to byte slice")
		return
	}

	vals := Values{}
	err = json.Unmarshal(bytes, &vals)
	if err != nil {
		log.Println("error unmarshaling json values:", err)
		return
	}

	fmt.Println("Values:", vals)
}

func parseCSV(i interface{}) {
	fmt.Printf("parsing csv...\n%s\n", i)
}

func parseXML(i interface{}) {
	fmt.Printf("parsing xml...\n%s\n", i)
}
