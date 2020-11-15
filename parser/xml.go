package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

func parseXML(rawValues interface{}) (*Values, error) {
	fmt.Println("parsing xml...")

	s, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("could not convert input to string")
	}

	s = strings.TrimPrefix(s, "<?xmlversion=\"1.0\"encoding=\"UTF-8\"?>")

	vals := &Values{}
	err := xml.Unmarshal([]byte(s), vals)
	if err != nil {
		return nil, err
	}

	return vals, nil
}
