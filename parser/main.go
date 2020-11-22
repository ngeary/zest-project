package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/ngeary/zest-project/anonymizer"
)

// Request represents a json request to insert data
type Request struct {
	RequestID string `json:"request_id"`
	Rows      []*Row
}

// Row struct
type Row struct {
	RowID   string    `json:"row_id"`
	Sources []*Source `json:"sources"`
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
	dataDir := "../data/"
	seenFiles := make(map[string]bool)

	for {
		fileInfos, err := ioutil.ReadDir(dataDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, fi := range fileInfos {
			if seenFiles[fi.Name()] || !strings.HasSuffix(strings.ToLower(fi.Name()), ".json") {
				continue
			}

			if fi.Name() != "dataset1.json" {
				continue
			}

			err = parse(dataDir + fi.Name())
			if err != nil {
				log.Println(err)
			}

			seenFiles[fi.Name()] = true
		}

		time.Sleep(time.Second)
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

	// var vals *Values

	// for _, r := range req.Rows {
	// 	for _, s := range r.Sources {
	// 		fmt.Printf("\nRequest ID: %s\tRow ID: %s\tSource Name: %s\n", req.RequestID, r.RowID, s.Name)

	// 		switch strings.ToLower(s.Format) {
	// 		case "json":
	// 			vals, err = parseJSON(s.Values)
	// 		case "csv":
	// 			vals, err = parseCSV(s.Values)
	// 		case "xml":
	// 			vals, err = parseXML(s.Values)
	// 		default:
	// 			err = errors.New("unrecognized data format: " + s.Format)
	// 		}

	// 		if err != nil {
	// 			log.Println("parsing error:", err)
	// 		}

	// 		fmt.Println("Values:", vals)
	// 	}
	// }

	for _, row := range req.Rows {
		for _, source := range row.Sources {
			if source.Name != "app_data" {
				continue
			}

			bytes, err := json.Marshal(source.Values)
			if err != nil {
				return err
			}

			vals := map[string]json.RawMessage{}
			err = json.Unmarshal(bytes, &vals)
			if err != nil {
				return err
			}

			anonData := anonymizer.GetAnonymousValues()
			for k, v := range anonData {
				vals[k] = v
			}

			source.Values = vals
		}
	}

	return writeToFile(&req)
}

func writeToFile(r *Request) error {
	bytes, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("../anon_data/4.json", bytes, 0644)
}
