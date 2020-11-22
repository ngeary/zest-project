package main

import (
	"encoding/json"
	"fmt"
)

func parseJSON(rawValues interface{}) (*Values, error) {
	fmt.Println("parsing json...")

	bytes, err := json.Marshal(rawValues)
	if err != nil {
		return nil, err
	}

	vals := Values{}
	err = json.Unmarshal(bytes, &vals)
	if err != nil {
		return nil, err
	}

	return &vals, nil
}

func jsonToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	bytes, err := json.Marshal(rawValues)
	if err != nil {
		return nil, err
	}

	vals := map[string]json.RawMessage{}
	err = json.Unmarshal(bytes, &vals)

	return vals, err
}
