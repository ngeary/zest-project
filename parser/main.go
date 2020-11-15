package main

import (
	"encoding/json"
	"errors"
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
	ID                 interface{} `json:"id" csv:"id"`
	MemberID           interface{} `json:"member_id" csv:"member_id"`
	FirstName          interface{} `json:"first_name" csv:"first_name"`
	LastName           interface{} `json:"last_name" csv:"last_name"`
	Address            interface{} `json:"address" csv:"address"`
	DOB                interface{} `json:"dob" csv:"dob"`
	CountryID          interface{} `json:"CountryID" xml:"CountryID"`
	Employer           interface{} `json:"Employer" xml:"Employer"`
	EmploymentType     interface{} `json:"EmploymentType" xml:"EmploymentType"`
	EmpOrderNum        interface{} `json:"EmpOrderNum" xml:"EmpOrderNum"`
	GrossMonthlyIncome interface{} `json:"GrossMonthlyIncome" xml:"GrossMonthlyIncome"`
	Position           interface{} `json:"Position" xml:"Position"`
	RetiredFlag        interface{} `json:"RetiredFlag" xml:"RetiredFlag"`
	SelfEmpFlag        interface{} `json:"SelfEmpFlag" xml:"SelfEmpFlag"`
	State              interface{} `json:"State" xml:"State"`
}

func main() {
	files := []string{
		// "../data/ng-dataset1.json",
		"../data/dataset1.json",
		"../data/dataset2.json",
		"../data/dataset3.json",
		"../data/dataset4.json",
	}

	var err error

	for _, f := range files {
		err = parse(f)
		if err != nil {
			log.Println(err)
		}
	}
}

func parse(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	file = removeXMLDeclarations(file)

	req := Request{}

	err = json.Unmarshal(file, &req)
	if err != nil {
		return err
	}

	var vals *Values

	for _, r := range req.Rows {
		for _, s := range r.Sources {
			fmt.Printf("\nRequest ID: %s\tRow ID: %s\tSource Name: %s\n", req.RequestID, r.RowID, s.Name)

			switch strings.ToLower(s.Format) {
			case "json":
				vals, err = parseJSON(s.Values)
			case "csv":
				vals, err = parseCSV(s.Values)
			case "xml":
				vals, err = parseXML(s.Values)
			default:
				err = errors.New("unrecognized data format: " + s.Format)
			}

			if err != nil {
				log.Println("parsing error:", err)
			}

			fmt.Println("Values:", vals)
		}
	}

	return nil
}
