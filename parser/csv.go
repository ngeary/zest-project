package main

import (
	"errors"
	"fmt"

	"github.com/gocarina/gocsv"
)

func parseCSV(rawValues interface{}) (*Values, error) {
	fmt.Println("parsing csv...")

	s, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("could not convert input to string")
	}

	vals := []*Values{}
	err := gocsv.UnmarshalString(s, &vals)
	if err != nil {
		return nil, err
	}

	if len(vals) == 0 {
		return nil, errors.New("no values found")
	}

	return vals[0], nil
}
