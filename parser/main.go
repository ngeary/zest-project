package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	Values  Values `json:"values"`
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
}
