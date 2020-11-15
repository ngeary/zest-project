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

	vals := &Values{}
	err = json.Unmarshal(bytes, vals)
	if err != nil {
		return nil, err
	}

	return vals, nil
}
