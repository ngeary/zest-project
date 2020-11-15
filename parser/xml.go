package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
)

func parseXML(rawValues interface{}) (*Values, error) {
	fmt.Println("parsing xml...")

	s, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("could not convert input to string")
	}

	vals := Values{}
	err := xml.Unmarshal([]byte(s), &vals)
	if err != nil {
		return nil, err
	}

	return &vals, nil
}

func removeXMLDeclarations(bytes []byte) []byte {
	// regex matches the following pattern: <?xml + (any number of any character) + version + (any number of any character) + ?>
	re := regexp.MustCompile(`<\?xml.*version.*\?>`)

	// replace each occurrence of XML declaration with empty byte slice
	return re.ReplaceAll(bytes, []byte{})
}
