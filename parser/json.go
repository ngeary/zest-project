package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func parseJSON(rawValues interface{}) {
	fmt.Println("parsing json...")

	bytes, err := json.Marshal(rawValues)
	if err != nil {
		log.Println("error converting raw data to byte slice")
		return
	}

	vals := Values{}
	err = json.Unmarshal(bytes, &vals)
	if err != nil {
		log.Println("error unmarshaling json values:", err)
		return
	}

	fmt.Println("Values:", vals)
}
