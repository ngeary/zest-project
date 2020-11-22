package main

import (
	"errors"
	"fmt"

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
