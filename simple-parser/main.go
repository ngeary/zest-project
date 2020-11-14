package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	data := map[string]interface{}{}

	file, err := ioutil.ReadFile("/Users/nick/github-repos/zest-project/data/ngtest1.json")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v\n", err)
	}

	for _, d := range data {
		log.Println(d)
	}
}
