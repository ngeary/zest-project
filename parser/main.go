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
	ID                 int    `json:"id" csv:"id"`
	MemberID           string `json:"member_id" csv:"member_id"`
	FirstName          string `json:"first_name" csv:"first_name"`
	LastName           string `json:"last_name" csv:"last_name"`
	Address            string `json:"address" csv:"address"`
	DOB                string `json:"dob" csv:"dob"`
	CountryID          int    `json:"CountryID"`
	Employer           string `json:"Employer"`
	EmploymentType     int    `json:"EmploymentType"`
	EmpOrderNum        int    `json:"EmpOrderNum"`
	GrossMonthlyIncome int    `json:"GrossMonthlyIncome"`
	Position           string `json:"Position"`
	RetiredFlag        string `json:"RetiredFlag"`
	SelfEmpFlag        string `json:"SelfEmpFlag"`
	State              string `json:"State"`
}

func main() {
	req := Request{}

	file, err := ioutil.ReadFile("../data/dataset2.json")
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
