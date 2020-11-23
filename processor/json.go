package main

import (
	"encoding/json"
)

func jsonToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	bytes, err := json.Marshal(rawValues)
	if err != nil {
		return nil, err
	}

	vals := map[string]json.RawMessage{}
	err = json.Unmarshal(bytes, &vals)

	return vals, err
}
