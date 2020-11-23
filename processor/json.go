package main

import (
	"encoding/json"
)

// jsonToMap formats json raw values into a map
func jsonToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	bytes, err := json.Marshal(rawValues)
	if err != nil {
		return nil, err
	}

	vals := map[string]json.RawMessage{}
	err = json.Unmarshal(bytes, &vals)

	return vals, err
}
