package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

func parseXML(rawValues interface{}) {
	fmt.Println("parsing xml...")

	s, ok := rawValues.(string)
	if !ok {
		fmt.Println("could not convert input to string")
		return
	}

	s = strings.TrimPrefix(s, "<?xmlversion=\"1.0\"encoding=\"UTF-8\"?>")

	vals := Values{}
	err := xml.Unmarshal([]byte(s), &vals)
	if err != nil {
		log.Println("error unmarshaling xml values:", err)
		return
	}

	fmt.Println("Values:", vals)
}
