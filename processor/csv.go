package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"strings"
)

// csvToMap formats csv raw values into a map
func csvToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	valsString, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("type assertion failed: input is not a string")
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
