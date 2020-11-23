package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"regexp"
	"strings"
)

func removeXMLDeclarations(input []byte) []byte {
	// regex matches the following pattern: <?xml + (any number of any character) + version + (any number of any character) + ?>
	re := regexp.MustCompile(`<\?xml.*version.*\?>`)

	// replace each occurrence of XML declaration with empty byte slice
	return re.ReplaceAll(input, []byte{})
}

// xmlToMap formats xml raw values into a map
func xmlToMap(rawValues interface{}) (map[string]json.RawMessage, error) {
	valsString, ok := rawValues.(string)
	if !ok {
		return nil, errors.New("type assertion failed: input is not a string")
	}

	vals := map[string]json.RawMessage{}
	d := xml.NewDecoder(strings.NewReader(valsString))
	var field, value string

	for token, err := d.Token(); err == nil; token, err = d.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			field = strings.TrimSpace(t.Name.Local)
		case xml.CharData:
			value = string([]byte(t))
		case xml.EndElement:
			if t.Name.Local == "root" {
				continue
			}
			vals[field] = []byte("\"" + strings.TrimSpace(value) + "\"")
			field, value = "", ""
		}
	}

	return vals, nil
}
