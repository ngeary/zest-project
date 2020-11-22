package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gocarina/gocsv"
)

// CsvValues struct
type CsvValues struct {
	ID        string `csv:"id"`
	MemberID  string `csv:"member_id"`
	FirstName string `csv:"first_name"`
	LastName  string `csv:"last_name"`
	Address   string `csv:"address"`
	DOB       string `csv:"dob"`
}

func parseCSV(rawValues interface{}) (*Values, error) {
	fmt.Println("parsing csv...")

	s, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("could not convert input to string")
	}

	vals := []CsvValues{}
	err := gocsv.UnmarshalString(s, &vals)
	if err != nil {
		return nil, err
	}

	if len(vals) == 0 {
		return nil, errors.New("no values found")
	}

	v := &Values{
		ID:        vals[0].ID,
		MemberID:  vals[0].MemberID,
		FirstName: vals[0].FirstName,
		LastName:  vals[0].LastName,
		Address:   vals[0].Address,
		DOB:       vals[0].DOB,
	}

	return v, nil
}

func csvToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	valsString, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("input not a string")
	}

	r := csv.NewReader(strings.NewReader(valsString))

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, errors.New("input does not have at least two lines")
	}

	if len(records[0]) != len(records[1]) {
		return nil, errors.New("number of values does not match number of fields")
	}

	vals := map[string]json.RawMessage{}

	for i, field := range records[0] {
		k := strings.TrimSpace(field)
		v := "\"" + strings.TrimSpace(records[1][i]) + "\""

		vals[k] = []byte(v)
	}

	return vals, nil
}
