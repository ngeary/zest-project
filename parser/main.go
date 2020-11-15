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
	ID                 int    `json:"id" csv:"id"`
	MemberID           string `json:"member_id" csv:"member_id"`
	FirstName          string `json:"first_name" csv:"first_name"`
	LastName           string `json:"last_name" csv:"last_name"`
	Address            string `json:"address" csv:"address"`
	DOB                string `json:"dob" csv:"dob"`
	CountryID          int    `json:"CountryID" xml:"CountryID"`
	Employer           string `json:"Employer" xml:"Employer"`
	EmploymentType     int    `json:"EmploymentType" xml:"EmploymentType"`
	EmpOrderNum        int    `json:"EmpOrderNum" xml:"EmpOrderNum"`
	GrossMonthlyIncome int    `json:"GrossMonthlyIncome" xml:"GrossMonthlyIncome"`
	Position           string `json:"Position" xml:"Position"`
	RetiredFlag        string `json:"RetiredFlag" xml:"RetiredFlag"`
	SelfEmpFlag        string `json:"SelfEmpFlag" xml:"SelfEmpFlag"`
	State              string `json:"State" xml:"State"`
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
